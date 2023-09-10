package slicex_test

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ayonli/goext/slicex"
)

func ExampleAt_int() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	v1, ok1 := slicex.At(list, 3)
	v2, ok2 := slicex.At(list, -3)
	v3, ok3 := slicex.At(list, 10)

	fmt.Println(v1, ok1)
	fmt.Println(v2, ok2)
	fmt.Println(v3, ok3)
	// Output:
	// 3 true
	// 7 true
	// 0 false
}

func ExampleAt_string() {
	list := []string{"Hello", "World", "Hi", "A-yon"}
	v1, ok1 := slicex.At(list, 1)
	v2, ok2 := slicex.At(list, -2)
	v3, ok3 := slicex.At(list, 4)

	fmt.Println(v1, ok1)
	fmt.Println(v2, ok2)
	fmt.Printf("%#v %v\n", v3, ok3)
	// Output:
	// World true
	// Hi true
	// "" false
}

func ExampleLastIndex_int() {
	list := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
	idx1 := slicex.LastIndex(list, 3)
	idx2 := slicex.LastIndex(list, 10)

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 7
	// -1
}

func ExampleLastIndex_string() {
	list := []string{"Hello", "World", "Hello", "A-yon"}
	idx1 := slicex.LastIndex(list, "Hello")
	idx2 := slicex.LastIndex(list, "Hi")

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 2
	// -1
}

func ExampleCount_int() {
	list := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
	count1 := slicex.Count(list, 3)
	count2 := slicex.Count(list, 10)

	fmt.Println(count1)
	fmt.Println(count2)
	// Output:
	// 2
	// 0
}

func ExampleCount_string() {
	list := []string{"Hello", "World", "Hello", "A-yon"}
	count1 := slicex.Count(list, "Hello")
	count2 := slicex.Count(list, "Hi")

	fmt.Println(count1)
	fmt.Println(count2)
	// Output:
	// 2
	// 0
}

func ExampleConcat_int() {
	list1 := []int{0, 1, 2}
	list2 := []int{3, 4, 5}
	list3 := []int{6, 7, 8}
	list4 := []int{9}
	list5 := slicex.Concat(list1, list2, list3, list4)

	fmt.Println(list5)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleConcat_string() {
	list1 := []string{"Hello", "World"}
	list2 := []string{"Hi", "A-yon"}
	list3 := slicex.Concat(list1, list2)

	fmt.Println(list3)
	// Output:
	// [Hello World Hi A-yon]
}

func ExampleUniq_int() {
	list1 := []int{1, 2, 3, 3, 4, 3, 2, 5, 5, 1}
	list2 := slicex.Uniq(list1)

	fmt.Println(list2)
	// Output:
	// [1 2 3 4 5]
}

func ExampleUniq_string() {
	list1 := []string{"Hello", "World", "Hello", "A-yon"}
	list2 := slicex.Uniq(list1)

	fmt.Println(list2)
	// Output:
	// [Hello World A-yon]
}

func ExampleUniqBy() {
	list1 := []map[string]string{
		{
			"id":  "world",
			"tag": "A",
		},
		{
			"id":  "ayon",
			"tag": "B",
		},
		{
			"id":  "world",
			"tag": "C",
		},
		{
			"name": "foo", // this item will be removed from the result since we order the slice
			"tag":  "D",   // by `id` which is missing here.
		},
		nil, // nil will be ignored and removed
	}
	list2 := slicex.UniqBy(list1, "id")

	fmt.Println(list2)
	// Output:
	// [map[id:world tag:A] map[id:ayon tag:B]]
}

func ExampleFlat_int() {
	list1 := [][]int{{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}}
	list2 := slicex.Flat(list1)

	fmt.Println(list2)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleFlat_string() {
	list1 := [][]string{{"Hello", "World"}, {"Hi", "A-yon"}}
	list2 := slicex.Flat(list1)

	fmt.Println(list2)
	// Output:
	// [Hello World Hi A-yon]
}

func ExampleSlice_int() {
	list1 := []int{0, 1, 2}
	list2 := list1[0:2] // list2 shares the same underlying array with list1
	list3 := slicex.Slice(list1, 0, 2)
	list4 := slicex.Slice(list1, 0, -1)
	list5 := slicex.Slice(list1, -2, 4)
	list6 := slicex.Slice(list1, 3, 4)
	list7 := slicex.Slice(list1, 3, 2)

	list2[0] = 10  // modifying list2 will affect list1, and vice versa
	list3[0] = 100 // modifying list3 doesn't have side effect, and vice versa

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(list4)
	fmt.Println(list5)
	fmt.Println(list6)
	fmt.Println(list7)
	// Output:
	// [10 1 2]
	// [10 1]
	// [100 1]
	// [0 1]
	// [1 2]
	// []
	// []
}

func ExampleSlice_string() {
	list1 := []string{"Hello", "World", "Hi", "A-yon"}
	list2 := list1[0:2] // list2 shares the same underlying array with list1
	list3 := slicex.Slice(list1, 0, 2)
	list4 := slicex.Slice(list1, 0, -2)

	list2[0] = "Hi"   // modifying list2 will affect list1, and vice versa
	list3[0] = "Hola" // modifying list3 doesn't have side effect, and vice versa

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// [Hi World Hi A-yon]
	// [Hi World]
	// [Hola World]
	// [Hello World]
}

func ExampleSplit_int() {
	list1 := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1}
	list2 := slicex.Split(list1, 2)
	list3 := slicex.Split(list1, 5)
	list4 := slicex.Split(list1, 1)

	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// [[0 1] [3 4 5 4 3] [1]]
	// [[0 1 2 3 4] [4 3 2 1]]
	// [[0] [2 3 4 5 4 3 2] []]
}

func ExampleSplit_string() {
	list1 := []string{"foo", "bar", "foo", "abc", "def", "foo", "bar"}
	list2 := slicex.Split(list1, "foo")
	list3 := slicex.Split(list1, "bar")

	fmt.Println(list2)
	fmt.Println(list3)
	// Output:
	// [[] [bar] [abc def] [bar]]
	// [[foo] [foo abc def foo] []]
}

func ExampleChunk_int() {
	list1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	list2 := slicex.Chunk(list1, 2)
	list3 := slicex.Chunk(list1, 3)

	fmt.Println(list2)
	fmt.Println(list3)
	// Output:
	// [[0 1] [2 3] [4 5] [6 7] [8 9]]
	// [[0 1 2] [3 4 5] [6 7 8] [9]]
}

func ExampleChunk_string() {
	list1 := []string{"Hello", "World", "Hi", "A-yon"}
	list2 := slicex.Chunk(list1, 2)
	list3 := slicex.Chunk(list1, 3)

	fmt.Println(list2)
	fmt.Println(list3)
	// Output:
	// [[Hello World] [Hi A-yon]]
	// [[Hello World Hi] [A-yon]]
}

func ExampleJoin_int() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	str1 := slicex.Join(list, ",")
	str2 := slicex.Join(list, " ")

	fmt.Println(str1)
	fmt.Println(str2)
	// Output:
	// 0,1,2,3,4,5,6,7,8,9
	// 0 1 2 3 4 5 6 7 8 9
}

