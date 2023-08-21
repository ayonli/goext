package oop_test

import (
	"fmt"

	"github.com/ayonli/goext/oop"
)

func ExampleNewBiMap() {
	m := oop.NewBiMap[string, string]()

	fmt.Println(m)
	fmt.Printf("%#v\n", m)
	// Output:
	// &oop.BiMap[]
	// &oop.BiMap[string, string]{}
}

func ExampleBiMap_Set() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World") // Set() method can be chained

	fmt.Println(m) // keys' order is preserved
	fmt.Printf("%#v", m)
	// Output:
	// &oop.BiMap[foo:Hello bar:World]
	// &oop.BiMap[string, string]{"foo":"Hello", "bar":"World"}
}

func ExampleBiMap_Get() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("foo1"))
	// Output:
	// Hello true
	//  false
}

func ExampleBiMap_GetKey() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.GetKey("Hello"))
	fmt.Println(m.GetKey("Hi"))
	// Output:
	// foo true
	//  false
}

func ExampleBiMap_Has() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Has("foo"))
	fmt.Println(m.Has("foo1"))
	// Output:
	// true
	// false
}

func ExampleBiMap_HasValue() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.HasValue("Hello"))
	fmt.Println(m.HasValue("Hi"))
	// Output:
	// true
	// false
}

func ExampleBiMap_Delete() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	ok1 := m.Delete("foo") // succeed
	ok2 := m.Delete("foo") // failed

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &oop.BiMap[bar:World]
	// true
	// false
}

func ExampleBiMap_DeleteValue() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	ok1 := m.DeleteValue("Hello") // succeed
	ok2 := m.DeleteValue("Hello") // failed

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &oop.BiMap[bar:World]
	// true
	// false
}

func ExampleBiMap_Clear() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	m.Clear()

	fmt.Println(m)
	// Output:
	// &oop.BiMap[]
}

func ExampleBiMap_Keys() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Keys()) // keys' order is preserved
	// Output:
	// [foo bar]
}

func ExampleBiMap_Values() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Values()) // values' order is the same with keys'
	// Output:
	// [Hello World]
}

func ExampleBiMap_ToMap() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.ToMap()) // the printed representation is ordered alphabetically, but the real value is not
	// Output:
	// map[bar:World foo:Hello]
}

func ExampleBiMap_ForEach() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	m.ForEach(func(value string, key string) {
		fmt.Println(key + " => " + value)
	})
	// Output:
	// foo => Hello
	// bar => World
}

func ExampleBiMap_Size() {
	m := oop.NewBiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Size())
	// Output:
	// 2
}
