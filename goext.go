package goext

import (
	"errors"
	"fmt"
)

// Ok asserts a typical Golang function call which returns a result and an error is successful and
// returns the result, it panics if the return error is not nil. This function should be composite
// with the `goext.Try()` function, allowing the program to bubble the error and catch it from
// outside.
//
// Example:
//
//	_, err := goext.Try(func () int {
//		res1 := goext.Ok(someCall())
//		res2 := goext.Ok(anotherCall())
//		return 0
//	})
func Ok[R any](res R, err error) R {
	if err == nil {
		return res
	} else {
		panic(err)
	}
}

// Try runs a function in a safe context where if it or what's inside it panics, the panic reason
// can be caught and returned as a normal error.
func Try[R any](fn func() R) (res R, err error) {
	defer func() {
		if re := recover(); re != nil {
			if _err, ok := re.(error); ok {
				err = _err
			} else if str, ok := re.(string); ok {
				err = errors.New(str)
			} else {
				err = errors.New(fmt.Sprint(re))
			}
		}
	}()

	return fn(), nil
}

// Wrap returns a new function that wraps the `goext.Try()`, rendering the new function already
// catchable.
//
// Deprecated: this function is not good.
func Wrap[R any](fn func(args ...any) R) func(args ...any) (R, error) {
	return func(args ...any) (R, error) {
		return Try(func() R {
			return fn(args...)
		})
	}
}
