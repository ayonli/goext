package collections

import (
	"strconv"
	"testing"

	"github.com/ayonli/goext/oop"
	"github.com/stretchr/testify/assert"
)

func TestMap(suit *testing.T) {
	suit.Run("NewMap", func(t *testing.T) {
		m := NewMap[string, string]([]MapEntry[string, string]{})

		assert.Equal(t, []mapRecordItem[string, string](nil), m.records)
		assert.Equal(t, 0, m.size)
	})

	suit.Run("Set", func(t *testing.T) {
		m := NewMap[string, string]([]MapEntry[string, string]{})
		m.Set("foo", "Hello").Set("bar", "World")

		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "foo", Value: "Hello", Deleted: false},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
		assert.Equal(t, 2, m.size)

		m.Set("foo", "Hi")
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "foo", Value: "Hi", Deleted: false},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
	})

	suit.Run("Get", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		v1, ok1 := m.Get("foo")
		v2, ok2 := m.Get("bar1")

		assert.Equal(t, "Hello", v1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, "", v2)
		assert.Equal(t, false, ok2)
	})

	suit.Run("Has", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, true, m.Has("foo"))
		assert.Equal(t, false, m.Has("bar1"))
	})

	suit.Run("Delete", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		ok1 := m.Delete("foo")
		ok2 := m.Delete("bar1")

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
		assert.Equal(t, 1, m.size)
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "", Value: "", Deleted: true},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)

		m2 := NewMap([]MapEntry[int, string]{})

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
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})
		m.Clear()

		assert.Equal(t, 0, m.size)
		assert.Equal(t, []mapRecordItem[string, string](nil), m.records)
	})

	suit.Run("Keys", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, []string{"foo", "bar"}, m.Keys())
	})

	suit.Run("Values", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, []string{"Hello", "World"}, m.Values())
	})

	suit.Run("ForEach", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		entries := &oop.List[*[]string]{}

		m.ForEach(func(value string, key string) {
			entries.Push(&[]string{key, value})
		})

		assert.Equal(t, []*[]string{
			{"foo", "Hello"},
			{"bar", "World"},
		}, entries.Values())
	})

	suit.Run("Size", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, 2, m.Size())

		m.Delete("foo")

		assert.Equal(t, 1, m.Size())
		assert.Equal(t, 1, len(m.Keys()))
		assert.Equal(t, 1, len(m.Values()))
		assert.Equal(t, []mapRecordItem[string, string]{
			{Key: "", Value: "", Deleted: true},
			{Key: "bar", Value: "World", Deleted: false},
		}, m.records)
	})

	suit.Run("ToMap", func(t *testing.T) {
		m := NewMap([]MapEntry[string, string]{
			{"foo", "Hello"},
			{"bar", "World"},
		})

		assert.Equal(t, map[string]string{
			"foo": "Hello",
			"bar": "World",
		}, m.ToMap())
	})
}
