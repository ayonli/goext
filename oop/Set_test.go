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

		assert.Equal(t, []mapRecordItem[string, int](nil), s1.m.records)
		assert.Equal(t, 0, s1.m.size)
		assert.Equal(t, []mapRecordItem[string, int]{
			{Key: "Hello", Value: 0, Deleted: false},
			{Key: "World", Value: 1, Deleted: false},
			{Key: "A-yon", Value: 2, Deleted: false},
		}, s2.m.records)
		assert.Equal(t, 3, s2.m.size)
	})

	suit.Run("Add", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World").Add("Hello")

		assert.Equal(t, []mapRecordItem[string, int]{
			{Key: "Hello", Value: 0, Deleted: false},
			{Key: "World", Value: 1, Deleted: false},
		}, s.m.records)
		assert.Equal(t, 2, s.m.size)
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
		assert.Equal(t, 1, s.m.size)
		assert.Equal(t, []mapRecordItem[string, int]{
			{Key: "", Value: 0, Deleted: true},
			{Key: "World", Value: 1, Deleted: false},
		}, s.m.records)

		s2 := NewSet([]int{})

		for i := 0; i < 100; i++ {
			s2.Add(i)
		}

		for i := 0; i < 100; i++ {
			s2.Delete(i)
		}

		assert.Equal(t, 0, s2.Size())
		assert.Equal(t, 33, len(s2.m.records))
	})

	suit.Run("Clear", func(t *testing.T) {
		s := NewSet([]string{})
		s.Add("Hello").Add("World")
		s.Clear()

		assert.Equal(t, 0, s.m.size)
		assert.Equal(t, []mapRecordItem[string, int](nil), s.m.records)
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
		assert.Equal(t, []mapRecordItem[string, int]{
			{Key: "", Value: 0, Deleted: true},
			{Key: "World", Value: 1, Deleted: false},
		}, s.m.records)
	})
}
