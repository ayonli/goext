package oop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(suit *testing.T) {
	suit.Run("NewSet", func(t *testing.T) {
		s1 := NewSet([]string{})
		s2 := NewSet([]string{
			"Hello",
			"World",
			"Hello",
			"A-yon",
		})

		assert.Equal(t, []mapRecordItem[int, string]{}, s1.records)
		assert.Equal(t, 0, s1.size)
		assert.Equal(t, []mapRecordItem[int, string]{
			{Key: 0, Value: "Hello", Deleted: false},
			{Key: 1, Value: "World", Deleted: false},
			{Key: 2, Value: "A-yon", Deleted: false},
		}, s2.records)
		assert.Equal(t, 3, s2.size)
	})

	suit.Run("Add", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World").Add("Hello")

		assert.Equal(t, []mapRecordItem[int, string]{
			{Key: 0, Value: "Hello", Deleted: false},
			{Key: 1, Value: "World", Deleted: false},
		}, s.records)
		assert.Equal(t, 2, s.size)
	})

	suit.Run("Has", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World")

		assert.Equal(t, true, s.Has("Hello"))
		assert.Equal(t, false, s.Has("World1"))
	})

	suit.Run("Delete", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World")
		ok1 := s.Delete("Hello")
		ok2 := s.Delete("Hola")

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
		assert.Equal(t, 1, s.size)
		assert.Equal(t, []mapRecordItem[int, string]{
			{Key: 0, Value: "", Deleted: true},
			{Key: 1, Value: "World", Deleted: false},
		}, s.records)

		m2 := NewSet([]int{})

		for i := 0; i < 100; i++ {
			m2.Add(i)
		}

		for i := 0; i < 100; i++ {
			m2.Delete(i)
		}

		assert.Equal(t, 0, m2.Size())
		assert.Equal(t, 33, len(m2.records))
	})

	suit.Run("Clear", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World")
		s.Clear()

		assert.Equal(t, 0, s.size)
		assert.Equal(t, []mapRecordItem[int, string]{}, s.records)
	})

	suit.Run("Values", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World")

		assert.Equal(t, []string{"Hello", "World"}, s.Values())
	})

	suit.Run("ForEach", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World")
		entries := &List[string]{}

		s.ForEach(func(value string) {
			entries.Push(value)
		})

		assert.Equal(t, []string{"Hello", "World"}, entries.Values())
	})

	suit.Run("Size", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World")

		assert.Equal(t, 2, s.Size())

		s.Delete("Hello")

		assert.Equal(t, 1, s.Size())
		assert.Equal(t, 1, len(s.Values()))
		assert.Equal(t, []mapRecordItem[int, string]{
			{Key: 0, Value: "", Deleted: true},
			{Key: 1, Value: "World", Deleted: false},
		}, s.records)
	})
}
