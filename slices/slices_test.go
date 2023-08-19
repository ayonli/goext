package slices

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	stringExt "github.com/ayonli/goext/strings"
	"github.com/stretchr/testify/assert"
)

func TestAt(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		v1, ok1 := At(list, 3)
		v2, ok2 := At(list, -3)
		v3, ok3 := At(list, 10)

		assert.Equal(t, 3, v1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, 7, v2)
		assert.Equal(t, true, ok2)
		assert.Equal(t, 0, v3)
		assert.Equal(t, false, ok3)
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hi", "A-yon"}
		v1, ok1 := At(list, 1)
		v2, ok2 := At(list, -2)
		v3, ok3 := At(list, 4)

		assert.Equal(t, "World", v1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, "Hi", v2)
		assert.Equal(t, true, ok2)
		assert.Equal(t, "", v3)
		assert.Equal(t, false, ok3)
	})
}

func TestLastIndex(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
		idx1 := LastIndex(list, 3)
		idx2 := LastIndex(list, 10)

		assert.Equal(t, 7, idx1)
		assert.Equal(t, -1, idx2)
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hello", "A-yon"}
		idx1 := LastIndex(list, "Hello")
		idx2 := LastIndex(list, "Hi")

		assert.Equal(t, 2, idx1)
		assert.Equal(t, -1, idx2)
	})
}

func TestCount(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
		count1 := Count(list, 3)
		count2 := Count(list, 10)

		assert.Equal(t, 2, count1)
		assert.Equal(t, 0, count2)
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hello", "A-yon"}
		count1 := Count(list, "Hello")
		count2 := Count(list, "Hi")

		assert.Equal(t, 2, count1)
		assert.Equal(t, 0, count2)
	})
}

func TestConcat(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2}
		list2 := []int{3, 4, 5}
		list3 := []int{6, 7, 8}
		list4 := []int{9}
		list5 := Concat(list1, list2, list3, list4)

		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, list5)
		assert.Equal(t, 10, len(list5))
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World"}
		list2 := []string{"Hi", "A-yon"}
		list3 := Concat(list1, list2)

		assert.Equal(t, []string{"Hello", "World", "Hi", "A-yon"}, list3)
	})
}

func TestUniq(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list1 := []int{1, 2, 3, 3, 4, 3, 2, 5, 5, 1}
		list2 := Uniq(list1)

		assert.Equal(t, []int{1, 2, 3, 4, 5}, list2)
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World", "Hello", "A-yon"}
		list2 := Uniq(list1)

		assert.Equal(t, []string{"Hello", "World", "A-yon"}, list2)
	})
}

func TestUniqBy(t *testing.T) {
	t.Run("[]map[string]string", func(t *testing.T) {
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
		}
		list2 := UniqBy(list1, "id")

		assert.Equal(t, []map[string]string{
			{
				"id":  "world",
				"tag": "A",
			},
			{
				"id":  "ayon",
				"tag": "B",
			},
		}, list2)
	})
}

func TestFlat(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list1 := [][]int{{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}}
		list2 := Flat(list1)

		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, list2)
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := [][]string{{"Hello", "World"}, {"Hi", "A-yon"}}
		list2 := Flat(list1)

		assert.Equal(t, []string{"Hello", "World", "Hi", "A-yon"}, list2)
	})
}

func TestSlice(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2}
		list2 := list1[0:2] // list2 shares the same underlying array with list1
		list3 := Slice(list1, 0, 2)
		list4 := Slice(list1, 0, -1)
		list5 := Slice(list1, 2, 4)
		list6 := Slice(list1, 3, 4)
		list7 := Slice(list1, 3, 2)

		list2[0] = 10  // modifying list2 will affect list1, and vice versa
		list3[0] = 100 // modifying list3 doesn't have side effect, and vice versa

		assert.Equal(t, []int{10, 1, 2}, list1)
		assert.Equal(t, []int{10, 1}, list2)
		assert.Equal(t, []int{100, 1}, list3)
		assert.Equal(t, []int{0, 1}, list4)
		assert.Equal(t, []int{2}, list5)
		assert.Equal(t, []int{}, list6)
		assert.Equal(t, []int{}, list7)
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World", "Hi", "A-yon"}
		list2 := list1[0:2] // list2 shares the same underlying array with list1
		list3 := Slice(list1, 0, 2)
		list4 := Slice(list1, 0, -2)

		list2[0] = "Hi"   // modifying list2 will affect list1, and vice versa
		list3[0] = "Hola" // modifying list3 doesn't have side effect, and vice versa

		assert.Equal(t, []string{"Hi", "World", "Hi", "A-yon"}, list1)
		assert.Equal(t, []string{"Hi", "World"}, list2)
		assert.Equal(t, []string{"Hola", "World"}, list3)
		assert.Equal(t, []string{"Hello", "World"}, list4)
	})
}

