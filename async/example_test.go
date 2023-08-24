package async_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/ayonli/goext/async"
)

func ExampleWait() {
	res, err := async.Wait(func() (string, error) {
		// this function runs in another goroutine
		return "Hello, World!", nil
	})
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// Hello, World!
	// <nil>
}

func ExampleWait_error() {
	res, err := async.Wait(func() (string, error) {
		// this function runs in another goroutine
		return "", errors.New("something went wrong")
	})
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	//
	// something went wrong
}

func ExampleWaitRace() {
	res, err := async.WaitRace(func() (string, error) {
		time.Sleep(time.Microsecond * 1)
		return "Hello, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 10)
		return "Hi, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 20)
		return "", errors.New("something went wrong")
	})

	fmt.Println(res)
	fmt.Printf("%#v\n", err)
	// Output:
	// Hello, World!
	// <nil>
}

func ExampleWaitRace_error() {
	res, errors := async.WaitRace(func() (string, error) {
		time.Sleep(time.Microsecond * 20)
		return "Hello, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 10)
		return "Hi, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 1)
		return "", errors.New("something went wrong")
	})

	fmt.Printf("%#v\n", res)
	fmt.Println(errors)
	// Output:
	// ""
	// something went wrong
}

func ExampleWaitAny() {
	res, errors := async.WaitAny(func() (string, error) {
		time.Sleep(time.Microsecond * 1)
		return "Hello, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 10)
		return "Hi, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 20)
		return "", errors.New("something went wrong")
	})

	fmt.Println(res)
	fmt.Printf("%#v\n", errors)
	// Output:
	// Hello, World!
	// []error(nil)
}

func ExampleWaitAny_error() {
	res, errors := async.WaitAny(func() (string, error) {
		time.Sleep(time.Microsecond * 20)
		return "", errors.New("something went wrong")
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 10)
		return "", errors.New("something went wrong")
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 1)
		return "", errors.New("something went wrong")
	})

	fmt.Printf("%#v\n", res)
	fmt.Println(errors)
	// Output:
	// ""
	// [something went wrong something went wrong something went wrong]
}

func ExampleWaitAll() {
	results, err := async.WaitAll(func() (string, error) {
		time.Sleep(time.Microsecond * 1)
		return "Hello, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 10)
		return "Hi, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 20)
		return "你好，世界！", nil
	})

	fmt.Println(results)
	fmt.Println(err)
	// Output:
	// [Hello, World! Hi, World! 你好，世界！]
	// <nil>
}

func ExampleWaitAll_error() {
	results, err := async.WaitAll(func() (string, error) {
		time.Sleep(time.Microsecond * 20)
		return "Hello, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 10)
		return "Hi, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 1)
		return "你好，世界！", errors.New("something went wrong")
	})

	fmt.Printf("%#v\n", results)
	fmt.Println(err)
	// Output:
	// []string(nil)
	// something went wrong
}

func ExampleWaitAllSettled() {
	results := async.WaitAllSettled(func() (string, error) {
		time.Sleep(time.Microsecond * 1)
		return "Hello, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 10)
		return "Hi, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 20)
		return "", errors.New("something went wrong")
	})

	for _, result := range results {
		if result.Error != nil {
			fmt.Println(result.Error)
		} else {
			fmt.Println(result.Value)
		}
	}
	// Output:
	// Hello, World!
	// Hi, World!
	// something went wrong
}

func ExampleTry() {
	texture := func(good bool) string {
		if !good {
			panic("something went wrong")
		}

		return "everything looks fine"
	}

	_, err := async.Try(func() (string, error) {
		return texture(false), nil
	})

	fmt.Println(err)
	// Output:
	// something went wrong
}
