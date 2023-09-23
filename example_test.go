package goext_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/ayonli/goext"
)

func ExampleReadAll() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	fmt.Println(goext.ReadAll(ch))
	// Output:
	// [1 2 3]
}

func ExampleOk() {
	texture := func(good bool) (string, error) {
		if !good {
			return "", errors.New("something went wrong")
		}

		return "everything looks fine", nil
	}

	res, err := goext.Try(func() string {
		text := goext.Ok(texture(true))
		return text
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	// Output:
	// everything looks fine
}

func ExampleOk_error() {
	texture := func(good bool) (string, error) {
		if !good {
			return "", errors.New("something went wrong")
		}

		return "everything looks fine", nil
	}

	res, err := goext.Try(func() string {
		text := goext.Ok(texture(false))
		return text
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	// Output:
	// something went wrong
}

func ExampleTry() {
	texture := func(good bool) string {
		if !good {
			panic("something went wrong")
		}

		return "everything looks fine"
	}

	res, err := goext.Try(func() string {
		text := texture(true)
		return text
	})

	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// everything looks fine
	// <nil>
}

func ExampleTry_panicString() {
	texture := func(good bool) string {
		if !good {
			panic("something went wrong")
		}

		return "everything looks fine"
	}

	res, err := goext.Try(func() string {
		text := texture(false)
		return text
	})

	fmt.Printf("%#v\n", res)
	fmt.Println(err)
	// Output:
	// ""
	// something went wrong
}

func ExampleTry_panicError() {
	texture := func(good bool) string {
		if !good {
			panic(errors.New("something went wrong"))
		}

		return "everything looks fine"
	}

	res, err := goext.Try(func() string {
		text := texture(false)
		return text
	})

	fmt.Printf("%#v\n", res)
	fmt.Println(err)
	// Output:
	// ""
	// something went wrong
}

func ExampleTry_panicAny() {
	texture := func(good bool) string {
		if !good {
			panic(1)
		}

		return "everything looks fine"
	}

	res, err := goext.Try(func() string {
		text := texture(false)
		return text
	})

	fmt.Printf("%#v\n", res)
	fmt.Println(err)
	// Output:
	// ""
	// 1
}

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