func TestChunk(t *testing.T) {
	t.Run("[]int->[][]int", func(t *testing.T) {
		list1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		list2 := Chunk(list1, 2)
		list3 := Chunk(list1, 3)

		assert.Equal(t, [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}, list2)
		assert.Equal(t, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}}, list3)
	})

	t.Run("[]int->[][]string", func(t *testing.T) {
		list1 := []string{"Hello", "World", "Hi", "A-yon"}
		list2 := Chunk(list1, 2)
		list3 := Chunk(list1, 3)

		assert.Equal(t, [][]string{{"Hello", "World"}, {"Hi", "A-yon"}}, list2)
		assert.Equal(t, [][]string{{"Hello", "World", "Hi"}, {"A-yon"}}, list3)
	})
}

func TestJoin(t *testing.T) {
	t.Run("[]int->string", func(t *testing.T) {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		str1 := Join(list, ",")
		str2 := Join(list, " ")

		assert.Equal(t, "0,1,2,3,4,5,6,7,8,9", str1)
		assert.Equal(t, "0 1 2 3 4 5 6 7 8 9", str2)
	})

	t.Run("[]string->string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hi", "A-yon"}
		str1 := Join(list, ", ")
		str2 := Join(list, " | ")

		assert.Equal(t, "Hello, World, Hi, A-yon", str1)
		assert.Equal(t, "Hello | World | Hi | A-yon", str2)
	})
}

func TestEvery(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2}
		ok1 := Every(list, func(item int, _ int) bool { return item < 3 })
		ok2 := Every(list, func(item int, _ int) bool { return item < 2 })

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World"}
		ok1 := Every(list1, func(item string, _ int) bool { return strings.Contains(item, "o") })
		ok2 := Every(list1, func(item string, _ int) bool { return strings.Contains(item, "H") })

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
	})
}

func TestSome(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2}
		ok1 := Some(list, func(item int, _ int) bool { return item < 1 })
		ok2 := Some(list, func(item int, _ int) bool { return item < 0 })

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World"}
		ok1 := Some(list1, func(item string, _ int) bool { return strings.Contains(item, "H") })
		ok2 := Some(list1, func(item string, _ int) bool { return strings.Contains(item, "i") })

		assert.Equal(t, true, ok1)
		assert.Equal(t, false, ok2)
	})
}

func TestFind(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		value1, ok1 := Find(list, func(item int, _ int) bool { return item >= 5 })
		value2, ok2 := Find(list, func(item int, _ int) bool { return item >= 10 })

		assert.Equal(t, 5, value1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, 0, value2)
		assert.Equal(t, false, ok2)
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hi", "A-yon"}
		value1, ok1 := Find(list, func(item string, idx int) bool { return idx == 2 })
		value2, ok2 := Find(list, func(item string, idx int) bool { return idx == 5 })

		assert.Equal(t, "Hi", value1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, "", value2)
		assert.Equal(t, false, ok2)
	})
}

func TestFindLast(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		value1, ok1 := FindLast(list, func(item int, _ int) bool { return item >= 5 })
		value2, ok2 := FindLast(list, func(item int, _ int) bool { return item >= 10 })

		assert.Equal(t, 9, value1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, 0, value2)
		assert.Equal(t, false, ok2)
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hi", "A-yon"}
		value1, ok1 := FindLast(list, func(item string, _ int) bool {
			return stringExt.StartsWith(item, "H")
		})
		value2, ok2 := FindLast(list, func(item string, _ int) bool {
			return stringExt.StartsWith(item, "o")
		})

		assert.Equal(t, "Hi", value1)
		assert.Equal(t, true, ok1)
		assert.Equal(t, "", value2)
		assert.Equal(t, false, ok2)
	})
}

