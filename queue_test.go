package goext_test

import (
	"errors"
	"fmt"

	"github.com/ayonli/goext"
)

func ExampleQueue() {
	out := make(chan []string)
	list := []string{}
	queue := goext.Queue(func(str string) {
		list = append(list, str)

		if len(list) == 2 {
			out <- list
		}
	}, 0)

	go func() {
		queue.Push("foo")
	}()

	go func() {
		queue.Push("bar")
	}()

	fmt.Println(len(<-out))
	queue.Close()
	// Output:
	// 2
}

func ExampleQueue_error() {
	out := make(chan error)
	queue := goext.Queue(func(str string) {
		if str == "error" {
			panic(errors.New("something went wrong"))
		}
	}, 0)

	queue.OnError(func(err error) {
		out <- err
	})

	queue.Push("error")

	err := <-out
	fmt.Println(err)
	// Output:
	// something went wrong
}
