package collections

import (
	"encoding/json"
	"strings"

	"github.com/ayonli/goext/mapx"
)

// Thread-safe case-insensitive map, keys are case-insensitive.
type CiMap[K ~string, V any] struct {
	Map[K, V]
	keys []K
}

// Creates a new instance of the CiMap.
func NewCiMap[K ~string, V string](initial []MapEntry[K, V]) *CiMap[K, V] {
	m := &CiMap[K, V]{}

	for _, entry := range initial {
		m.Set(entry.Key, entry.Value)
	}

	return m
}

// Sets a pair of key and value in the map. If the key already exists, it changes the corresponding
// value; otherwise, it adds the new pair into the map.
func (self *CiMap[K, V]) Set(key K, value V) *CiMap[K, V] {
	self.mut.Lock()
	defer self.mut.Unlock()
	return self.set(key, value)
}

func (self *CiMap[K, V]) set(key K, value V) *CiMap[K, V] {
	id := strings.ToLower(string(key))
	idx := self.findIndex(K(id))

	if idx == -1 {
		self.records = append(self.records, mapRecordItem[K, V]{
			Key:     K(id),
			Value:   value,
			Deleted: false,
		})
		self.keys = append(self.keys, key)
		self.size++
	} else {
		record := &self.records[idx]
		record.Key = K(id)
		record.Value = value
		self.keys[idx] = key // also update the key
	}

	return self
}

// Retrieves a value by the given key. If the key doesn't exist, it returns the zero-value of type
// `V` and `false`.
func (self *CiMap[K, V]) Get(key K) (V, bool) {
	return self.Map.Get(K(strings.ToLower(string(key))))
}

// Checks if the given key exists in the map.
func (self *CiMap[K, V]) Has(key K) bool {
	return self.Map.Has(K(strings.ToLower(string(key))))
}

// Retrieves a value by the given key. If the key doesn't exist yet, invokes the `setter` function
// to set the value and return it.
//
// This function is atomic.
func (self *CiMap[K, V]) EnsureGet(key K, setter func() V) V {
	self.mut.Lock()
	defer self.mut.Unlock()

	id := strings.ToLower(string(key))
	idx := self.findIndex(K(id))
	var value V

	if idx == -1 {
		value = setter()
		self.set(key, value)
	} else {
		record := self.records[idx]
		value = record.Value
	}

	return value
}

// Removes the key-value pair by the given key.
func (self *CiMap[K, V]) Delete(key K) bool {
	self.mut.Lock()
	defer self.mut.Unlock()

	id := strings.ToLower(string(key))
	idx := self.findIndex(K(id))

	if idx == -1 {
		return false
	}

	self.deleteAt(idx)
	self.keys[idx] = ""

	return true
}

func (self *CiMap[K, V]) Clear() {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.records = nil
	self.size = 0
	self.keys = nil
}

// Retrieves all the keys in the map.
func (self *CiMap[K, V]) Keys() []K {
	self.mut.RLock()
	defer self.mut.RUnlock()

	items := make([]K, self.size)
	idx := 0

	for _, record := range self.records {
		if !record.Deleted {
			items[idx] = self.keys[idx]
			idx++
		}
	}

	return items
}

// Returns a channel for the map entries that can be used in the `for...range...` loop.
func (self *CiMap[K, V]) Entries() <-chan MapEntry[K, V] {
	channel := make(chan MapEntry[K, V])

	go func() {
		self.ForEach(func(value V, key K) {
			channel <- MapEntry[K, V]{Key: key, Value: value}
		})
		close(channel)
	}()

	return channel
}

// Loop through all the key-value pairs in the map and invoke the given function against them.
func (self *CiMap[K, V]) ForEach(fn func(value V, key K)) {
	for idx, record := range self.records {
		if !record.Deleted {
			fn(record.Value, self.keys[idx])
		}
	}
}

// Creates a builtin `map` based on this map.
func (self *CiMap[K, V]) ToMap() map[K]V {
	self.mut.RLock()
	defer self.mut.RUnlock()

	items := map[K]V{}

	for idx, record := range self.records {
		if !record.Deleted {
			key := self.keys[idx]
			items[key] = record.Value
		}
	}

	return items
}

func (self *CiMap[K, V]) getNormalizedRecords() []mapRecordItem[K, V] {
	records := make([]mapRecordItem[K, V], self.size)
	idx := 0

	for i, record := range self.records {
		if !record.Deleted {
			records[idx] = mapRecordItem[K, V]{
				Key:     self.keys[i],
				Value:   record.Value,
				Deleted: false,
			}
			idx++
		}
	}

	return records
}

func (self *CiMap[K, V]) String() string {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.formatString("collections.CiMap", self.getNormalizedRecords())
}

func (self *CiMap[K, V]) GoString() string {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.formatGoString("collections.CiMap", self.getNormalizedRecords())
}

func (self *CiMap[K, V]) UnmarshalJSON(data []byte) error {
	self.mut.Lock()
	defer self.mut.Unlock()

	var m map[K]V

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	for _, key := range mapx.Keys(m) { // mapx.Keys() guarantees keys are ordered alphabetically
		self.set(key, m[key])
	}

	return nil
}

func (self *CiMap[K, V]) MarshalJSON() ([]byte, error) {
	self.mut.RLock()
	defer self.mut.RUnlock()

	str := "{"
	started := false

	for idx, record := range self.records {
		if !record.Deleted {
			if started {
				str += ","
			} else {
				started = true
			}

			key := self.keys[idx]
			keyBytes, err := json.Marshal(key)

			if err != nil {
				return []byte{}, err
			} else {
				str += string(keyBytes) + ":"
			}

			valueBytes, err := json.Marshal(record.Value)

			if err != nil {
				return []byte{}, err
			} else {
				str += string(valueBytes)
			}
		}
	}

	str += "}"

	return []byte(str), nil
}
