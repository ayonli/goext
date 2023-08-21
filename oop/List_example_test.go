package oop_test

import (
	"fmt"

	"github.com/ayonli/goext/oop"
)

func ExampleNewList() {
	list := oop.NewList([]string{"foo", "bar"})

	fmt.Println(list)
	fmt.Printf("%#v\n", list)
	// Output:
	// &[foo bar]
	// &oop.List[string]{"foo", "bar"}
}

func ExampleList_At() {
	list := oop.NewList([]string{"foo", "bar"})

	fmt.Println(list.At(0))
	fmt.Println(list.At(-1)) // negative indexing is supported
	fmt.Println(list.At(2))  // exceeding boundary returns zero-value and `false`
	// Output:
	// foo true
	// bar true
	//  false
}

func ExampleList_IndexOf() {
	list := oop.NewList([]string{"foo", "bar"})

	fmt.Println(list.IndexOf("foo"))
	fmt.Println(list.IndexOf("Hello"))
	// Output:
	// 0
	// -1
}

func ExampleList_LastIndexOf() {
	list := oop.NewList([]string{"foo", "bar", "foo", "bar"})

	fmt.Println(list.LastIndexOf("foo"))
	fmt.Println(list.LastIndexOf("Hello"))
	// Output:
	// 2
	// -1
}

func ExampleList_Length() {
	list := oop.NewList([]string{"foo", "bar"})

	fmt.Println(list.Length())
	// Output:
	// 2
}

func ExampleList_Values() {
	list := oop.NewList([]string{"foo", "bar"})

	fmt.Println(list.Values())
	// Output:
	// [foo bar]
}

func ExampleList_Clone() {
	list1 := oop.NewList([]string{"foo", "bar"})
	list2 := list1.Clone()

	fmt.Println(list2)
	// Output:
	// &[foo bar]
}

func ExampleList_Equal() {
	list1 := oop.NewList([]string{"foo", "bar"})
	list2 := oop.NewList([]string{"foo", "bar"})
	list3 := oop.NewList([]string{"Hello", "World"})

	fmt.Println(list1.Equal(list2))
	fmt.Println(list1.Equal(list3))
	// Output:
	// true
	// false
}

func ExampleList_Contains() {
	list := oop.NewList([]string{"foo", "bar"})

	fmt.Println(list.Contains("foo"))
	fmt.Println(list.Contains("Hello"))
	// Output:
	// true
	// false
}

func ExampleList_Count() {
	list := oop.NewList([]string{"foo", "bar", "foo"})

	fmt.Println(list.Count("foo"))
	fmt.Println(list.Count("bar"))
	fmt.Println(list.Count("Hello"))
	// Output:
	// 2
	// 1
	// 0
}

func ExampleList_Concat() {
	list1 := oop.NewList([]string{"foo", "bar"})
	list2 := list1.Concat(oop.NewList([]string{"Hello", "World"}))

	fmt.Println(list2)
	// Output:
	// &[foo bar Hello World]
}

func ExampleList_Uniq() {
	list := oop.NewList([]string{"hello", "world", "hi", "world"})

	fmt.Println(list.Uniq())
	// Output:
	// &[hello world hi]
}

func ExampleList_Slice() {
	list := oop.NewList([]string{"foo", "bar", "hello", "world"})

	fmt.Println(list.Slice(0, 2))
	fmt.Println(list.Slice(-2, list.Length())) // negative indexing is supported
	// Output:
	// &[foo bar]
	// &[hello world]
}

func ExampleList_Chunk() {
	list := oop.NewList([]string{"foo", "bar", "hello", "world"})
	chunks := list.Chunk(2)

	fmt.Println(chunks)
	fmt.Printf("%#v\n", chunks)
	// Output:
	// [&[foo bar] &[hello world]]
	// []*oop.List[string]{&oop.List[string]{"foo", "bar"}, &oop.List[string]{"hello", "world"}}
}

func ExampleList_Join() {
	list := oop.NewList([]string{"foo", "bar"})

	fmt.Println(list.Join(","))
	// Output:
	// foo,bar
}

func ExampleList_Replace() {
	list := oop.NewList([]string{"foo", "bar", "hello", "world"})

	list.Replace(2, 4, "hi", "ayon") // Replace() mutates the list and returns itself

	fmt.Println(list)
	// Output:
	// &[foo bar hi ayon]
}

func ExampleList_Reverse() {
	list := oop.NewList([]string{"foo", "bar"})

	list.Reverse() // Reverse() mutates the list and returns itself

	fmt.Println(list)
	// Output:
	// &[bar foo]
}

func ExampleList_ToReversed() {
	list1 := oop.NewList([]string{"foo", "bar"})
	list2 := list1.ToReversed() // ToReversed() returns a copy of the list with all items reversed

	fmt.Println(list1)
	fmt.Println(list2)
	// Output:
	// &[foo bar]
	// &[bar foo]
}

func ExampleList_Sort() {
	list := oop.NewList([]string{"foo", "bar"})

	list.Sort() // Sort() mutates the list and returns itself

	fmt.Println(list)
	// Output:
	// &[bar foo]
}

func ExampleList_ToSorted() {
	list1 := oop.NewList([]string{"foo", "bar"})
	list2 := list1.ToSorted() // ToSorted() returns a copy of the list with all items reversed

	fmt.Println(list1)
	fmt.Println(list2)
	// Output:
	// &[foo bar]
	// &[bar foo]
}

func ExampleList_Every() {
	list1 := oop.NewList([]string{"foo", "bar"})
	ok1 := list1.Every(func(item string, idx int) bool { return len(item) > 0 })
	ok2 := list1.Every(func(item string, idx int) bool { return item[0:1] == "f" })

	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// true
	// false
}

