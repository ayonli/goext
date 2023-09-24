package goext_test

import (
	"fmt"
	"time"

	"github.com/ayonli/goext"
)

func ExampleThrottle_withoutKey() {
	fn := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "")

	fmt.Println(fn(1))
	fmt.Println(fn(2))

	time.Sleep(time.Millisecond * 6)
	fmt.Println(fn(3))

	// Output:
	// 2 <nil>
	// 2 <nil>
	// 6 <nil>
}

func ExampleThrottle_withKey() {
	res1, err1 := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "foo")(1)
	res2, err2 := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "foo")(2)

	fmt.Println(res1, err1)
	fmt.Println(res2, err2)

	time.Sleep(time.Millisecond * 6)
	res3, err3 := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "foo")(3)
	fmt.Println(res3, err3)

	// Output:
	// 2 <nil>
	// 2 <nil>
	// 6 <nil>
}
