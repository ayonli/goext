package collections

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/ayonli/goext/mapx"
	"github.com/ayonli/goext/slicex"
)

type mapRecordItem[K comparable, V any] struct {
	Key     K
	Value   V
	Deleted bool
}

type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// Map is an object-oriented collection of map with ordered keys and thread-safe by default.
//
// Unlike the builtin `map` type, this Map stores data in a underlying list, which provides ordered
// keys sequence. However, the op time is O(n), which might be inefficient for large amount of data.
// Use with caution.
type Map[K comparable, V any] struct {
	records []mapRecordItem[K, V]
	size    int
	mut     sync.RWMutex
}

// Creates a new instance of the Map.
func NewMap[K comparable, V any](initial []MapEntry[K, V]) *Map[K, V] {
	m := &Map[K, V]{}

	for _, entry := range initial {
		m.set(entry.Key, entry.Value)
	}

	return m
}

func (self *Map[K, V]) findIndex(key K) int {
	return slices.IndexFunc(self.records, func(record mapRecordItem[K, V]) bool {
		return record.Key == key && !record.Deleted
	})
}

// Sets a pair of key and value in the map. If the key already exists, it changes the corresponding
// value; otherwise, it adds the new pair into the map.
func (self *Map[K, V]) Set(key K, value V) *Map[K, V] {
	self.mut.Lock()
	defer self.mut.Unlock()
	return self.set(key, value)
}

func (self *Map[K, V]) set(key K, value V) *Map[K, V] {
	idx := self.findIndex(key)

	if idx == -1 {
		self.records = append(self.records, mapRecordItem[K, V]{
			Key:     key,
			Value:   value,
			Deleted: false,
		})
		self.size++
	} else {
		self.records[idx].Value = value
	}

	return self
}

// Retrieves a value by the given key. If the key doesn't exist, it returns the zero-value of type
// `V` and `false`.
func (self *Map[K, V]) Get(key K) (V, bool) {
	self.mut.RLock()
	defer self.mut.RUnlock()

	idx := self.findIndex(key)

	if idx == -1 {
		return *new(V), false
	}

	record := self.records[idx]
	return record.Value, true
}

// Checks if the given key exists in the map.
func (self *Map[K, V]) Has(key K) bool {
	idx := self.findIndex(key)
	return idx != -1
}

// Retrieves a value by the given key. If the key doesn't exist yet, invokes the `setter` function
// to set the value and return it.
//
// This function is atomic.
func (self *Map[K, V]) EnsureGet(key K, setter func() V) V {
	self.mut.Lock()
	defer self.mut.Unlock()

	idx := self.findIndex(key)
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

func (self *Map[K, V]) deleteAt(idx int) bool {
	if idx == -1 {
		return false
	}

	record := &self.records[idx] // must use & (ref) in order to mutate the object
	record.Key = *new(K)
	record.Value = *new(V)
	record.Deleted = true
	self.size--

	// Optimize memory, when too much records are deleted, re-allocate the internal list.
	if limit := len(self.records); limit >= 100 && self.size <= int(limit/3) {
		self.records = slicex.Filter(self.records, func(item mapRecordItem[K, V], idx int) bool {
			return !item.Deleted
		})
	}

	return true
}

// Removes the key-value pair by the given key.
func (self *Map[K, V]) Delete(key K) bool {
	self.mut.Lock()
	defer self.mut.Unlock()

	idx := self.findIndex(key)
	return self.deleteAt(idx)
}

// Removes and returns the key-value pair by the given key.
func (self *Map[K, V]) Pop(key K) (V, bool) {
	self.mut.Lock()
	defer self.mut.Unlock()

	idx := self.findIndex(key)

	if idx == -1 {
		return *new(V), false
	}

	record := self.records[idx]
	self.deleteAt(idx)

	return record.Value, true
}

// Empties the map and resets its size.
func (self *Map[K, V]) Clear() {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.records = nil
	self.size = 0
}

// Retrieves all the keys in the map.
func (self *Map[K, V]) Keys() []K {
	self.mut.RLock()
	defer self.mut.RUnlock()

	items := make([]K, self.size)
	idx := 0

	for _, record := range self.records {
		if !record.Deleted {
			items[idx] = record.Key
			idx++
		}
	}

	return items
}

// Retrieves all the values in the map.
func (self *Map[K, V]) Values() []V {
	self.mut.RLock()
	defer self.mut.RUnlock()

	items := make([]V, self.size)
	idx := 0

	for _, record := range self.records {
		if !record.Deleted {
			items[idx] = record.Value
			idx++
		}
	}

	return items
}

// Returns a channel for the map entries that can be used in the `for...range...` loop.
func (self *Map[K, V]) Entries() <-chan MapEntry[K, V] {
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
func (self *Map[K, V]) ForEach(fn func(value V, key K)) {
	for _, record := range self.records {
		if !record.Deleted {
			fn(record.Value, record.Key)
		}
	}
}

// Returns the size of the map.
func (self *Map[K, V]) Size() int {
	return self.size
}

// Creates a builtin `map` based on this map.
func (self *Map[K, V]) ToMap() map[K]V {
	self.mut.RLock()
	defer self.mut.RUnlock()

	items := map[K]V{}

	for _, record := range self.records {
		if !record.Deleted {
			items[record.Key] = record.Value
		}
	}

	return items
}

func (self *Map[K, V]) formatString(typeName string, records []mapRecordItem[K, V]) string {
	str := "&" + typeName + "["
	started := false

	for _, record := range records {
		if record.Deleted {
			continue
		}

		if started {
			str += " "
		} else {
			started = true
		}

		str += fmt.Sprint(record.Key) + ":" + fmt.Sprint(record.Value)
	}

	str += "]"
	return str
}

func (self *Map[K, V]) String() string {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.formatString("collections.Map", self.records)
}

func (self *Map[K, V]) formatGoString(typeName string, records []mapRecordItem[K, V]) string {
	mapStr := fmt.Sprintf("%#v", map[K]V{})
	idx1 := strings.Index(mapStr, "[")
	idx2 := strings.Index(mapStr, "]")
	idx3 := strings.Index(mapStr, "{")
	keyType := mapStr[idx1+1 : idx2]
	valueType := mapStr[idx2+1 : idx3]

	str := "&" + typeName + "[" + keyType + ", " + valueType + "]{"
	started := false

	for _, record := range records {
		if !record.Deleted {
			if started {
				str += ", "
			} else {
				started = true
			}

			str += fmt.Sprintf("%#v:%#v", record.Key, record.Value)
		}
	}

	str += "}"
	return str
}

func (self *Map[K, V]) GoString() string {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.formatGoString("collections.Map", self.records)
}

func (self *Map[K, V]) UnmarshalJSON(data []byte) error {
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

func (self *Map[K, V]) MarshalJSON() ([]byte, error) {
	self.mut.RLock()
	defer self.mut.RUnlock()

	str := "{"
	started := false

	for _, record := range self.records {
		if !record.Deleted {
			if started {
				str += ","
			} else {
				started = true
			}

			keyBytes, err := json.Marshal(record.Key)

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
