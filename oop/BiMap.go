package oop

import "slices"

// Bi-directional map, keys and values are unique and map to each other.
type BiMap[K comparable, V comparable] struct {
	records []MapRecordItem[K, V]
	size    int
}

// Creates a new instance of the BiMap.
func NewBiMap[K comparable, V comparable]() *BiMap[K, V] {
	self := BiMap[K, V]{
		records: []MapRecordItem[K, V]{},
		size:    0,
	}
	return &self
}

func (self *BiMap[K, V]) findIndex(key K) int {
	return slices.IndexFunc(self.records, func(record MapRecordItem[K, V]) bool {
		return record.Key == key && !record.Deleted
	})
}

func (self BiMap[K, V]) findIndexByValue(value V) int {
	return slices.IndexFunc(self.records, func(record MapRecordItem[K, V]) bool {
		return record.Value == value && !record.Deleted
	})
}

// Sets a pair of key and value in the map. If the key already exists, it changes the corresponding
// value; if the value already exists, it changes the corresponding key; if both are missing, it
// adds the new pair into the map.
func (self *BiMap[K, V]) Set(key K, value V) *BiMap[K, V] {
	idx := self.findIndex(key)

	if idx == -1 {
		idx = self.findIndexByValue(value)
	}

	if idx == -1 {
		self.records = append(self.records, MapRecordItem[K, V]{
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

// Retrieves a value by the given key. If the key doesn't exist, it returns the zero-value of type
// `V` and `false`.
func (self *BiMap[K, V]) Get(key K) (V, bool) {
	idx := self.findIndex(key)

	if idx == -1 {
		return *new(V), false
	}

	record := self.records[idx]
	return record.Value, true
}

// Retrieves a key by the given value. If the value doesn't exist, it returns the zero-value of type
// `K` and `false`.
func (self *BiMap[K, V]) GetKey(value V) (K, bool) {
	idx := self.findIndexByValue(value)

	if idx == -1 {
		return *new(K), false
	}

	record := self.records[idx]
	return record.Key, true
}

// Checks if the given key exists in the map.
func (self *BiMap[K, V]) Has(key K) bool {
	idx := self.findIndex(key)
	return idx != -1
}

// Checks if the given value exists in the map.
func (self *BiMap[K, V]) HasValue(value V) bool {
	idx := self.findIndexByValue(value)
	return idx != -1
}

func (self *BiMap[K, V]) deleteAt(idx int) bool {
	if idx == -1 {
		return false
	}

	record := &self.records[idx]
	record.Key = *new(K)
	record.Value = *new(V)
	record.Deleted = true
	self.size--

	return true
}

// Removes the key-value pair by the given key.
func (self *BiMap[K, V]) Delete(key K) bool {
	idx := self.findIndex(key)
	return self.deleteAt(idx)
}

// Removes the key-value pair by the given value.
func (self *BiMap[K, V]) DeleteValue(value V) bool {
	idx := self.findIndexByValue(value)
	return self.deleteAt(idx)
}

// Empties the map and reset its size.
func (self *BiMap[K, V]) Clear() {
	self.records = []MapRecordItem[K, V]{}
	self.size = 0
}

// Retrieves all the keys in the map.
func (self *BiMap[K, V]) Keys() []K {
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
func (self *BiMap[K, V]) Values() []V {
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

// Loop through all the key-value pairs in the map and invoke the given function against them.
func (self *BiMap[K, V]) ForEach(fn func(value V, key K)) {
	for _, record := range self.records {
		if !record.Deleted {
			fn(record.Value, record.Key)
		}
	}
}

// Returns the size of the map.
func (self *BiMap[K, V]) Size() int {
	return self.size
}
