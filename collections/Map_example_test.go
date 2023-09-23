package collections_test

import (
	"encoding/json"
	"fmt"

	"github.com/ayonli/goext/collections"
)

func ExampleMap() {
	m := &collections.Map[string, string]{} // use & for literal creation
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m)
	// Output:
	// &collections.Map[foo:Hello bar:World]
}

func ExampleMap_json() {
	m := &collections.Map[string, string]{}
	m.Set("foo", "Hello").Set("bar", "World")

	data, _ := json.Marshal(m)
	fmt.Println(string(data))

	m2 := &collections.Map[string, string]{}
	json.Unmarshal(data, m2)
	fmt.Println(m2)
	// Output:
	// {"foo":"Hello","bar":"World"}
	// &collections.Map[bar:World foo:Hello]
}

func ExampleNewMap() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m)
	// Output:
	// &collections.Map[foo:Hello bar:World]
}

func ExampleMap_Set() {
	m := collections.NewMap([]collections.MapEntry[string, string]{})
	m.Set("foo", "Hello").Set("bar", "World") // Set() method can be chained

	fmt.Println(m) // keys' order is preserved
	fmt.Printf("%#v\n", m)
	// Output:
	// &collections.Map[foo:Hello bar:World]
	// &collections.Map[string, string]{"foo":"Hello", "bar":"World"}
}

func ExampleMap_Get() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("foo1"))
	// Output:
	// Hello true
	//  false
}

func ExampleMap_Has() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Has("foo"))
	fmt.Println(m.Has("foo1"))
	// Output:
	// true
	// false
}

func ExampleMap_EnsureGet() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hi"},
	})

	val1 := m.EnsureGet("foo", func() string { return "Hello" })
	val2 := m.EnsureGet("bar", func() string { return "World" })

	fmt.Println(val1)
	fmt.Println(val2)
	// Output:
	// Hi
	// World
}

func ExampleMap_Delete() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	ok1 := m.Delete("foo") // succeed
	ok2 := m.Delete("foo") // fail

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &collections.Map[bar:World]
	// true
	// false
}

func ExampleMap_Clear() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	m.Clear()

	fmt.Println(m)
	// Output:
	// &collections.Map[]
}

func ExampleMap_Keys() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Keys()) // keys' order is preserved
	// Output:
	// [foo bar]
}

func ExampleMap_Values() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Values()) // values' order is the same as keys'
	// Output:
	// [Hello World]
}

func ExampleMap_Entries() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	for entry := range m.Entries() {
		fmt.Println(entry.Key, "=>", entry.Value)
	}
	// Output:
	// foo => Hello
	// bar => World
}

func ExampleMap_ForEach() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	m.ForEach(func(value string, key string) {
		fmt.Println(key, "=>", value)
	})
	// Output:
	// foo => Hello
	// bar => World
}

func ExampleMap_Size() {
	m := collections.NewMap([]collections.MapEntry[string, string]{})
	fmt.Println(m.Size())

	m.Set("foo", "Hello")
	fmt.Println(m.Size())

	m.Set("bar", "World")
	fmt.Println(m.Size())
	// Output:
	// 0
	// 1
	// 2
}

func ExampleMap_ToMap() {
	m := collections.NewMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.ToMap()) // the printed representation is ordered alphabetically, but the real value is not
	// Output:
	// map[bar:World foo:Hello]
}
