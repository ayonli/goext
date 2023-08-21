package oop_test

import (
	"fmt"

	"github.com/ayonli/goext/oop"
)

func ExampleNewCiMap() {
	m := oop.NewCiMap[string, string]()

	fmt.Println(m)
	fmt.Printf("%#v\n", m)
	// Output:
	// &oop.CiMap[]
	// &oop.CiMap[string, string]{}
}

func ExampleCiMap_Set() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World") // Set() method can be chained

	fmt.Println(m) // keys' names and their order are preserved
	fmt.Printf("%#v\n", m)
	// Output:
	// &oop.CiMap[Foo:Hello bar:World]
	// &oop.CiMap[string, string]{"Foo":"Hello", "bar":"World"}
}

func ExampleCiMap_Get() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("foo1"))
	// Output:
	// Hello true
	//  false
}

func ExampleCiMap_Has() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Has("foo"))
	fmt.Println(m.Has("foo1"))
	// Output:
	// true
	// false
}

func ExampleCiMap_Delete() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	ok1 := m.Delete("foo") // succeed
	ok2 := m.Delete("foo") // failed

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &oop.CiMap[bar:World]
	// true
	// false
}

func ExampleCiMap_Clear() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	m.Clear()

	fmt.Println(m)
	// Output:
	// &oop.CiMap[]
}

func ExampleCiMap_Keys() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Keys()) // keys' names and their order are preserved
	// Output:
	// [Foo bar]
}

func ExampleCiMap_Values() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Values()) // values' order is the same with keys'
	// Output:
	// [Hello World]
}

func ExampleCiMap_ToMap() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.ToMap()) // the printed representation is ordered alphabetically, but the real value is not
	// Output:
	// map[Foo:Hello bar:World]
}

func ExampleCiMap_ForEach() {
	m := oop.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	m.ForEach(func(value string, key string) {
		fmt.Println(key + " => " + value)
	})
	// Output:
	// Foo => Hello
	// bar => World
}

func ExampleCiMap_Size() {
	m := oop.NewCiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Size())
	// Output:
	// 2
}
