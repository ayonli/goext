package collections_test

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/ayonli/goext/collections"
)

func ExampleSet() {
	s := &collections.Set[string]{} // use & for literal creation
	s.Add("Hello").Add("World")

	fmt.Println(s)
	// Output:
	// &collections.Set[Hello World]
}

func ExampleSet_json() {
	s := &collections.Set[string]{} // use & for literal creation
	s.Add("Hello").Add("World")

	data, _ := json.Marshal(s)
	fmt.Println(string(data))

	s2 := &collections.Set[string]{}
	json.Unmarshal(data, s2)
	fmt.Printf("%#v", s2)
	// Output:
	// ["Hello","World"]
	// &collections.Set[string]{"Hello", "World"}
}

func ExampleNewSet() {
	s := collections.NewSet([]string{
		"Hello",
		"World",
		"Hello",
		"A-yon",
	})

	fmt.Println(s)
	// Output:
	// &collections.Set[Hello World A-yon]
}

func ExampleSet_Add() {
	s := collections.NewSet([]string{})

	s.Add("Hello").Add("World") // Add() method can be chained
	s.Add("Hello")              // duplicate adding values will not effect

	fmt.Println(s)
	// Output:
	// &collections.Set[Hello World]
}

func ExampleSet_Has() {
	s := collections.NewSet([]string{"Hello"})

	fmt.Println(s.Has("Hello"))
	fmt.Println(s.Has("World"))
	// Output:
	// true
	// false
}

func ExampleSet_Delete() {
	s := collections.NewSet([]string{"Hello", "World"})

	ok1 := s.Delete("Hello") // succeed
	ok2 := s.Delete("Hello") // failed

	fmt.Println(s)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &collections.Set[World]
	// true
	// false
}

func ExampleSet_Pop() {
	s := collections.NewSet([]string{
		"Hello",
		"World",
		"Hi",
		"A-yon",
	})

	fmt.Println(s.Pop(0))
	fmt.Println(s.Pop(-1))
	fmt.Println(s.Size())

	s.Clear()
	fmt.Println(s.Pop(0))

	// Output:
	// Hello true
	// A-yon true
	// 2
	//  false
}

func ExampleSet_Random() {
	list := []string{
		"Hello",
		"World",
		"Hi",
		"A-yon",
	}
	s := collections.NewSet(list)

	item, _ := s.Random()
	fmt.Println(slices.Contains(list, item))
	fmt.Println(s.Size())

	s.Clear()
	fmt.Println(s.Random())

	// Output:
	// true
	// 3
	//  false
}

func ExampleSet_Clear() {
	s := collections.NewSet([]string{"Hello", "World"})

	s.Clear()

	fmt.Println(s)
	// Output:
	// &collections.Set[]
}

func ExampleSet_Values() {
	s := collections.NewSet([]string{"Hello", "World"})

	fmt.Println(s.Values())
	// Output:
	// [Hello World]
}

func ExampleSet_ForEach() {
	s := collections.NewSet([]string{"Hello", "World"})

	s.ForEach(func(item string) {
		fmt.Println(item)
	})
	// Output:
	// Hello
	// World
}

func ExampleSet_Size() {
	s := collections.NewSet([]string{"Hello", "World"})

	fmt.Println(s.Size())
	// Output:
	// 2
}
