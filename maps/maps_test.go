package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssign(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m1 := Assign(map[string]string{}, map[string]string{
			"foo": "Hello",
		})
		m2 := Assign(map[string]string{}, m1, map[string]string{
			"bar": "World",
		})

		assert.Equal(t, map[string]string{
			"foo": "Hello",
		}, m1)
		assert.Equal(t, map[string]string{
			"foo": "Hello",
			"bar": "World",
		}, m2)
	})
}

func TestPatch(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m1 := Patch(map[string]string{}, map[string]string{
			"foo": "Hello",
		})
		m2 := Patch(map[string]string{}, m1, map[string]string{
			"foo": "Hi",
			"bar": "World",
		})

		assert.Equal(t, map[string]string{
			"foo": "Hello",
		}, m1)
		assert.Equal(t, map[string]string{
			"foo": "Hello",
			"bar": "World",
		}, m2)
	})
}

func TestKeys(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m := map[string]string{
			"foo": "Hello",
			"bar": "World",
		}
		keys := Keys(m)

		assert.Equal(t, []string{"bar", "foo"}, keys)
	})

	t.Run("map[int]string", func(t *testing.T) {
		m := map[int]string{
			0: "Hello",
			1: "World",
		}
		keys := Keys(m)

		assert.Equal(t, []int{0, 1}, keys)
	})
}

func TestValues(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m := map[string]string{
			"foo": "Hello",
			"bar": "World",
		}
		values := Values(m)

		assert.Equal(t, []string{"World", "Hello"}, values)
	})

	t.Run("map[int]string", func(t *testing.T) {
		m := map[int]string{
			0: "Hello",
			1: "World",
		}
		values := Values(m)

		assert.Equal(t, []string{"Hello", "World"}, values)
	})
}

func TestPick(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m1 := map[string]string{
			"foo": "Hello",
			"bar": "World",
		}
		m2 := Pick(m1, []string{"foo"})

		assert.Equal(t, map[string]string{
			"foo": "Hello",
		}, m2)
	})

	t.Run("map[int]string", func(t *testing.T) {
		m1 := map[int]string{
			0: "Hello",
			1: "World",
		}
		m2 := Pick(m1, []int{0})

		assert.Equal(t, map[int]string{
			0: "Hello",
		}, m2)
	})
}

func TestOmit(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m1 := map[string]string{
			"foo": "Hello",
			"bar": "World",
		}
		m2 := Omit(m1, []string{"bar"})

		assert.Equal(t, map[string]string{
			"foo": "Hello",
		}, m2)
	})

	t.Run("map[int]string", func(t *testing.T) {
		m1 := map[int]string{
			0: "Hello",
			1: "World",
		}
		m2 := Omit(m1, []int{1})

		assert.Equal(t, map[int]string{
			0: "Hello",
		}, m2)
	})
}