func ExampleJoin_string() {
	list := []string{"Hello", "World", "Hi", "A-yon"}
	str1 := slicex.Join(list, ", ")
	str2 := slicex.Join(list, " | ")

	fmt.Println(str1)
	fmt.Println(str2)
	// Output:
	// Hello, World, Hi, A-yon
	// Hello | World | Hi | A-yon
}

func ExampleJoin_any() {
	list := []any{"Hello", 1, 2.0, map[string]uint8{"foo": 1}, nil}
	str1 := slicex.Join(list, ", ")
	str2 := slicex.Join(list, " | ")

	fmt.Println(str1)
	fmt.Println(str2)
	// Output:
	// Hello, 1, 2, map[foo:1], <nil>
	// Hello | 1 | 2 | map[foo:1] | <nil>
}

func ExampleEvery_int() {
	list := []int{0, 1, 2}
	ok1 := slicex.Every(list, func(item int, _ int) bool { return item < 3 })
	ok2 := slicex.Every(list, func(item int, _ int) bool { return item < 2 })

	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// true
	// false
}

func ExampleEvery_string() {
	list1 := []string{"Hello", "World"}
	ok1 := slicex.Every(list1, func(item string, _ int) bool { return strings.Contains(item, "o") })
	ok2 := slicex.Every(list1, func(item string, _ int) bool { return strings.Contains(item, "H") })

	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// true
	// false
}

func ExampleSome_int() {
	list := []int{0, 1, 2}
	ok1 := slicex.Some(list, func(item int, _ int) bool { return item < 1 })
	ok2 := slicex.Some(list, func(item int, _ int) bool { return item < 0 })

	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// true
	// false
}

func ExampleSome_string() {
	list1 := []string{"Hello", "World"}
	ok1 := slicex.Some(list1, func(item string, _ int) bool { return strings.Contains(item, "H") })
	ok2 := slicex.Some(list1, func(item string, _ int) bool { return strings.Contains(item, "i") })

	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// true
	// false
}

