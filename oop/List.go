package oop

import (
	"fmt"
	"slices"

	"github.com/ayonli/goext/slicex"
)

// List is an objected-oriented abstract that works around the slice.
type List[T comparable] []T

func NewList[T comparable](base []T) *List[T] {
	ins := List[T](base)
	return &ins
}

func (self *List[T]) At(i int) (T, bool) {
	return slicex.At(*self, i)
}

func (self *List[T]) IndexOf(item T) int {
	return slices.Index(*self, item)
}

func (self *List[T]) LastIndexOf(item T) int {
	return slicex.LastIndex(*self, item)
}

func (self *List[T]) Length() int {
	return len(*self)
}

func (self *List[T]) Values() []T {
	return []T(*self)
}

func (self *List[T]) String() string {
	return "&" + fmt.Sprint(*self)
}

func (self *List[T]) GoString() string {
	return "&" + fmt.Sprintf("%#v", *self)
}

func (self *List[T]) Clone() *List[T] {
	list := slices.Clone(*self)
	return &list
}

func (self *List[T]) Equal(another *List[T]) bool {
	return slices.Equal(*self, *another)
}

func (self *List[T]) Contains(item T) bool {
	return slices.Contains(*self, item)
}

func (self *List[T]) Count(item T) int {
	return slicex.Count(*self, item)
}

func (self *List[T]) valuedWithOthers(others []*List[T]) []List[T] {
	sources := append([]List[T]{}, *self)
	return append(sources, slicex.Map(others, func(list *List[T], _ int) List[T] {
		if list == nil {
			return nil
		} else {
			return *list
		}
	})...)
}

func (self *List[T]) Concat(others ...*List[T]) *List[T] {
	list := slicex.Concat(self.valuedWithOthers(others)...)
	return &list
}

func (self *List[T]) Uniq() *List[T] {
	list := slicex.Uniq(*self)
	return &list
}

func (self *List[T]) Slice(start int, end int) *List[T] {
	list := slicex.Slice(*self, start, end)
	return &list
}

func (self *List[T]) Chunk(length int) []*List[T] { // TODO
	return slicex.Map(slicex.Chunk(*self, length), func(list List[T], _ int) *List[T] {
		return &list
	})
}

func (self *List[T]) Join(sep string) string {
	return slicex.Join(*self, sep)
}

func (self *List[T]) Replace(start int, end int, values ...T) *List[T] {
	*self = slices.Replace(*self, start, end, values...)
	return self
}

func (self *List[T]) Reverse() *List[T] {
	slices.Reverse(*self)
	return self
}

func (self *List[T]) ToReversed() *List[T] {
	return self.Clone().Reverse()
}

func (self *List[T]) Sort() *List[T] {
	slices.SortStableFunc(*self, slicex.CompareItems)
	return self
}

func (self *List[T]) ToSorted() *List[T] {
	return self.Clone().Sort()
}

func (self *List[T]) Every(fn func(item T, idx int) bool) bool {
	return slicex.Every(*self, fn)
}

func (self *List[T]) Some(fn func(item T, idx int) bool) bool {
	return slicex.Some(*self, fn)
}

func (self *List[T]) Find(fn func(item T, idx int) bool) (T, bool) {
	return slicex.Find(*self, fn)
}

func (self *List[T]) FindLast(fn func(item T, idx int) bool) (T, bool) {
	return slicex.FindLast(*self, fn)
}

func (self *List[T]) FindIndex(fn func(item T, idx int) bool) int {
	return slicex.FindIndex(*self, fn)
}

func (self *List[T]) FindLastIndex(fn func(item T, idx int) bool) int {
	return slicex.FindLastIndex(*self, fn)
}

func (self *List[T]) Filter(fn func(item T, idx int) bool) *List[T] {
	list := slicex.Filter(*self, fn)
	return &list
}

func (self *List[T]) ForEach(fn func(item T, idx int)) *List[T] {
	for idx, item := range *self {
		fn(item, idx)
	}

	return self
}

func (self *List[T]) Pop() T {
	return slicex.Pop(self)
}

func (self *List[T]) Push(items ...T) int {
	return slicex.Push(self, items...)
}

func (self *List[T]) Shift() T {
	return slicex.Shift(self)
}

func (self *List[T]) Unshift(items ...T) int {
	return slicex.Unshift(self, items...)
}

func (self *List[T]) Shuffle() *List[T] {
	slicex.Shuffle(*self)
	return self
}

func (self *List[T]) Diff(others ...*List[T]) *List[T] {
	sources := append([]List[T]{}, slicex.Map(others, func(list *List[T], _ int) List[T] {
		if list == nil {
			return nil
		} else {
			return *list
		}
	})...)
	list := slicex.Diff(*self, sources...)
	return &list
}

func (self *List[T]) Xor(others ...*List[T]) *List[T] {
	list := slicex.Xor(self.valuedWithOthers(others)...)
	return &list
}

func (self *List[T]) Union(others ...*List[T]) *List[T] {
	list := slicex.Union(self.valuedWithOthers(others)...)
	return &list
}

func (self *List[T]) Intersect(others ...*List[T]) *List[T] {
	list := slicex.Intersect(self.valuedWithOthers(others)...)
	return &list
}
