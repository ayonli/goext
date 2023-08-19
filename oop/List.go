package oop

import (
	"slices"

	sliceExt "github.com/ayonli/goext/slices"
)

type List[T comparable] []T

func NewList[T comparable](records []T) *List[T] {
	list := List[T](records)
	return &list
}

func (self *List[T]) At(i int) (T, bool) {
	return sliceExt.At(*self, i)
}

func (self *List[T]) Index(item T) int {
	return slices.Index(*self, item)
}

func (self *List[T]) LastIndex(item T) int {
	return sliceExt.LastIndex(*self, item)
}

func (self *List[T]) Length() int {
	return len([]T(*self))
}

func (self *List[T]) Clone() *List[T] {
	list := slices.Clone(*self)
	return &list
}

func (self *List[T]) Contains(item T) bool {
	return slices.Contains(*self, item)
}

func (self *List[T]) Equal(another List[T]) bool {
	return slices.Equal(*self, another)
}

func (self *List[T]) Count(item T) int {
	return sliceExt.Count(*self, item)
}

func (self *List[T]) Concat(others ...List[T]) *List[T] {
	sources := append([]List[T]{}, *self)
	sources = append(sources, others...)
	list := sliceExt.Concat(sources...)
	return &list
}

func (self *List[T]) Uniq() *List[T] {
	list := sliceExt.Uniq(*self)
	return &list
}

func (self *List[T]) Slice(start int, end int) *List[T] {
	list := sliceExt.Slice(*self, start, end)
	return &list
}

func (self *List[T]) Chunk(length int) []List[T] {
	return sliceExt.Chunk(*self, length)
}

func (self *List[T]) Replace(start int, end int, values ...T) *List[T] {
	limit := self.Length()

	if start < 0 {
		start = limit + start
	}

	if end < 0 {
		end = limit + end
	}

	if end > limit {
		end = limit
	}

	if start >= end || start >= limit {
		return self // nothing can be changed, return self
	}

	list := slices.Replace(*self, start, end, values...)
	return &list
}

func (self *List[T]) Reverse() *List[T] {
	slices.Reverse(*self)
	return self
}

func (self *List[T]) ToReversed() *List[T] {
	return self.Clone().Reverse()
}

func (self *List[T]) Every(fn func(item T, idx int) bool) bool {
	return sliceExt.Every(*self, fn)
}

func (self *List[T]) Some(fn func(item T, idx int) bool) bool {
	return sliceExt.Some(*self, fn)
}

func (self *List[T]) Find(fn func(item T, idx int) bool) (T, bool) {
	return sliceExt.Find(*self, fn)
}

func (self *List[T]) FindLast(fn func(item T, idx int) bool) (T, bool) {
	return sliceExt.FindLast(*self, fn)
}

func (self *List[T]) Filter(fn func(item T, idx int) bool) *List[T] {
	list := sliceExt.Filter(*self, fn)
	return &list
}

func (self *List[T]) Pop() T {
	return sliceExt.Pop(self)
}

func (self *List[T]) Push(items ...T) int {
	return sliceExt.Push(self, items...)
}

func (self *List[T]) Shift() T {
	return sliceExt.Shift(self)
}

func (self *List[T]) Unshift(items ...T) int {
	return sliceExt.Unshift(self, items...)
}

func (self *List[T]) Shuffle() *List[T] {
	sliceExt.Shuffle(*self)
	return self
}