func ExampleFind_int() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	value1, ok1 := slicex.Find(list, func(item int, _ int) bool { return item >= 5 })
	value2, ok2 := slicex.Find(list, func(item int, _ int) bool { return item >= 10 })

	fmt.Println(value1, ok1)
	fmt.Println(value2, ok2)
	// Output:
	// 5 true
	// 0 false
}

func ExampleFind_string() {
	list := []string{"Hello", "World", "Hi", "A-yon"}
	value1, ok1 := slicex.Find(list, func(item string, idx int) bool { return idx == 2 })
	value2, ok2 := slicex.Find(list, func(item string, idx int) bool { return idx == 5 })

	fmt.Println(value1, ok1)
	fmt.Printf("%#v %v\n", value2, ok2)
	// Output:
	// Hi true
	// "" false
}

func ExampleFindLast_int() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	value1, ok1 := slicex.FindLast(list, func(item int, _ int) bool { return item >= 5 })
	value2, ok2 := slicex.FindLast(list, func(item int, _ int) bool { return item >= 10 })

	fmt.Println(value1, ok1)
	fmt.Println(value2, ok2)
	// Output:
	// 9 true
	// 0 false
}

func ExampleFindLast_string() {
	list := []string{"Hello", "World", "Hi", "A-yon"}
	value1, ok1 := slicex.FindLast(list, func(item string, _ int) bool {
		return strings.HasPrefix(item, "H")
	})
	value2, ok2 := slicex.FindLast(list, func(item string, _ int) bool {
		return strings.HasPrefix(item, "o")
	})

	fmt.Println(value1, ok1)
	fmt.Printf("%#v %v\n", value2, ok2)
	// Output:
	// Hi true
	// "" false
}

func ExampleFindIndex_int() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	idx1 := slicex.FindIndex(list, func(item int, _ int) bool { return item >= 5 })
	idx2 := slicex.FindIndex(list, func(item int, _ int) bool { return item >= 10 })

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 5
	// -1
}

func ExampleFindIndex_string() {
	list := []string{"Hello", "World", "Hi", "A-yon"}
	idx1 := slicex.FindLastIndex(list, func(item string, idx int) bool { return item == "Hi" })
	idx2 := slicex.FindLastIndex(list, func(item string, idx int) bool { return item == "Haha" })

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 2
	// -1
}

func ExampleFindLastIndex_int() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	idx1 := slicex.FindLastIndex(list, func(item int, _ int) bool { return item >= 5 })
	idx2 := slicex.FindLastIndex(list, func(item int, _ int) bool { return item >= 10 })

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 9
	// -1
}

func ExampleFindLastIndex_string() {
	list := []string{"Hello", "World", "Hello", "A-yon"}
	idx1 := slicex.FindLastIndex(list, func(item string, idx int) bool { return item == "Hello" })
	idx2 := slicex.FindLastIndex(list, func(item string, idx int) bool { return item == "Haha" })

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 2
	// -1
}

func ExampleFilter_int() {
	list1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	list2 := slicex.Filter(list1, func(item int, _ int) bool { return item%2 == 0 })

	fmt.Println(list2)
	// Output:
	// [0 2 4 6 8]
}

func ExampleFilter_string() {
	list1 := []string{"Hello", "World", "Hi", "A-yon"}
	list2 := slicex.Filter(list1, func(item string, _ int) bool { return strings.HasPrefix(item, "H") })

	fmt.Println(list2)
	// Output:
	// [Hello Hi]
}

func ExampleForEach_int() {
	list1 := []int{0, 1, 2}

	slicex.ForEach(list1, func(item int, _ int) {
		fmt.Println(item)
	})
	// Output:
	// 0
	// 1
	// 2
}

func ExampleForEach_string() {
	list1 := []string{"Hello", "World"}

	slicex.ForEach(list1, func(item string, _ int) {
		fmt.Println(item)
	})
	// Output:
	// Hello
	// World
}

func ExampleMap_int() {
	list1 := []int{0, 1, 2}
	list2 := slicex.Map(list1, func(item int, _ int) int { return item * 2 })

	fmt.Println(list2)
	// Output:
	// [0 2 4]
}

func ExampleMap_string() {
	list1 := []string{"Hello", "World"}
	list2 := slicex.Map(list1, func(item string, _ int) string { return strings.ToUpper(item) })

	fmt.Println(list2)
	// Output:
	// [HELLO WORLD]
}

func ExampleReduce_int() {
	list := []int{0, 1, 2}
	sum := slicex.Reduce(list, func(sum int, item int, _ int) int { return sum + item }, 0)

	fmt.Println(sum)
	// Output:
	// 3
}

