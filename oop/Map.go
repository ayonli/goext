package oop

import (
	"slices"
)

type Map[K comparable, V any] struct {
	keys         []K
	values       []V
	size         int
	kPlaceholder *K
	vPlaceholder *V
}

func NewMap[K comparable, V any](entries [][]any) *Map[K, V] {
	self := Map[K, V]{
		keys:         []K{},
		values:       []V{},
		kPlaceholder: new(K),
		vPlaceholder: new(V),
	}

	for _, item := range entries {
		key, kOk := item[0].(K)
		value, vOk := item[1].(V)

		if !kOk || !vOk {
			continue
		} else {
			idx := slices.Index(self.keys, key)

			if idx == -1 {
				self.keys = append(self.keys, key)
				self.values = append(self.values, value)
			} else {
				self.values[idx] = value
			}
		}
	}

	self.size = len(self.keys)

	return &self
}

func (self *Map[K, V]) Set(key K, value V) *Map[K, V] {
	idx := slices.Index(self.keys, key)

	if idx == -1 {
		self.keys = append(self.keys, key)
		self.values = append(self.values, value)
		self.size++
	} else {
		self.values[idx] = value
	}

	return self
}

func (self *Map[K, V]) Get(key K) (V, bool) {
	idx := slices.Index(self.keys, key)

	if idx == -1 {
		return *self.vPlaceholder, false
	}

	return self.values[idx], true
}

func (self *Map[K, V]) Delete(key K) bool {
	idx := slices.Index(self.keys, key)

	if idx == -1 {
		return false
	}

	self.keys[idx] = *self.kPlaceholder
	self.values[idx] = *self.vPlaceholder
	self.size--
	return true
}

func (self *Map[K, V]) Clear() {
	self.keys = []K{}
	self.values = []V{}
	self.size = 0
}

func (self *Map[K, V]) ForEach(fn func(value V, key K)) {
	for idx, key := range self.keys {
		if &key != self.kPlaceholder {
			value := self.values[idx]
			fn(value, key)
		}
	}
}

func (self *Map[K, V]) Has(key K) bool {
	for _, ele := range self.keys {
		if &ele == &key {
			return true
		}
	}

	return false
}

func (self *Map[K, V]) Keys() []K {
	items := make([]K, self.size)
	idx := 0

	for _, item := range self.keys {
		if &item != self.kPlaceholder {
			items[idx] = item
			idx++
		}
	}

	return items
}

func (self *Map[K, V]) Values() []V {
	items := make([]V, self.size)
	idx := 0

	for _, item := range self.values {
		if &item != self.vPlaceholder {
			items[idx] = item
			idx++
		}
	}

	return items
}

func (self *Map[K, V]) Entries() [][]any {
	entries := make([][]any, self.size)
	idx := 0

	for i, key := range self.keys {
		if &key != self.kPlaceholder {
			value := self.values[i]

			entries[idx] = []any{key, value}
			idx++
		}
	}

	return entries
}

func (self *Map[K, V]) Size() int {
	return self.size
}
