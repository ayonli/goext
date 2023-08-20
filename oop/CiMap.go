package oop

import (
	"slices"
	"strings"

	"github.com/ayonli/goext/slicex"
)

type CiMapRecordItem[K comparable, V any] struct {
	Id      string
	Key     K
	Value   V
	Deleted bool
}

// Case-insensitive map, keys are case-insensitive.
type CiMap[K ~string, V any] struct {
	records []CiMapRecordItem[K, V]
	size    int
}

// Creates a new instance of the CiMap.
func NewCiMap[K ~string, V any]() *CiMap[K, V] {
	self := CiMap[K, V]{
		records: []CiMapRecordItem[K, V]{},
		size:    0,
	}
	return &self
}

func (self *CiMap[K, V]) findIndex(id string) int {
	return slices.IndexFunc(self.records, func(record CiMapRecordItem[K, V]) bool {
		return record.Id == id && !record.Deleted
	})
}

// Sets a pair of key and value in the map. If the key already exists, it changes the corresponding
// value; otherwise, it adds the new pair into the map.
func (self *CiMap[K, V]) Set(key K, value V) *CiMap[K, V] {
	id := strings.ToLower(string(key))
	idx := self.findIndex(id)

	if idx == -1 {
		self.records = append(self.records, CiMapRecordItem[K, V]{
			Id:      id,
			Key:     key,
			Value:   value,
			Deleted: false,
		})
		self.size++
	} else {
		record := &self.records[idx]
		record.Key = key // also update the key
		record.Value = value
	}

	return self
}

// Retrieves a value by the given key. If the key doesn't exist, it returns the zero-value of type
// `V` and `false`.
func (self *CiMap[K, V]) Get(key K) (V, bool) {
	id := strings.ToLower(string(key))
	idx := self.findIndex(id)

	if idx == -1 {
		return *new(V), false
	}

	record := self.records[idx]
	return record.Value, true
}

// Checks if the given key exists in the map.
func (self *CiMap[K, V]) Has(key K) bool {
	id := strings.ToLower(string(key))
	idx := self.findIndex(id)
	return idx != -1
}

// Removes the key-value pair by the given key.
func (self *CiMap[K, V]) Delete(key K) bool {
	id := strings.ToLower(string(key))
	idx := self.findIndex(id)

	if idx == -1 {
		return false
	}

	record := &self.records[idx]
	record.Id = ""
	record.Key = *new(K)
	record.Value = *new(V)
	record.Deleted = true
	self.size--

	// Optimize memory, when too much records are deleted, re-allocate the internal list.
	if limit := len(self.records); limit >= 100 && self.size <= int(limit/3) {
		self.records = slicex.Filter(self.records, func(item CiMapRecordItem[K, V], idx int) bool {
			return !item.Deleted
		})
	}

	return true
}

// Empties the map and resets its size.
func (self *CiMap[K, V]) Clear() {
	self.records = []CiMapRecordItem[K, V]{}
	self.size = 0
}

// Retrieves all the keys in the map.
func (self *CiMap[K, V]) Keys() []K {
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
func (self *CiMap[K, V]) Values() []V {
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
func (self *CiMap[K, V]) ToMap() map[K]V {
	items := map[K]V{}

	for _, record := range self.records {
		if !record.Deleted {
			items[record.Key] = record.Value
		}
	}

	return items
}

// Loop through all the key-value pairs in the map and invoke the given function against them.
func (self *CiMap[K, V]) ForEach(fn func(value V, key K)) {
	for _, record := range self.records {
		if !record.Deleted {
			fn(record.Value, record.Key)
		}
	}
}

// Returns the size of the map.
func (self *CiMap[K, V]) Size() int {
	return self.size
}
