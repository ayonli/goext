package goext_test

import (
	"errors"
	"fmt"

	"github.com/ayonli/goext"
)

func ExampleTry_success() {
	texture := func(good bool) string {
		if !good {
			panic("something went wrong")
		}

		return "everything looks fine"
	}

	res, err := goext.Try(func() (string, error) {
		return texture(true), nil
	})

	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// everything looks fine
	// <nil>
}

func ExampleTry_error() {
	res, err := goext.Try(func() (string, error) {
		return "", errors.New("something went wrong")
	})

	fmt.Printf("%#v\n", res)
	fmt.Println(err)
	// Output:
	// ""
	// something went wrong
}

func ExampleTry_panicString() {
	texture := func(good bool) string {
		if !good {
			panic("something went wrong")
		}

		return "everything looks fine"
	}

	res, err := goext.Try(func() (string, error) {
		return texture(false), nil
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

	res, err := goext.Try(func() (string, error) {
		return texture(false), nil
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

	res, err := goext.Try(func() (string, error) {
		return texture(false), nil
	})

	fmt.Printf("%#v\n", res)
	fmt.Println(err)
	// Output:
	// ""
	// 1
}