func ExampleReduce_map() {
	list := []map[string]string{
		{
			"id":  "foo",
			"tag": "Hello",
		},
		{
			"id":  "bar",
			"tag": "World",
		},
	}
	record := slicex.Reduce(list, func(
		record map[string]map[string]string,
		item map[string]string,
		_ int,
	) map[string]map[string]string {
		record[item["id"]] = item
		return record
	}, map[string]map[string]string{})

	fmt.Println(record)
	// Output:
	// map[bar:map[id:bar tag:World] foo:map[id:foo tag:Hello]]
}

func ExamplePop_int() {
	list := []int{0, 1, 2}
	item := slicex.Pop(&list)

	fmt.Println(list)
	fmt.Println(item)
	// Output:
	// [0 1]
	// 2
}

func ExamplePop_string() {
	list := []string{"Hello", "World", "Hi", "A-yon"}
	item := slicex.Pop(&list)

	fmt.Println(list)
	fmt.Println(item)
	// Output:
	// [Hello World Hi]
	// A-yon
}

func ExamplePush_int() {
	list := []int{0, 1, 2}
	length := slicex.Push(&list, 3, 4)

	fmt.Println(list)
	fmt.Println(length)
	// Output:
	// [0 1 2 3 4]
	// 5
}

func ExamplePush_string() {
	list := []string{"Hello", "World"}
	length := slicex.Push(&list, "Hi", "A-yon")

	fmt.Println(list)
	fmt.Println(length)
	// Output:
	// [Hello World Hi A-yon]
	// 4
}

func ExampleShift_int() {
	list := []int{0, 1, 2}
	item := slicex.Shift(&list)

	fmt.Println(list)
	fmt.Println(item)
	// Output:
	// [1 2]
	// 0
}

func ExampleShift_string() {
	list := []string{"Hello", "World", "Hi", "A-yon"}
	item := slicex.Shift(&list)

	fmt.Println(list)
	fmt.Println(item)
	// Output:
	// [World Hi A-yon]
	// Hello
}

func ExampleUnshift_int() {
	list := []int{0, 1, 2}
	length := slicex.Unshift(&list, -2, -1)

	fmt.Println(list)
	fmt.Println(length)
	// Output:
	// [-2 -1 0 1 2]
	// 5
}

func ExampleUnshift_string() {
	list := []string{"Hello", "World"}
	length := slicex.Unshift(&list, "Hi", "A-yon")

	fmt.Println(list)
	fmt.Println(length)
	// Output:
	// [Hi A-yon Hello World]
	// 4
}

func ExampleShuffle_int() {
	list1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	list2 := slices.Clone(list1)
	list3 := slices.Clone(list1)

	slicex.Shuffle(list2) // now list2 is in any order
	slicex.Shuffle(list3) // now list3 is in any order

	fmt.Printf("list2(len: %d) != list1(len: %d): %v\n", len(list2), len(list1), !slices.Equal(list2, list1))
	fmt.Printf("list3(len: %d) != list1(len: %d): %v\n", len(list3), len(list1), !slices.Equal(list3, list1))
	fmt.Printf("list3(len: %d) != list2(len: %d): %v\n", len(list3), len(list2), !slices.Equal(list3, list2))
	// Output:
	// list2(len: 10) != list1(len: 10): true
	// list3(len: 10) != list1(len: 10): true
	// list3(len: 10) != list2(len: 10): true
}

func ExampleShuffle_sting() {
	list1 := []string{"Hello", "World", "Hi", "A-yon"}
	list2 := slices.Clone(list1)
	list3 := slices.Clone(list1)

	slicex.Shuffle(list2) // now list2 is in any order
	slicex.Shuffle(list3) // now list3 is in any order

	fmt.Printf("list2(len: %d) != list1(len: %d): %v\n", len(list2), len(list1), !slices.Equal(list2, list1))
	fmt.Printf("list3(len: %d) != list1(len: %d): %v\n", len(list3), len(list1), !slices.Equal(list3, list1))
	fmt.Printf("list3(len: %d) != list2(len: %d): %v\n", len(list3), len(list2), !slices.Equal(list3, list2))
	// Output:
	// list2(len: 4) != list1(len: 4): true
	// list3(len: 4) != list1(len: 4): true
	// list3(len: 4) != list2(len: 4): true
}

