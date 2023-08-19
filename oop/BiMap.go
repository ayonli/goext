package oop

import "slices"

type BiMap[K comparable, V comparable] struct {
	keys         []K
	values       []V
	size         int
	kPlaceholder *K
	vPlaceholder *V
}

func NewBiMap[K comparable, V comparable](entries [][]any) *BiMap[K, V] {
	self := BiMap[K, V]{
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
				idx = slices.Index(self.values, value)

				if idx == -1 {
					self.keys = append(self.keys, key)
					self.values = append(self.values, value)
				} else {
					self.keys[idx] = key
				}
			} else {
				self.values[idx] = value
			}
		}
	}

	self.size = len(self.keys)

	return &self
}

func (self *BiMap[K, V]) Set(key K, value V) *BiMap[K, V] {
	idx := slices.Index(self.keys, key)

	if idx == -1 {
		idx = slices.Index(self.values, value)

		if idx == -1 {
			self.keys = append(self.keys, key)
			self.values = append(self.values, value)
			self.size++
		} else {
			self.keys[idx] = key
		}
	} else {
		self.values[idx] = value
	}

	return self
}

func (self *BiMap[K, V]) Get(key K) (V, bool) {
	idx := slices.Index(self.keys, key)

	if idx == -1 {
		return *self.vPlaceholder, false
	}

	return self.values[idx], true
}

func (self *BiMap[K, V]) GetKey(value V) (K, bool) {
	idx := slices.Index(self.values, value)

	if idx == -1 {
		return *self.kPlaceholder, false
	}

	return self.keys[idx], true
}

func (self *BiMap[K, V]) Delete(key K) bool {
	idx := slices.Index(self.keys, key)

	if idx == -1 {
		return false
	}

	self.keys[idx] = *self.kPlaceholder
	self.values[idx] = *self.vPlaceholder
	self.size--
	return true
}

func (self *BiMap[K, V]) DeleteValue(value V) bool {
	idx := slices.Index(self.values, value)

	if idx == -1 {
		return false
	}

	self.keys[idx] = *self.kPlaceholder
	self.values[idx] = *self.vPlaceholder
	self.size--
	return true
}

func (self *BiMap[K, V]) Clear() {
	self.keys = []K{}
	self.values = []V{}
	self.size = 0
}

func (self *BiMap[K, V]) ForEach(fn func(value V, key K)) {
	for idx, key := range self.keys {
		if &key != self.kPlaceholder {
			value := self.values[idx]
			fn(value, key)
		}
	}
}

func (self *BiMap[K, V]) Has(key K) bool {
	for _, ele := range self.keys {
		if &ele == &key {
			return true
		}
	}

	return false
}

func (self *BiMap[K, V]) HasValue(value V) bool {
	for _, ele := range self.values {
		if &ele == &value {
			return true
		}
	}

	return false
}

func (self *BiMap[K, V]) Keys() []K {
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

func (self *BiMap[K, V]) Values() []V {
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

func (self *BiMap[K, V]) Entries() [][]any {
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

func (self *BiMap[K, V]) Size() int {
	return self.size
}
