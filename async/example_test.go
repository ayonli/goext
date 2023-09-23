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

func ExampleWaitTimeout() {
	res1, err1 := async.WaitTimeout(func() (string, error) {
		return "Hello, World!", nil
	}, time.Millisecond*10)
	res2, err2 := async.WaitTimeout(func() (string, error) {
		time.Sleep(time.Millisecond * 20)
		return "Hello, World!", nil
	}, time.Millisecond*10)

	fmt.Println(res1)
	fmt.Println(err1)
	fmt.Printf("%#v\n", res2)
	fmt.Println(err2)
	// Output:
	// Hello, World!
	// <nil>
	// ""
	// context deadline exceeded
}

func ExampleWaitAfter() {
	start := time.Now()
	res, err := async.WaitAfter(func() (string, error) {
		return "Hello, World!", nil
	}, time.Millisecond*10)

	fmt.Println(res)
	fmt.Println(err)
	fmt.Println(time.Since(start).Milliseconds() >= 10) // may exceed 10 due to context change
	// Output:
	// Hello, World!
	// <nil>
	// true
}

func ExampleWaitUntil() {
	result := 0
	async.WaitUntil(func() bool {
		result++
		return result == 10
	})
	fmt.Println(result)
	// Output:
	// 10
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
	// Output:
	// 2
}

func ExampleAsyncTask() {
	task := &async.AsyncTask[string]{}

	go func() {
		task.Resolve("Hello, World!")
	}()

	res, _ := task.Result()

	fmt.Println(res)
	// Output:
	// Hello, World!
}

func ExampleAsyncTask_Resolve() {
	task := &async.AsyncTask[string]{}

	go func() {
		task.Resolve("Hello, World!")

		// Resolve and Reject can only be called once and once one of them has been called,
		// calling the other will not effect.
		task.Reject(errors.New("something went wrong"))
	}()

	res, err := task.Result()

	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// Hello, World!
	// <nil>
}

func ExampleAsyncTask_Reject() {
	task := &async.AsyncTask[string]{}

	go func() {
		task.Reject(errors.New("something went wrong"))

		// Resolve and Reject can only be called once and once one of them has been called,
		// calling the other will not effect.
		task.Resolve("Hello, World!")
	}()

	res, err := task.Result()

	fmt.Printf("%#v\n", res)
	fmt.Println(err)
	// Output:
	// ""
	// something went wrong
}

func ExampleAsyncTask_Result() {
	task := &async.AsyncTask[string]{}

	go func() {
		task.Resolve("Hello, World!")
	}()

	res1, _ := task.Result()
	res2, _ := task.Result() // successive calls returns the same result

	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// Hello, World!
	// Hello, World!
}
