package oop

import (
	"slices"
)

type MapRecordItem[K comparable, V any] struct {
	Key     K
	Value   V
	Deleted bool
}

type Map[K comparable, V any] struct {
	records []MapRecordItem[K, V]
	size    int
}

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

func (self *Map[K, V]) Get(key K) (V, bool) {
	idx := self.findIndex(key)

	if idx == -1 {
		return *new(V), false
	}

	record := self.records[idx]
	return record.Value, true
}

func (self *Map[K, V]) Has(key K) bool {
	idx := self.findIndex(key)
	return idx != -1
}

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

func (self *Map[K, V]) Clear() {
	self.records = []MapRecordItem[K, V]{}
	self.size = 0
}

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

func (self *Map[K, V]) ForEach(fn func(value V, key K)) {
	for _, record := range self.records {
		if !record.Deleted {
			fn(record.Value, record.Key)
		}
	}
}

func (self *Map[K, V]) Size() int {
	return self.size
}
