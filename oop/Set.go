package oop

import (
	"slices"
)

// Set is an object-oriented collection that stores unique items.
type Set[T comparable] struct {
	records []MapRecordItem[int, T]
	size    int
}

// Creates a new instance of the Set.
func NewSet[T comparable](base []T) *Set[T] {
	self := Set[T]{
		records: []MapRecordItem[int, T]{},
		size:    0,
	}

	for _, item := range base {
		self.Add(item)
	}

	return &self
}

func (self *Set[T]) findIndex(item T) int {
	return slices.IndexFunc(self.records, func(record MapRecordItem[int, T]) bool {
		return record.Value == item && !record.Deleted
	})
}

// Adds an item to the set. If the item already exists, the set remains untouched.
func (self *Set[T]) Add(item T) *Set[T] {
	if !self.Has(item) {
		self.records = append(self.records, MapRecordItem[int, T]{
			Key:     self.size,
			Value:   item,
			Deleted: false,
		})
		self.size++
	}

	return self
}

// Checks if the given item exists in the set.
func (self *Set[T]) Has(item T) bool {
	idx := self.findIndex(item)
	return idx != -1
}

// Removes the item from the set.
func (self *Set[T]) Delete(item T) bool {
	idx := self.findIndex(item)

	if idx == -1 {
		return false
	}

	record := &self.records[idx]
	record.Key = 0
	record.Value = *new(T)
	record.Deleted = true
	self.size--

	return true
}

// Empties the set and reset its size.
func (self *Set[T]) Clear() {
	self.records = []MapRecordItem[int, T]{}
	self.size = 0
}

// Retrieves all the values in the set.
func (self *Set[T]) Values() []T {
	items := make([]T, self.size)
	idx := 0

	for _, record := range self.records {
		if !record.Deleted {
			items[idx] = record.Value
			idx++
		}
	}

	return items
}

// Loop through all the items in the set and invoke the given function against them.
func (self *Set[T]) ForEach(fn func(item T)) {
	for _, record := range self.records {
		if !record.Deleted {
			fn(record.Value)
		}
	}
}

// Returns the size of the set.
func (self *Set[T]) Size() int {
	return self.size
}
