package collections

import (
	"encoding/json"
	"slices"

	"github.com/ayonli/goext/mapx"
)

// Thread-safe bi-directional map, keys and values are unique and map to each other.
type BiMap[K comparable, V comparable] struct {
	Map[K, V]
}

// Creates a new instance of the BiMap.
func NewBiMap[K comparable, V comparable](initial []MapEntry[K, V]) *BiMap[K, V] {
	m := &BiMap[K, V]{}

	for _, entry := range initial {
		m.set(entry.Key, entry.Value)
	}

	return m
}

func (self BiMap[K, V]) findIndexByValue(value V) int {
	return slices.IndexFunc(self.records, func(record mapRecordItem[K, V]) bool {
		return record.Value == value && !record.Deleted
	})
}

// Sets a pair of key and value in the map. If the key already exists, it changes the corresponding
// value; if the value already exists, it changes the corresponding key; if both are missing, it
// adds the new pair into the map.
func (self *BiMap[K, V]) Set(key K, value V) *BiMap[K, V] {
	self.mut.Lock()
	defer self.mut.Unlock()
	return self.set(key, value)
}

func (self *BiMap[K, V]) set(key K, value V) *BiMap[K, V] {
	idx := self.findIndex(key)

	if idx == -1 {
		idx = self.findIndexByValue(value)
	}

	if idx == -1 {
		self.records = append(self.records, mapRecordItem[K, V]{
			Key:     key,
			Value:   value,
			Deleted: false,
		})
		self.size++
	} else {
		record := &self.records[idx]
		record.Key = key     // update both the key
		record.Value = value // and the value
	}

	return self
}

// Retrieves a key by the given value. If the value doesn't exist, it returns the zero-value of type
// `K` and `false`.
func (self *BiMap[K, V]) GetKey(value V) (K, bool) {
	self.mut.RLock()
	defer self.mut.RUnlock()

	idx := self.findIndexByValue(value)

	if idx == -1 {
		return *new(K), false
	}

	record := self.records[idx]
	return record.Key, true
}

// Checks if the given value exists in the map.
func (self *BiMap[K, V]) HasValue(value V) bool {
	idx := self.findIndexByValue(value)
	return idx != -1
}

// Removes the key-value pair by the given value.
func (self *BiMap[K, V]) DeleteValue(value V) bool {
	self.mut.Lock()
	defer self.mut.Unlock()

	idx := self.findIndexByValue(value)
	return self.deleteAt(idx)
}

func (self *BiMap[K, V]) String() string {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.formatString("collections.BiMap", self.records)
}

func (self *BiMap[K, V]) GoString() string {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.formatGoString("collections.BiMap", self.records)
}

func (self *BiMap[K, V]) UnmarshalJSON(data []byte) error {
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