func ExampleOrderBy() {
	list1 := []map[string]any{
		{
			"id":  "world",
			"age": 53,
			"tag": "A",
		},
		{
			"id":  "ayon",
			"age": 28,
			"tag": "B",
		},
		{
			"id":  "claire",
			"age": 25,
			"tag": "B",
		},
	}
	list2 := slicex.OrderBy(list1, "age", "asc")
	list3 := slicex.OrderBy(list2, "age", "desc")
	list4 := slicex.OrderBy(list1, "id", "asc")
	list5 := slicex.OrderBy(list2, "tag", "desc")
	list6 := slicex.OrderBy(list1, "name", "desc") // this has no effect since key 'name' doesn't exist
	list7 := slicex.OrderBy(list1, "age", "foo")   // this has no effect since the `order` can only be either 'asc' or 'desc'

	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(list4)
	fmt.Println(list5)
	fmt.Println(list6)
	fmt.Println(list7)
	// Output:
	// [map[age:25 id:claire tag:B] map[age:28 id:ayon tag:B] map[age:53 id:world tag:A]]
	// [map[age:53 id:world tag:A] map[age:28 id:ayon tag:B] map[age:25 id:claire tag:B]]
	// [map[age:28 id:ayon tag:B] map[age:25 id:claire tag:B] map[age:53 id:world tag:A]]
	// [map[age:28 id:ayon tag:B] map[age:25 id:claire tag:B] map[age:53 id:world tag:A]]
	// [map[age:53 id:world tag:A] map[age:28 id:ayon tag:B] map[age:25 id:claire tag:B]]
	// [map[age:53 id:world tag:A] map[age:28 id:ayon tag:B] map[age:25 id:claire tag:B]]
}

func ExampleGroupBy() {
	list := []map[string]string{
		{
			"group": "world",
			"tag":   "A",
		},
		{
			"group": "room",
			"tag":   "B",
		},
		{
			"group": "room",
			"tag":   "C",
		},
	}
	record := slicex.GroupBy(list, func(item map[string]string, _ int) string {
		return item["group"]
	})

	fmt.Println(record)
	// Output:
	// map[room:[map[group:room tag:B] map[group:room tag:C]] world:[map[group:world tag:A]]]
}

func ExampleDiff_int() {
	list1 := []int{0, 1, 2, 3, 4, 5}
	list2 := []int{2, 3, 4, 5, 6, 7}
	list3 := slicex.Diff(list1, list2)

	fmt.Println(list3)
	// Output:
	// [0 1]
}

func ExampleDiff_string() {
	list1 := []string{"Hello", "World"}
	list2 := []string{"Hi", "World"}
	list3 := slicex.Diff(list1, list2)

	fmt.Println(list3)
	// Output:
	// [Hello]
}

func ExampleXor_int() {
	list1 := []int{0, 1, 2, 3, 4, 5}
	list2 := []int{2, 3, 4, 5, 6, 7}
	list3 := slicex.Xor(list1, list2)
	list4 := slicex.Xor(list1, list2, nil) // nil will be ignored

	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// [0 1 6 7]
	// [0 1 6 7]
}

func ExampleXor_string() {
	list1 := []string{"Hello", "World"}
	list2 := []string{"Hi", "World"}
	list3 := slicex.Xor(list1, list2)
	list4 := slicex.Xor(list1, list2, nil) // nil will be ignored

	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// [Hello Hi]
	// [Hello Hi]
}

func ExampleUnion_int() {
	list1 := []int{0, 1, 2, 3, 4, 5}
	list2 := []int{2, 3, 4, 5, 6, 7}
	list3 := slicex.Union(list1, list2)
	list4 := slicex.Union(list1, list2, nil) // nil will be ignored

	fmt.Println(list3)
	fmt.Println(list4)
	// Output:
	// [0 1 2 3 4 5 6 7]
	// [0 1 2 3 4 5 6 7]
}

func ExampleUnion_string() {
	list1 := []string{"Hello", "World"}
	list2 := []string{"Hi", "World"}
	list3 := slicex.Union(list1, list2)
	list4 := slicex.Union(list1, list2, nil) // nil will be ignored
	fmt.Println(list4)

	fmt.Println(list3)
	// Output:
	// [Hello World Hi]
	// [Hello World Hi]
}

func ExampleIntersect_int() {
	list1 := []int{0, 1, 2, 3, 4, 5}
	list2 := []int{2, 3, 4, 5, 6, 7}
	list3 := slicex.Intersect(list1, list2)

	fmt.Println(list3)
	// Output:
	// [2 3 4 5]
}

func ExampleIntersect_string() {
	list1 := []string{"Hello", "World"}
	list2 := []string{"Hi", "World"}
	list3 := slicex.Intersect(list1, list2)

	fmt.Println(list3)
	// Output:
	// [World]
}