func TestFilter(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		list2 := Filter(list1, func(item int, _ int) bool { return item%2 == 0 })

		assert.Equal(t, []int{0, 2, 4, 6, 8}, list2)
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World", "Hi", "A-yon"}
		list2 := Filter(list1, func(item string, _ int) bool { return stringExt.StartsWith(item, "H") })

		assert.Equal(t, []string{"Hello", "Hi"}, list2)
	})
}

func TestMap(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2}
		list2 := Map(list1, func(item int, _ int) int { return item * 2 })

		assert.Equal(t, []int{0, 2, 4}, list2)
	})

	t.Run("[]int->[]string", func(t *testing.T) {
		list1 := []int{0, 1, 2}
		list2 := Map(list1, func(item int, _ int) string { return "'" + fmt.Sprint(item) + "'" })

		assert.Equal(t, []string{"'0'", "'1'", "'2'"}, list2)
	})
}

func TestReduce(t *testing.T) {
	t.Run("[]int->int", func(t *testing.T) {
		list := []int{0, 1, 2}
		sum := Reduce(list, func(sum int, item int, _ int) int { return sum + item }, 0)

		assert.Equal(t, 3, sum)
	})

	t.Run("[]map[string]string->map[string]map[string]string", func(t *testing.T) {
		list := []map[string]string{
			{
				"id": "world",
			},
			{
				"id": "ayon",
			},
		}
		record := Reduce(list, func(
			record map[string]map[string]string,
			item map[string]string,
			_ int,
		) map[string]map[string]string {
			record[item["id"]] = item
			return record
		}, map[string]map[string]string{})

		assert.Equal(t, map[string]map[string]string{
			"world": {
				"id": "world",
			},
			"ayon": {
				"id": "ayon",
			},
		}, record)
	})
}

func TestPop(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2}
		item := Pop(&list)

		assert.Equal(t, []int{0, 1}, list)
		assert.Equal(t, 2, item)
		assert.Equal(t, 2, len(list))
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hi", "A-yon"}
		item := Pop(&list)

		assert.Equal(t, []string{"Hello", "World", "Hi"}, list)
		assert.Equal(t, "A-yon", item)
		assert.Equal(t, 3, len(list))
		assert.Equal(t, 3, len(list))
	})
}

func TestPush(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2}
		length := Push(&list, 3, 4)

		assert.Equal(t, []int{0, 1, 2, 3, 4}, list)
		assert.Equal(t, 5, length)
		assert.Equal(t, 5, len(list))
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World"}
		length := Push(&list, "Hi", "A-yon")

		assert.Equal(t, []string{"Hello", "World", "Hi", "A-yon"}, list)
		assert.Equal(t, 4, length)
		assert.Equal(t, 4, len(list))
	})
}

func TestShift(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2}
		item := Shift(&list)

		assert.Equal(t, []int{1, 2}, list)
		assert.Equal(t, 0, item)
		assert.Equal(t, 2, len((list)))
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World", "Hi", "A-yon"}
		item := Shift(&list)

		assert.Equal(t, []string{"World", "Hi", "A-yon"}, list)
		assert.Equal(t, "Hello", item)
		assert.Equal(t, 3, len(list))
	})
}

func TestUnshift(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list := []int{0, 1, 2}
		length := Unshift(&list, -2, -1)

		assert.Equal(t, []int{-2, -1, 0, 1, 2}, list)
		assert.Equal(t, 5, length)
		assert.Equal(t, 5, len(list))
	})

	t.Run("[]string", func(t *testing.T) {
		list := []string{"Hello", "World"}
		length := Unshift(&list, "Hi", "A-yon")

		assert.Equal(t, []string{"Hi", "A-yon", "Hello", "World"}, list)
		assert.Equal(t, 4, length)
		assert.Equal(t, 4, len(list))
	})
}

