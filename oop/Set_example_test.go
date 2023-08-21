package oop_test

import (
	"fmt"

	"github.com/ayonli/goext/oop"
)

func ExampleNewSet() {
	s1 := oop.NewSet([]string{}) // create an empty set
	s2 := oop.NewSet([]string{   // create a set with initial values
		"Hello",
		"World",
		"Hello",
		"A-yon",
	})

	fmt.Println(s1)
	fmt.Printf("%#v\n", s2)
	// Output:
	// &oop.Set[]
	// &oop.Set[string]{"Hello", "World", "A-yon"}
}

func ExampleSet_Add() {
	s := oop.NewSet([]string{})

	s.Add("Hello").Add("World") // Add() method can be chained
	s.Add("Hello")              // duplicate adding values will not effect

	fmt.Println(s)
	// Output:
	// &oop.Set[Hello World]
}

func ExampleSet_Has() {
	s := oop.NewSet([]string{})
	s.Add("Hello")

	fmt.Println(s.Has("Hello"))
	fmt.Println(s.Has("World"))
	// Output:
	// true
	// false
}

func ExampleSet_Delete() {
	s := oop.NewSet([]string{})
	s.Add("Hello").Add("World")

	ok1 := s.Delete("Hello") // succeed
	ok2 := s.Delete("Hello") // failed

	fmt.Println(s)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// &oop.Set[World]
	// true
	// false
}

func ExampleSet_Clear() {
	s := oop.NewSet([]string{})
	s.Add("Hello").Add("World")

	s.Clear()

	fmt.Println(s)
	// Output:
	// &oop.Set[]
}

func ExampleSet_Values() {
	s := oop.NewSet([]string{})
	s.Add("Hello").Add("World")

	fmt.Println(s.Values())
	// Output:
	// [Hello World]
}

func ExampleSet_ForEach() {
	s := oop.NewSet([]string{})
	s.Add("Hello").Add("World")

	s.ForEach(func(item string) {
		fmt.Println(item)
	})
	// Output:
	// Hello
	// World
}

func ExampleSet_Size() {
	s := oop.NewSet([]string{})
	s.Add("Hello").Add("World")

	fmt.Println(s.Size())
	// Output:
	// 2
}
