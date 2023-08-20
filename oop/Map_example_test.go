package oop_test

import (
	"fmt"

	"github.com/ayonli/goext/oop"
)

func ExampleNewMap() {
	m := oop.NewMap[string, string]()

	fmt.Println(m)
	// Output:
	// Map[]
}

func ExampleMap_Set() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World") // Set() method can be chained

	fmt.Println(m) // keys' order is preserved
	// Output:
	// Map[foo:Hello bar:World]
}

func ExampleMap_Get() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("foo1"))
	// Output:
	// Hello true
	//  false
}

func ExampleMap_Has() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Has("foo"))
	fmt.Println(m.Has("foo1"))
	// Output:
	// true
	// false
}

func ExampleMap_Delete() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	ok1 := m.Delete("foo") // succeed
	ok2 := m.Delete("foo") // failed

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// Map[bar:World]
	// true
	// false
}

func ExampleMap_Clear() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	m.Clear()

	fmt.Println(m)
	// Output:
	// Map[]
}

func ExampleMap_Keys() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Keys()) // keys' order is preserved
	// Output:
	// [foo bar]
}

func ExampleMap_Values() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Values()) // values' order is the same with keys'
	// Output:
	// [Hello World]
}

func ExampleMap_ToMap() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.ToMap()) // the printed representation is ordered alphabetically, but the real value is not
	// Output:
	// map[bar:World foo:Hello]
}

func ExampleMap_ForEach() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	m.ForEach(func(value string, key string) {
		fmt.Println(key + " => " + value)
	})
	// Output:
	// foo => Hello
	// bar => World
}

func ExampleMap_Size() {
	m := oop.NewMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Size())
	// Output:
	// 2
}