package collections_test

import (
	"encoding/json"
	"fmt"

	"github.com/ayonli/goext/collections"
)

func ExampleBiMap() {
	m := &collections.BiMap[string, string]{} // use & for literal creation
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m)
	// Output:
	// &collections.BiMap[foo:Hello bar:World]
}

func ExampleBiMap_json() {
	m := &collections.BiMap[string, string]{}
	m.Set("foo", "Hello").Set("bar", "World")

	data, _ := json.Marshal(m)
	fmt.Println(string(data))

	m2 := &collections.BiMap[string, string]{}
	json.Unmarshal(data, m2)
	fmt.Println(m2)
	// Output:
	// {"foo":"Hello","bar":"World"}
	// &collections.BiMap[bar:World foo:Hello]
}

func ExampleNewBiMap() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m)
	// Output:
	// &collections.BiMap[foo:Hello bar:World]
}

func ExampleBiMap_Set() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{})
	m.Set("foo", "Hello").Set("bar", "World") // Set() method can be chained

	fmt.Println(m) // keys' order is preserved
	fmt.Printf("%#v", m)
	// Output:
	// &collections.BiMap[foo:Hello bar:World]
	// &collections.BiMap[string, string]{"foo":"Hello", "bar":"World"}
}

func ExampleBiMap_Get() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("foo1"))
	// Output:
	// Hello true
	//  false
}

func ExampleBiMap_GetKey() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.GetKey("Hello"))
	fmt.Println(m.GetKey("Hi"))
	// Output:
	// foo true
	//  false
}

func ExampleBiMap_Has() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Has("foo"))
	fmt.Println(m.Has("foo1"))
	// Output:
	// true
	// false
}

func ExampleBiMap_HasValue() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.HasValue("Hello"))
	fmt.Println(m.HasValue("Hi"))
	// Output:
	// true
	// false
}

func ExampleBiMap_Delete() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	ok1 := m.Delete("foo") // succeed
	ok2 := m.Delete("foo") // fail

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &collections.BiMap[bar:World]
	// true
	// false
}

func ExampleBiMap_DeleteValue() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	ok1 := m.DeleteValue("Hello") // succeed
	ok2 := m.DeleteValue("Hello") // failed

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &collections.BiMap[bar:World]
	// true
	// false
}

func ExampleBiMap_Clear() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	m.Clear()

	fmt.Println(m)
	// Output:
	// &collections.BiMap[]
}

func ExampleBiMap_Keys() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Keys()) // keys' order is preserved
	// Output:
	// [foo bar]
}

func ExampleBiMap_Values() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.Values()) // values' order is the same as keys'
	// Output:
	// [Hello World]
}

func ExampleBiMap_Entries() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
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

func ExampleBiMap_ForEach() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
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

func ExampleBiMap_Size() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{})
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

func ExampleBiMap_ToMap() {
	m := collections.NewBiMap([]collections.MapEntry[string, string]{
		{"foo", "Hello"},
		{"bar", "World"},
	})

	fmt.Println(m.ToMap()) // the printed representation is ordered alphabetically, but the real value is not
	// Output:
	// map[bar:World foo:Hello]
}
