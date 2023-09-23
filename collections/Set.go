package collections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Set is an object-oriented collection that stores unique items and is thread-safe.
type Set[T comparable] struct {
	m Map[T, int]
}

// Creates a new instance of the Set.
func NewSet[T comparable](base []T) *Set[T] {
	self := Set[T]{}

	for _, item := range base {
		if !self.m.Has(item) {
			self.m.set(item, self.m.size)
		}
	}

	return &self
}

// Adds an item to the set. If the item already exists, the set remains untouched.
func (self *Set[T]) Add(item T) *Set[T] {
	self.m.EnsureGet(item, func() int { return self.m.size })
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
	return "&collections.Set" + fmt.Sprint(self.Values())
}

func (self *Set[T]) GoString() string {
	str := fmt.Sprintf("%#v", self.Values())
	idx := strings.Index(str, "{")
	return "&collections.Set[" + str[2:idx] + "]" + str[idx:]
}

func (self *Set[T]) UnmarshalJSON(data []byte) error {
	self.m.mut.Lock()
	defer self.m.mut.Unlock()

	var s []T

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	for _, value := range s {
		if !self.m.Has(value) {
			self.m.set(value, self.m.size)
		}
	}

	return nil
}

func (self *Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Values())
}
