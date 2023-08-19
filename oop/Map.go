package oop

import (
	"slices"
)

type MapRecordItem[K comparable, V any] struct {
	Key     K
	Value   V
	Deleted bool
}

// Unlike the builtin `map` type, this Map stores data in a underlying list, which provides ordered
// keys sequence. However, the op time is O(n), which might be insufficient for large amount of data.
// Use with caution.
type Map[K comparable, V any] struct {
	records []MapRecordItem[K, V]
	size    int
}

// Creates a new instance of the Map.
func NewMap[K comparable, V any]() *Map[K, V] {
	self := Map[K, V]{
		records: []MapRecordItem[K, V]{},
		size:    0,
	}
	return &self
}

func (self *Map[K, V]) findIndex(key K) int {
	return slices.IndexFunc(self.records, func(record MapRecordItem[K, V]) bool {
		return record.Key == key && !record.Deleted
	})
}

// Sets a pair of key and value in the map. If the key already exists, it changes the corresponding
// value; otherwise, it adds the new pair into the map.
func (self *Map[K, V]) Set(key K, value V) *Map[K, V] {
	idx := self.findIndex(key)

	if idx == -1 {
		self.records = append(self.records, MapRecordItem[K, V]{
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

// Removes the key-value pair by the given key.
func (self *Map[K, V]) Delete(key K) bool {
	idx := self.findIndex(key)

	if idx == -1 {
		return false
	}

	record := &self.records[idx] // must use & (ref) in order to mutate the object
	record.Key = *new(K)
	record.Value = *new(V)
	record.Deleted = true
	self.size--

	return true
}

// Empties the map and resets its size.
func (self *Map[K, V]) Clear() {
	self.records = []MapRecordItem[K, V]{}
	self.size = 0
}

// Retrieves all the keys in the map.
func (self *Map[K, V]) Keys() []K {
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

// Creates a builtin `map` based on this map.
func (self *Map[K, V]) ToMap() map[K]V {
	items := map[K]V{}

	for _, record := range self.records {
		if !record.Deleted {
			items[record.Key] = record.Value
		}
	}

	return items
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
