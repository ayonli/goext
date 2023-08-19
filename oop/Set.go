package oop

import (
	"slices"

	sliceExt "github.com/ayonli/goext/slices"
)

type Set[T comparable] struct {
	items       []T
	size        int
	placeholder *T
}

func NewSet[T comparable](items []T) *Set[T] {
	self := Set[T]{}
	self.items = sliceExt.Uniq(items)
	self.size = len(self.items)
	self.placeholder = new(T)

	return &self
}

func (self *Set[T]) Add(item T) *Set[T] {
	if !slices.Contains(self.items, item) {
		self.items = append(self.items, item)
		self.size++
	}

	return self
}

func (self *Set[T]) Delete(item T) bool {
	idx := slices.Index(self.items, item)

	if idx == -1 {
		return false
	}

	self.items[idx] = *self.placeholder
	self.size--
	return true
}

func (self *Set[T]) Clear() {
	self.items = []T{}
	self.size = 0
}

func (self *Set[T]) ForEach(fn func(item T)) {
	for _, item := range self.items {
		if &item != self.placeholder {
			fn(item)
		}
	}
}

func (self *Set[T]) Has(item T) bool {
	for _, ele := range self.items {
		if &ele == &item {
			return true
		}
	}

	return false
}

func (self *Set[T]) Values() []T {
	items := make([]T, self.size)
	idx := 0

	for _, item := range self.items {
		if &item != self.placeholder {
			items[idx] = item
			idx++
		}
	}

	return items
}

func (self *Set[T]) Size() int {
	return self.size
}