func TestShuffle(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		list1 := []int{0, 12, 3, 4, 5, 6, 7, 8, 9}
		list2 := slices.Clone(list1)
		list3 := slices.Clone(list1)

		Shuffle(list2)
		Shuffle(list3)

		assert.NotEqual(t, list1, list2)
		assert.NotEqual(t, list1, list3)
		assert.NotEmpty(t, list2, list3)
	})

	t.Run("[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World", "Hi", "A-yon"}
		list2 := slices.Clone(list1)
		list3 := slices.Clone(list1)

		Shuffle(list2)
		Shuffle(list3)

		assert.NotEqual(t, list1, list2)
		assert.NotEqual(t, list1, list3)
		assert.NotEmpty(t, list2, list3)
	})
}

func TestOrderBy(t *testing.T) {
	t.Run("[]map[string]any", func(t *testing.T) {
		list1 := []map[string]any{
			{
				"id":  "world",
				"age": 53,
			},
			{
				"id":  "ayon",
				"age": 28,
			},
		}
		list2 := OrderBy(list1, "age", "asc")
		list3 := OrderBy(list2, "age", "desc")
		list4 := OrderBy(list1, "id", "asc")
		list5 := OrderBy(list2, "id", "desc")
		list6 := OrderBy(list1, "name", "desc")

		assert.Equal(t, []map[string]any{
			{
				"id":  "ayon",
				"age": 28,
			},
			{
				"id":  "world",
				"age": 53,
			},
		}, list2)
		assert.Equal(t, list1, list3)
		assert.Equal(t, list2, list4)
		assert.Equal(t, list1, list5)
		assert.Equal(t, list1, list6)
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("[]map[string]string->map[string][]map[string]string", func(t *testing.T) {
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
		record := GroupBy(list, func(item map[string]string, _ int) string {
			return item["group"]
		})

		assert.Equal(t, map[string][]map[string]string{
			"world": {
				{
					"group": "world",
					"tag":   "A",
				},
			},
			"room": {
				{
					"group": "room",
					"tag":   "B",
				},
				{
					"group": "room",
					"tag":   "C",
				},
			},
		}, record)

	})
}

func TestDiff(t *testing.T) {
	t.Run("...[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2, 3, 4, 5}
		list2 := []int{2, 3, 4, 5, 6, 7}
		list3 := Diff(list1, list2)

		assert.Equal(t, []int{0, 1}, list3)
	})

	t.Run("...[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World"}
		list2 := []string{"Hi", "World"}
		list3 := Diff(list1, list2)

		assert.Equal(t, []string{"Hello"}, list3)
	})
}

func TestXor(t *testing.T) {
	t.Run("...[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2, 3, 4, 5}
		list2 := []int{2, 3, 4, 5, 6, 7}
		list3 := Xor(list1, list2)

		assert.Equal(t, []int{0, 1, 6, 7}, list3)
	})

	t.Run("...[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World"}
		list2 := []string{"Hi", "World"}
		list3 := Xor(list1, list2)

		assert.Equal(t, []string{"Hello", "Hi"}, list3)
	})
}

func TestUnion(t *testing.T) {
	t.Run("...[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2, 3, 4, 5}
		list2 := []int{2, 3, 4, 5, 6, 7}
		list3 := Union(list1, list2)

		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7}, list3)
	})

	t.Run("...[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World"}
		list2 := []string{"Hi", "World"}
		list3 := Union(list1, list2)

		assert.Equal(t, []string{"Hello", "World", "Hi"}, list3)
	})
}

func TestIntersect(t *testing.T) {
	t.Run("...[]int", func(t *testing.T) {
		list1 := []int{0, 1, 2, 3, 4, 5}
		list2 := []int{2, 3, 4, 5, 6, 7}
		list3 := Intersect(list1, list2)

		assert.Equal(t, []int{2, 3, 4, 5}, list3)
	})

	t.Run("...[]string", func(t *testing.T) {
		list1 := []string{"Hello", "World"}
		list2 := []string{"Hi", "World"}
		list3 := Intersect(list1, list2)

		assert.Equal(t, []string{"World"}, list3)
	})
}
