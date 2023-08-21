package oop

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiMap(suit *testing.T) {
	suit.Run("NewBiMap", func(t *testing.T) {
		m := NewBiMap[string, string]()

		assert.Equal(t, []mapRecordItem[string, string](nil), m.records)
		assert.Equal(t, 0, m.size)
	})

	suit.Run("Set", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "foo", Value: "Hello", Deleted: false},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
		assert.Equal(t, 2, m.size)

		m.Set("foo1", "Hello")
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "foo1", Value: "Hello", Deleted: false},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
	})

	suit.Run("Get", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		v1, ok1 := m.Get("foo")
		v2, ok2 := m.Get("bar1")

		assert.Equal(t, "Hello", v1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, "", v2)
		assert.Equal(t, false, ok2)
	})

	suit.Run("GetKey", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		v1, ok1 := m.GetKey("Hello")
		v2, ok2 := m.GetKey("World1")

		assert.Equal(t, "foo", v1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, "", v2)
		assert.Equal(t, false, ok2)
	})

	suit.Run("Has", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, true, m.Has("foo"))
		assert.Equal(t, false, m.Has("bar1"))
	})

	suit.Run("HasValue", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, true, m.HasValue("Hello"))
		assert.Equal(t, false, m.HasValue("World1"))
	})

	suit.Run("Delete", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")
		ok1 := m.Delete("foo")
		ok2 := m.Delete("bar1")

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
		assert.Equal(t, 1, m.size)
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "", Value: "", Deleted: true},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
	})

	suit.Run("DeleteValue", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")
		ok1 := m.DeleteValue("Hello")
		ok2 := m.DeleteValue("World1")

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
		assert.Equal(t, 1, m.size)
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "", Value: "", Deleted: true},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)

		m2 := NewBiMap[int, string]()

		for i := 0; i < 100; i++ {
			m2.Set(i, strconv.Itoa(i))
		}

		for i := 0; i < 100; i++ {
			m2.Delete(i)
		}

		assert.Equal(t, 0, m2.Size())
		assert.Equal(t, 33, len(m2.records))
	})

	suit.Run("Clear", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")
		m.Clear()

		assert.Equal(t, 0, m.size)
		assert.Equal(t, []mapRecordItem[string, string](nil), m.records)
	})

	suit.Run("Keys", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, []string{"foo", "bar"}, m.Keys())
	})

	suit.Run("Values", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, []string{"Hello", "World"}, m.Values())
	})

	suit.Run("ToMap", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, map[string]string{
			"foo": "Hello",
			"bar": "World",
		}, m.ToMap())
	})

	suit.Run("ForEach", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")
		entries := &List[*[]string]{}

		m.ForEach(func(value string, key string) {
			entries.Push(&[]string{key, value})
		})

		assert.Equal(t, []*[]string{
			{"foo", "Hello"},
			{"bar", "World"},
		}, entries.Values())
	})

	suit.Run("Size", func(t *testing.T) {
		m := NewBiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, 2, m.Size())

		m.Delete("foo")
		m.DeleteValue("World")

		assert.Equal(t, 0, m.Size())
		assert.Equal(t, 0, len(m.Keys()))
		assert.Equal(t, 0, len(m.Values()))
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "", Value: "", Deleted: true},
			{Key: "", Value: "", Deleted: true},
		}, m.records)
	})
}
