package mapx_test

import (
	"fmt"

	"github.com/ayonli/goext/mapx"
)

func ExampleAssign() {
	m1 := mapx.Assign(map[string]string{}, map[string]string{
		"foo": "Hello",
	})
	m2 := mapx.Assign(map[string]string{}, m1, map[string]string{
		"bar": "World",
	})

	fmt.Println(m1)
	fmt.Println(m2)
	// Output:
	// map[foo:Hello]
	// map[bar:World foo:Hello]
}

func ExamplePatch() {
	m1 := mapx.Patch(map[string]string{}, map[string]string{
		"foo": "Hello",
	})
	m2 := mapx.Patch(map[string]string{}, m1, map[string]string{
		"foo": "Hi",
		"bar": "World",
	})

	fmt.Println(m1)
	fmt.Println(m2)
	// Output:
	// map[foo:Hello]
	// map[bar:World foo:Hello]
}

func ExampleKeys_string() {
	m := map[string]string{
		"foo": "Hello",
		"bar": "World",
	}
	keys := mapx.Keys(m)

	fmt.Println(keys)
	// Output:
	// [bar foo]
}

func ExampleKeys_int() {
	m := map[int]string{
		0: "Hello",
		1: "World",
	}
	keys := mapx.Keys(m)

	fmt.Println(keys)
	// Output:
	// [0 1]
}

func ExampleValues() {
	m := map[string]string{
		"foo": "Hello",
		"bar": "World",
	}
	values := mapx.Values(m)

	fmt.Println(values)
	// Output:
	// [World Hello]
}

func ExamplePick_string() {
	m1 := map[string]string{
		"foo": "Hello",
		"bar": "World",
	}
	m2 := mapx.Pick(m1, []string{"foo"})

	fmt.Println(m2)
	// Output:
	// map[foo:Hello]
}

func ExamplePick_int() {
	m1 := map[int]string{
		0: "Hello",
		1: "World",
	}
	m2 := mapx.Pick(m1, []int{0})

	fmt.Println(m2)
	// Output:
	// map[0:Hello]
}

func ExampleOmit_string() {
	m1 := map[string]string{
		"foo": "Hello",
		"bar": "World",
	}
	m2 := mapx.Omit(m1, []string{"bar"})

	fmt.Println(m2)
	// Output:
	// map[foo:Hello]
}

func ExampleOmit_int() {
	m1 := map[int]string{
		0: "Hello",
		1: "World",
	}
	m2 := mapx.Omit(m1, []int{1})

	fmt.Println(m2)
	// Output:
	// map[0:Hello]
}
