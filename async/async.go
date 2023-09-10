// Package async provides functions to run functions in other goroutines and wait for their results.
package async

import (
	"context"
	"time"

	"github.com/ayonli/goext/slicex"
)

// WaitResult represents a single result of the function passed to the `Wait` family functions.
type WaitResult[R any] struct {
	Value R
	Error error
}

type indexedResult[R any] struct {
	index  int
	result WaitResult[R]
}

// Wait runs the given function in another goroutine and waits its return value.
func Wait[R any](fn func() (R, error)) (R, error) {
	channel := make(chan WaitResult[R])

	go func() {
		res, err := fn()
		channel <- WaitResult[R]{Value: res, Error: err}
	}()

	res := <-channel
	return res.Value, res.Error
}

// WaitRace runs a series of functions in different goroutines and wait for anyone returns, either
// with or without error.
func WaitRace[F func() (R, error), R any](fns ...F) (R, error) {
	channel := make(chan WaitResult[R], len(fns))

	for _, fn := range fns {
		go func(fn F) {
			res, err := fn()
			channel <- WaitResult[R]{Value: res, Error: err}
		}(fn)
	}

	res := <-channel
	return res.Value, res.Error
}

// WaitAny runs a series of functions in different goroutines and wait for anyone returns
// successfully without error.
//
// If all functions failed (returned with error), all the errors will be grouped in a single slice
// and ordered accordingly.
func WaitAny[F func() (R, error), R any](fns ...F) (R, []error) {
	limit := len(fns)
	channel := make(chan indexedResult[R], limit)
	results := make([]WaitResult[R], limit)
	count := 0

	for i, fn := range fns {
		go func(fn F, i int) {
			res, err := fn()
			channel <- indexedResult[R]{
				index:  i,
				result: WaitResult[R]{Value: res, Error: err},
			}
		}(fn, i)
	}

	for count < limit {
		res := <-channel

		if res.result.Error == nil {
			return res.result.Value, nil
		} else {
			results[res.index] = res.result
		}

		count++
	}

	return *new(R), slicex.Map(results, func(item WaitResult[R], _ int) error {
		return item.Error
	})
}

// WaitAll runs a series of functions in different goroutines and wait for all return successfully
// or anyone fails.
//
// If all functions succeed (returned without error), all the results will be grouped in a single
// slice and ordered accordingly.
func WaitAll[F func() (R, error), R any](fns ...F) ([]R, error) {
	limit := len(fns)
	channel := make(chan indexedResult[R], limit)
	results := make([]WaitResult[R], limit)
	count := 0

	for i, fn := range fns {
		go func(fn F, i int) {
			res, err := fn()
			channel <- indexedResult[R]{
				index:  i,
				result: WaitResult[R]{Value: res, Error: err},
			}
		}(fn, i)
	}

	for count < limit {
		res := <-channel

		if res.result.Error != nil {
			return nil, res.result.Error
		} else {
			results[res.index] = res.result
		}

		count++
	}

	return slicex.Map(results, func(item WaitResult[R], _ int) R {
		return item.Value
	}), nil
}

// WaitAllSettled runs a series of functions in different goroutines and wait for all of them are
// settled, all the results are grouped in a single slice and ordered accordingly.
func WaitAllSettled[F func() (R, error), R any](fns ...F) []WaitResult[R] {
	limit := len(fns)
	channel := make(chan indexedResult[R], limit)
	results := make([]WaitResult[R], limit)
	count := 0

	for i, fn := range fns {
		go func(fn F, i int) {
			res, err := fn()
			channel <- indexedResult[R]{
				index:  i,
				result: WaitResult[R]{Value: res, Error: err},
			}
		}(fn, i)
	}

	for count < limit {
		res := <-channel
		results[res.index] = res.result
		count++
	}

	return results
}

// WaitTimeout runs the given function in another goroutine and shall return its result before the
// timeout limit.
func WaitTimeout[R any](fn func() (R, error), duration time.Duration) (R, error) {
	channel := make(chan WaitResult[R], 1)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	go func() {
		res, err := fn()
		channel <- WaitResult[R]{Value: res, Error: err}
	}()

	select {
	case res := <-channel:
		return res.Value, res.Error
	case <-ctx.Done():
		return *new(R), ctx.Err()
	}
}

// WaitAfter runs the given function in another goroutine and returns its result only after the
// given duration.
func WaitAfter[R any](fn func() (R, error), duration time.Duration) (R, error) {
	results := WaitAllSettled(fn, func() (R, error) {
		time.Sleep(duration)
		return *new(R), nil
	})
	result := results[0]
	return result.Value, result.Error
}

// Queue processes data sequentially by the given callback function that prevents concurrency
// conflicts, it returns a new function that pushes data into the queue.
//
// The callback function returns a boolean value indicates whether the queue has finished, once true,
// the internal channel will be closed and no more data shall be pushed.
func Queue[T any](callback func(data T) (fin bool)) func(data T) {
	c := make(chan T)

	go func() {
		for !callback(<-c) {
		}

		close(c)
	}()

	return func(data T) {
		c <- data
	}
}
