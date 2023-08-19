package oop

import (
	"slices"
)

type Set[T comparable] struct {
	records []MapRecordItem[int, T]
	size    int
}

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

func (self *Set[T]) Has(item T) bool {
	idx := self.findIndex(item)
	return idx != -1
}

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

func (self *Set[T]) Clear() {
	self.records = []MapRecordItem[int, T]{}
	self.size = 0
}

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

func (self *Set[T]) ForEach(fn func(item T)) {
	for _, record := range self.records {
		if !record.Deleted {
			fn(record.Value)
		}
	}
}

func (self *Set[T]) Size() int {
	return self.size
}
