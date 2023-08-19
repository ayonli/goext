package oop

import (
	"slices"
	"strings"
)

type CiMap[K ~string, V any] struct {
	keys        []K
	values      []V
	ids         []string
	size        int
	placeholder *V
}

func NewCiMap[K ~string, V any](initial map[K]V) *CiMap[K, V] {
	self := CiMap[K, V]{
		keys:        []K{},
		values:      []V{},
		ids:         []string{},
		placeholder: new(V),
	}

	for key, value := range initial {
		id := strings.ToLower(string(key))
		self.keys = append(self.keys, key)
		self.values = append(self.values, value)
		self.ids = append(self.ids, id)
	}

	self.size = len(self.ids)

	return &self
}

func (self *CiMap[K, V]) Set(key K, value V) *CiMap[K, V] {
	id := strings.ToLower(string(key))
	idx := slices.Index(self.ids, id)

	if idx == -1 {
		self.keys = append(self.keys, key)
		self.values = append(self.values, value)
		self.ids = append(self.ids, id)
		self.size++
	} else {
		self.keys[idx] = key // also update the key
		self.values[idx] = value
	}

	return self
}

func (self *CiMap[K, V]) Get(key K) (V, bool) {
	id := strings.ToLower(string(key))
	idx := slices.Index(self.ids, id)

	if idx == -1 {
		return *self.placeholder, false
	}

	return self.values[idx], true
}

func (self *CiMap[K, V]) Delete(key K) bool {
	id := strings.ToLower(string(key))
	idx := slices.Index(self.ids, id)

	if idx == -1 {
		return false
	}

	self.keys[idx] = ""
	self.values[idx] = *self.placeholder
	self.size--

	return true
}

func (self *CiMap[K, V]) Clear() {
	self.keys = []K{}
	self.values = []V{}
	self.ids = []string{}
	self.size = 0
}

func (self *CiMap[K, V]) ForEach(fn func(value V, key K)) {
	for idx := range self.ids {
		value := self.values[idx]

		if &value != self.placeholder {
			key := self.keys[idx]
			fn(value, key)
		}
	}
}

func (self *CiMap[K, V]) Has(key K) bool {
	id := strings.ToLower(string(key))
	idx := slices.Index(self.ids, id)

	if idx == -1 {
		return false
	}

	value := self.values[idx]

	if &value == self.placeholder {
		return false
	}

	return true
}

func (self *CiMap[K, V]) Keys() []K {
	items := make([]K, self.size)
	idx := 0

	for i, value := range self.values {
		if &value != self.placeholder {
			key := self.keys[i]
			items[idx] = key
			idx++
		}
	}

	return items
}

func (self *CiMap[K, V]) Values() []V {
	items := make([]V, self.size)
	idx := 0

	for _, value := range self.values {
		if &value != self.placeholder {
			items[idx] = value
			idx++
		}
	}

	return items
}

func (self *CiMap[K, V]) Entries() [][]any {
	entries := make([][]any, self.size)
	idx := 0

	for i, value := range self.values {
		if &value != self.placeholder {
			key := self.keys[i]
			entries[idx] = []any{key, value}
			idx++
		}
	}

	return entries
}

func (self *CiMap[K, V]) Size() int {
	return self.size
}
