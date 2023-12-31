package collections

import (
	"strconv"
	"testing"

	"github.com/ayonli/goext/oop"
	"github.com/stretchr/testify/assert"
)

func TestCiMap(suit *testing.T) {
	suit.Run("NewMap", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{})

		assert.Equal(t, []mapRecordItem[string, string](nil), m.records)
		assert.Equal(t, []string(nil), m.keys)
		assert.Equal(t, 0, m.size)
	})

	suit.Run("Set", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{})
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "foo", Value: "Hello", Deleted: false},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
		assert.Equal(t, []string{"foo", "bar"}, m.keys)
		assert.Equal(t, 2, m.size)

		m.Set("Foo", "Hi")
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "foo", Value: "Hi", Deleted: false},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
		assert.Equal(t, []string{"Foo", "bar"}, m.keys)
	})

	suit.Run("Get", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

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
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, true, m.Has("foo"))
		assert.Equal(t, true, m.Has("Foo"))
		assert.Equal(t, false, m.Has("bar1"))
	})

	suit.Run("Delete", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		ok1 := m.Delete("foo")
		ok2 := m.Delete("Bar")
		ok3 := m.Delete("bar")

		assert.Equal(t, true, ok1)
		assert.Equal(t, true, ok2)
		assert.Equal(t, false, ok3)
		assert.Equal(t, 0, m.size)
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "", Value: "", Deleted: true},
			{Key: "", Value: "", Deleted: true},
		}, m.records)
		assert.Equal(t, []string{"", ""}, m.keys)

		m2 := NewCiMap([]MapEntry[string, string]{})

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
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		m.Clear()

		assert.Equal(t, 0, m.size)
		assert.Equal(t, []mapRecordItem[string, string](nil), m.records)
		assert.Equal(t, []string(nil), m.keys)
	})

	suit.Run("Keys", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{
			{"Foo", "Hello"},
			{"Bar", "World"},
		})

		assert.Equal(t, []string{"Foo", "Bar"}, m.Keys())
	})

	suit.Run("Values", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, []string{"Hello", "World"}, m.Values())
	})

	suit.Run("ToMap", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, map[string]string{
			"foo": "Hello",
			"bar": "World",
		}, m.ToMap())
	})

	suit.Run("ForEach", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"Bar", "World"},
		})

		entries := &oop.List[*[]string]{}

		m.ForEach(func(value string, key string) {
			entries.Push(&[]string{key, value})
		})

		assert.Equal(t, []*[]string{
			{"foo", "Hello"},
			{"Bar", "World"},
		}, entries.Values())
	})

	suit.Run("Size", func(t *testing.T) {
		m := NewCiMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"Bar", "World"},
		})

		assert.Equal(t, 2, m.Size())

		m.Delete("Foo")

		assert.Equal(t, 1, m.Size())
		assert.Equal(t, 1, len(m.Keys()))
		assert.Equal(t, 1, len(m.Values()))
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "", Value: "", Deleted: true},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
		assert.Equal(t, []string{"", "Bar"}, m.keys)
	})
}
