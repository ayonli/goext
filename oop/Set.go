package oop

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Set is an object-oriented collection that stores unique items.
type Set[T comparable] struct {
	m Map[T, int]
}

// Creates a new instance of the Set.
func NewSet[T comparable](base []T) *Set[T] {
	self := Set[T]{}

	for _, item := range base {
		self.Add(item)
	}

	return &self
}

// Adds an item to the set. If the item already exists, the set remains untouched.
func (self *Set[T]) Add(item T) *Set[T] {
	if !self.Has(item) {
		self.m.Set(item, self.m.size)
	}

	return self
}

// Checks if the given item exists in the set.
func (self *Set[T]) Has(item T) bool {
	return self.m.Has(item)
}

// Removes the item from the set.
func (self *Set[T]) Delete(item T) bool {
	return self.m.Delete(item)
}

func (self *Set[T]) Clear() {
	self.m.Clear()
}

// Retrieves all the values in the set.
func (self *Set[T]) Values() []T {
	return self.m.Keys()
}

// Loop through all the items in the set and invoke the given function against them.
func (self *Set[T]) ForEach(fn func(item T)) {
	self.m.ForEach(func(_ int, key T) {
		fn(key)
	})
}

// Returns the size of the set.
func (self *Set[T]) Size() int {
	return self.m.size
}

func (self *Set[T]) String() string {
	return "&oop.Set" + fmt.Sprint(self.Values())
}

func (self *Set[T]) GoString() string {
	str := fmt.Sprintf("%#v", self.Values())
	idx := strings.Index(str, "{")
	return "&oop.Set[" + str[2:idx] + "]" + str[idx:]
}

func (self *Set[T]) UnmarshalJSON(data []byte) error {
	var s []T

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	for _, value := range s {
		self.Add(value)
	}

	return nil
}

func (self Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Values())
}
