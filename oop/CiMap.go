package oop

import (
	"slices"
	"strings"
)

type CiMapRecordItem[K comparable, V any] struct {
	Id      string
	Key     K
	Value   V
	Deleted bool
}

type CiMap[K ~string, V any] struct {
	records []CiMapRecordItem[K, V]
	size    int
}

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

func (self *CiMap[K, V]) Get(key K) (V, bool) {
	id := strings.ToLower(string(key))
	idx := self.findIndex(id)

	if idx == -1 {
		return *new(V), false
	}

	record := self.records[idx]
	return record.Value, true
}

func (self *CiMap[K, V]) Has(key K) bool {
	id := strings.ToLower(string(key))
	idx := self.findIndex(id)
	return idx != -1
}

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

	return true
}

func (self *CiMap[K, V]) Clear() {
	self.records = []CiMapRecordItem[K, V]{}
	self.size = 0
}

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

func (self *CiMap[K, V]) ForEach(fn func(value V, key K)) {
	for _, record := range self.records {
		if !record.Deleted {
			fn(record.Value, record.Key)
		}
	}
}

func (self *CiMap[K, V]) Size() int {
	return self.size
}
