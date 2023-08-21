package oop

import (
	"strings"
)

// Case-insensitive map, keys are case-insensitive.
type CiMap[K ~string, V any] struct {
	Map[K, V]
	keys []K
}

// Creates a new instance of the CiMap.
func NewCiMap[K ~string, V string]() *CiMap[K, V] {
	self := CiMap[K, V]{}
	return &self
}

// Sets a pair of key and value in the map. If the key already exists, it changes the corresponding
// value; otherwise, it adds the new pair into the map.
func (self *CiMap[K, V]) Set(key K, value V) *CiMap[K, V] {
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

// Removes the key-value pair by the given key.
func (self *CiMap[K, V]) Delete(key K) bool {
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
	self.Map.Clear()
	self.keys = nil
}

// Retrieves all the keys in the map.
func (self *CiMap[K, V]) Keys() []K {
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

// Creates a builtin `map` based on this map.
func (self *CiMap[K, V]) ToMap() map[K]V {
	items := map[K]V{}

	for idx, record := range self.records {
		if !record.Deleted {
			key := self.keys[idx]
			items[key] = record.Value
		}
	}

	return items
}

// Loop through all the key-value pairs in the map and invoke the given function against them.
func (self *CiMap[K, V]) ForEach(fn func(value V, key K)) {
	for idx, record := range self.records {
		if !record.Deleted {
			fn(record.Value, self.keys[idx])
		}
	}
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
	return self.formatString("oop.CiMap", self.getNormalizedRecords())
}

func (self *CiMap[K, V]) GoString() string {
	return self.formatGoString("oop.CiMap", self.getNormalizedRecords())
}
