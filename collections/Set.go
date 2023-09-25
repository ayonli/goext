package collections

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ayonli/goext/number"
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
	self.m.Use(item, func() int { return self.m.size })
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

// Removes and returns the first (if `order >= 0`) or the last (if `order < 0`) item from the set.
func (self *Set[T]) Pop(order int) (T, bool) {
	if self.Size() == 0 {
		return *new(T), false
	}

	self.m.mut.Lock()
	defer self.m.mut.Unlock()

	if order >= 0 {
		for idx, item := range self.m.records {
			if !item.Deleted {
				self.m.deleteAt(idx)
				return item.Key, true
			}
		}
	} else {
		for idx := len(self.m.records) - 1; idx >= 0; idx-- {
			item := self.m.records[idx]

			if !item.Deleted {
				self.m.deleteAt(idx)
				return item.Key, true
			}
		}
	}

	// never
	return *new(T), false
}

// Removes and returns a random item from the set.
func (self *Set[T]) Random() (T, bool) {
	if self.Size() == 0 {
		return *new(T), false
	}

	self.m.mut.Lock()
	defer self.m.mut.Unlock()

	pos := number.Random(0, self.Size()-1)
	idx := 0

	for i, item := range self.m.records {
		if !item.Deleted {
			if idx == pos {
				self.m.deleteAt(i)
				return item.Key, true
			}

			idx++
		}
	}

	// never
	return *new(T), true
}

// Empties the set and resets its size.
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