func ExampleList_Some() {
	list1 := oop.NewList([]string{"foo", "bar"})
	ok1 := list1.Some(func(item string, idx int) bool { return item[0:1] == "f" })
	ok2 := list1.Some(func(item string, idx int) bool { return len(item) > 3 })

	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// true
	// false
}

func ExampleList_Find() {
	list := oop.NewList([]string{"foo", "bar"})
	item1, ok1 := list.Find(func(item string, idx int) bool { return item[0:1] == "f" })
	item2, ok2 := list.Find(func(item string, idx int) bool { return item[0:1] == "a" })

	fmt.Println(item1, ok1)
	fmt.Println(item2, ok2)
	// Output:
	// foo true
	//  false
}

func ExampleList_FindLast() {
	list := oop.NewList([]string{"foo", "bar"})
	item1, ok1 := list.FindLast(func(item string, idx int) bool { return len(item) > 0 })
	item2, ok2 := list.FindLast(func(item string, idx int) bool { return len(item) > 3 })

	fmt.Println(item1, ok1)
	fmt.Println(item2, ok2)
	// Output:
	// bar true
	//  false
}

func ExampleList_FindIndex() {
	list := oop.NewList([]string{"foo", "bar"})
	idx1 := list.FindIndex(func(item string, idx int) bool { return item[0:1] == "f" })
	idx2 := list.FindIndex(func(item string, idx int) bool { return item[0:1] == "a" })

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 0
	// -1
}

func ExampleList_FindLastIndex() {
	list := oop.NewList([]string{"foo", "bar"})
	idx1 := list.FindLastIndex(func(item string, idx int) bool { return item[0:1] == "b" })
	idx2 := list.FindLastIndex(func(item string, idx int) bool { return item[0:1] == "a" })

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 1
	// -1
}

func ExampleList_Filter() {
	list1 := oop.NewList([]string{"foo", "bar", "hello", "world"})
	list2 := list1.Filter(func(item string, idx int) bool { return len(item) > 3 })

	fmt.Println(list2)
	// Output:
	// &[hello world]
}

func ExampleList_ForEach() {
	list := oop.NewList([]string{"foo", "bar"})

	list.ForEach(func(item string, idx int) {
		fmt.Println(idx, "=>", item)
	})
	// Output:
	// 0 => foo
	// 1 => bar
}

func ExampleList_Pop() {
	list := oop.NewList([]string{"foo", "bar"})
	item := list.Pop()

	fmt.Println(item)
	fmt.Println(list)
	// Output:
	// bar
	// &[foo]
}

func ExampleList_Push() {
	list := oop.NewList([]string{"foo", "bar"})
	length := list.Push("hello", "world")

	fmt.Println(length)
	fmt.Println(list)
	// Output:
	// 4
	// &[foo bar hello world]
}

func ExampleList_Shift() {
	list := oop.NewList([]string{"foo", "bar"})
	item := list.Shift()

	fmt.Println(item)
	fmt.Println(list)
	// Output:
	// foo
	// &[bar]
}

func ExampleList_Unshift() {
	list := oop.NewList([]string{"foo", "bar"})
	length := list.Unshift("hello", "world")

	fmt.Println(length)
	fmt.Println(list)
	// Output:
	// 4
	// &[hello world foo bar]
}

func ExampleList_Shuffle() {
	list1 := oop.NewList([]int{0, 12, 3, 4, 5, 6, 7, 8, 9})
	list2 := list1.Clone()
	list3 := list1.Clone()

	list2.Shuffle() // Shuffle() mutates the list and returns itself
	list3.Shuffle()

	fmt.Printf("list2(len: %d) != list1(len: %d): %v\n", list2.Length(), list1.Length(), !list2.Equal(list1))
	fmt.Printf("list3(len: %d) != list1(len: %d): %v\n", list3.Length(), list1.Length(), !list3.Equal(list1))
	fmt.Printf("list3(len: %d) != list2(len: %d): %v\n", list3.Length(), list2.Length(), !list3.Equal(list2))
	// Output:
	// list2(len: 9) != list1(len: 9): true
	// list3(len: 9) != list1(len: 9): true
	// list3(len: 9) != list2(len: 9): true
}

func ExampleList_Diff() {
	list1 := oop.NewList([]int{0, 1, 2, 3, 4, 5})
	list2 := oop.NewList([]int{2, 3, 4, 5, 6, 7})
	list3 := list1.Diff(list2)
	list4 := list1.Diff(list2, nil) // nil will be ignored

	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// &[0 1]
	// &[0 1]
}

func ExampleList_Xor() {
	list1 := oop.NewList([]int{0, 1, 2, 3, 4, 5})
	list2 := oop.NewList([]int{2, 3, 4, 5, 6, 7})
	list3 := list1.Xor(list2)
	list4 := list1.Xor(list2, nil) // nil will be ignored

	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// &[0 1 6 7]
	// &[0 1 6 7]
}

func ExampleList_Union() {
	list1 := oop.NewList([]int{0, 1, 2, 3, 4, 5})
	list2 := oop.NewList([]int{2, 3, 4, 5, 6, 7})
	list3 := list1.Union(list2)
	list4 := list1.Union(list2, nil) // nil will be ignored

	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// &[0 1 2 3 4 5 6 7]
	// &[0 1 2 3 4 5 6 7]
}

func ExampleList_Intersect() {
	list1 := oop.NewList([]int{0, 1, 2, 3, 4, 5})
	list2 := oop.NewList([]int{2, 3, 4, 5, 6, 7})
	list3 := list1.Intersect(list2)

	fmt.Println(list3)
	// Output:
	// &[2 3 4 5]
}
