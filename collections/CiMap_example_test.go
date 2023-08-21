package collections_test

import (
	"encoding/json"
	"fmt"

	"github.com/ayonli/goext/collections"
)

func ExampleCiMap() {
	m := &collections.CiMap[string, string]{} // use & for literal creation
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m)
	fmt.Println(m.Has("foo"))
	fmt.Println(m.Get("foo"))
	// Output:
	// &collections.CiMap[Foo:Hello bar:World]
	// true
	// Hello true
}

func ExampleCiMap_json() {
	m := &collections.CiMap[string, string]{}
	m.Set("Foo", "Hello").Set("bar", "World")

	data, _ := json.Marshal(m)
	fmt.Println(string(data))

	m2 := &collections.CiMap[string, string]{}
	json.Unmarshal(data, m2)
	fmt.Println(m2)
	// Output:
	// {"Foo":"Hello","bar":"World"}
	// &collections.CiMap[Foo:Hello bar:World]
}

func ExampleNewCiMap() {
	m := collections.NewCiMap[string, string]()

	fmt.Println(m)
	fmt.Printf("%#v\n", m)
	// Output:
	// &collections.CiMap[]
	// &collections.CiMap[string, string]{}
}

func ExampleCiMap_Set() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World") // Set() method can be chained

	fmt.Println(m) // keys' names and their order are preserved
	fmt.Printf("%#v\n", m)
	// Output:
	// &collections.CiMap[Foo:Hello bar:World]
	// &collections.CiMap[string, string]{"Foo":"Hello", "bar":"World"}
}

func ExampleCiMap_Get() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("foo1"))
	// Output:
	// Hello true
	//  false
}

func ExampleCiMap_Has() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Has("foo"))
	fmt.Println(m.Has("foo1"))
	// Output:
	// true
	// false
}

func ExampleCiMap_Delete() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	ok1 := m.Delete("foo") // succeed
	ok2 := m.Delete("foo") // failed

	fmt.Println(m)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &collections.CiMap[bar:World]
	// true
	// false
}

func ExampleCiMap_Clear() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	m.Clear()

	fmt.Println(m)
	// Output:
	// &collections.CiMap[]
}

func ExampleCiMap_Keys() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Keys()) // keys' names and their order are preserved
	// Output:
	// [Foo bar]
}

func ExampleCiMap_Values() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.Values()) // values' order is the same as keys'
	// Output:
	// [Hello World]
}

func ExampleCiMap_ToMap() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	fmt.Println(m.ToMap()) // the printed representation is ordered alphabetically, but the real value is not
	// Output:
	// map[Foo:Hello bar:World]
}

func ExampleCiMap_ForEach() {
	m := collections.NewCiMap[string, string]()
	m.Set("Foo", "Hello").Set("bar", "World")

	m.ForEach(func(value string, key string) {
		fmt.Println(key, "=>", value)
	})
	// Output:
	// Foo => Hello
	// bar => World
}

func ExampleCiMap_Size() {
	m := collections.NewCiMap[string, string]()
	m.Set("foo", "Hello").Set("bar", "World")

	fmt.Println(m.Size())
	// Output:
	// 2
}
