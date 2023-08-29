package goext_test

import (
	"errors"
	"fmt"

	"github.com/ayonli/goext"
)

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

func ExampleWrap() {
	texture := func(good bool) string {
		if !good {
			panic("something went wrong")
		}

		return "everything looks fine"
	}

	call := goext.Wrap(func(args ...any) string {
		text := texture(args[0].(bool))
		return text
	})

	res, err := call(true)

	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// everything looks fine
	// <nil>
}
