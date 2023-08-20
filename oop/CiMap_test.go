package oop

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCiMap(suit *testing.T) {
	suit.Run("NewMap", func(t *testing.T) {
		m := NewCiMap[string, string]()

		assert.Equal(t, []CiMapRecordItem[string, string]{}, m.records)
		assert.Equal(t, 0, m.size)
	})

	suit.Run("Set", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, []CiMapRecordItem[string, string]{
			{Id: "foo", Key: "foo", Value: "Hello", Deleted: false},
			{Id: "bar", Key: "bar", Value: "World", Deleted: false},
		}, m.records)
		assert.Equal(t, 2, m.size)

		m.Set("Foo", "Hi")
		assert.Equal(t, []CiMapRecordItem[string, string]{
			{Id: "foo", Key: "Foo", Value: "Hi", Deleted: false},
			{Id: "bar", Key: "bar", Value: "World", Deleted: false},
		}, m.records)
	})

	suit.Run("Get", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		v1, ok1 := m.Get("foo")
		v2, ok2 := m.Get("Bar")
		v3, ok3 := m.Get("foo1")

		assert.Equal(t, "Hello", v1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, "World", v2)
		assert.Equal(t, true, ok2)
		assert.Equal(t, "", v3)
		assert.Equal(t, false, ok3)
	})

	suit.Run("Has", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, true, m.Has("foo"))
		assert.Equal(t, true, m.Has("Foo"))
		assert.Equal(t, false, m.Has("bar1"))
	})

	suit.Run("Delete", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")
		ok1 := m.Delete("foo")
		ok2 := m.Delete("Bar")
		ok3 := m.Delete("bar")

		assert.Equal(t, true, ok1)
		assert.Equal(t, true, ok2)
		assert.Equal(t, false, ok3)
		assert.Equal(t, 0, m.size)
		assert.Equal(t, []CiMapRecordItem[string, string]{
			{Id: "", Key: "", Value: "", Deleted: true},
			{Id: "", Key: "", Value: "", Deleted: true},
		}, m.records)

		m2 := NewCiMap[string, string]()

		for i := 0; i < 100; i++ {
			m2.Set(strconv.Itoa(i), strconv.Itoa(i))
		}

		for i := 0; i < 100; i++ {
			m2.Delete(strconv.Itoa(i))
		}

		assert.Equal(t, 0, m2.Size())
		assert.Equal(t, 33, len(m2.records))
	})

	suit.Run("Clear", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")
		m.Clear()

		assert.Equal(t, 0, m.size)
		assert.Equal(t, []CiMapRecordItem[string, string]{}, m.records)
	})

	suit.Run("Keys", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("Foo", "Hello").Set("Bar", "World")

		assert.Equal(t, []string{"Foo", "Bar"}, m.Keys())
	})

	suit.Run("Values", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, []string{"Hello", "World"}, m.Values())
	})

	suit.Run("ToMap", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, map[string]string{
			"foo": "Hello",
			"bar": "World",
		}, m.ToMap())
	})

	suit.Run("ForEach", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("Bar", "World")
		entries := List[*[]string]{}

		m.ForEach(func(value string, key string) {
			entries.Push(&[]string{key, value})
		})

		assert.Equal(t, []*[]string{
			{"foo", "Hello"},
			{"Bar", "World"},
		}, entries.Values())
	})

	suit.Run("Size", func(t *testing.T) {
		m := NewCiMap[string, string]()
		m.Set("foo", "Hello").Set("Bar", "World")

		assert.Equal(t, 2, m.Size())

		m.Delete("Foo")

		assert.Equal(t, 1, m.Size())
		assert.Equal(t, 1, len(m.Keys()))
		assert.Equal(t, 1, len(m.Values()))
		assert.Equal(t, []CiMapRecordItem[string, string]{
			{Id: "", Key: "", Value: "", Deleted: true},
			{Id: "bar", Key: "Bar", Value: "World", Deleted: false},
		}, m.records)
	})
}
