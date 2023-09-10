package async_test

import (
	"errors"
	"fmt"
	"runtime"
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
		time.Sleep(time.Microsecond * 15)
		return "Hi, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 30)
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
		time.Sleep(time.Microsecond * 30)
		return "Hello, World!", nil
	}, func() (string, error) {
		time.Sleep(time.Microsecond * 15)
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

func ExampleQueue() {
	out := make(chan []string)
	list := []string{}
	push := async.Queue(func(str string) (fin bool) {
		list = append(list, str)
		fin = len(list) == 2

		if fin {
			// The order of the `list` is not stable, but we can guarantee that all two strings have
			// been appended and stored to the `list`.
			//
			// Without concurrency control, we may end up `list` only has one item left or
			// len(list) == 2 but when we trying to print or send it, it becomes 1.
			out <- list
		}

		return fin
	})

	go func() {
		push("foo")
	}()

	go func() {
		push("bar")
	}()

	fmt.Println(len(<-out))
	fmt.Println(runtime.NumGoroutine())
	// Output:
	// 2
	// 2
}
