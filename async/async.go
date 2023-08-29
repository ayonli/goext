// Package async provides functions to run functions in other goroutines and wait for their results.
package async

import (
	"errors"
	"fmt"
	"slices"

	"github.com/ayonli/goext/slicex"
)

type indexedResult[R any] struct {
	index  int
	value  R
	reason error
}

// WaitResult represents a single result of the function passed to the `WaitAllSettled()` function.
type WaitResult[R any] struct {
	Value R
	Error error
}

// Wait runs the given function in another goroutine and waits its return value.
func Wait[R any](fn func() (R, error)) (R, error) {
	resChan := make(chan R)
	errChan := make(chan error)

	go func() {
		res, err := fn()

		if err != nil {
			errChan <- err
		} else {
			resChan <- res
		}
	}()

	select {
	case res := <-resChan:
		return res, nil
	case err := <-errChan:
		return *new(R), err
	}
}

// WaitRace runs a series of functions in different goroutines and wait for anyone returns, either
// with or without error.
func WaitRace[F func() (R, error), R any](fns ...F) (R, error) {
	resChan := make(chan R)
	errChan := make(chan error)

	for _, fn := range fns {
		go func(fn F) {
			res, err := fn()

			if err != nil {
				errChan <- err
			} else {
				resChan <- res
			}
		}(fn)
	}

	select {
	case res := <-resChan:
		return res, nil
	case err := <-errChan:
		return *new(R), err
	}
}

// WaitAny runs a series of functions in different goroutines and wait for anyone returns
// successfully without error.
//
// If all functions failed (returned with error), all the errors will be grouped in a single slice
// and ordered accordingly.
func WaitAny[F func() (R, error), R any](fns ...F) (R, []error) {
	resultChan := make(chan indexedResult[R])
	limit := len(fns)

	for i, fn := range fns {
		go func(fn F, i int) {
			res, err := fn()

			if err != nil {
				resultChan <- indexedResult[R]{
					index:  i,
					value:  *new(R),
					reason: err,
				}
			} else {
				resultChan <- indexedResult[R]{
					index:  i,
					value:  res,
					reason: nil,
				}
			}
		}(fn, i)
	}

	results := []indexedResult[R]{}

	for len(results) < limit {
		result := <-resultChan

		if result.reason == nil {
			return result.value, nil
		} else {
			results = append(results, result)
		}
	}

	slices.SortStableFunc(results, func(a indexedResult[R], b indexedResult[R]) int {
		return a.index - b.index
	})

	return *new(R), slicex.Map(results, func(item indexedResult[R], _ int) error {
		return item.reason
	})
}

// WaitAll runs a series of functions in different goroutines and wait for all return successfully
// or anyone fails.
//
// If all functions succeed (returned without error), all the results will be grouped in a single
// slice and ordered accordingly.
func WaitAll[F func() (R, error), R any](fns ...F) ([]R, error) {
	resultChan := make(chan indexedResult[R])
	limit := len(fns)

	for i, fn := range fns {
		go func(fn F, i int) {
			res, err := fn()

			if err != nil {
				resultChan <- indexedResult[R]{
					index:  i,
					value:  *new(R),
					reason: err,
				}
			} else {
				resultChan <- indexedResult[R]{
					index:  i,
					value:  res,
					reason: nil,
				}
			}
		}(fn, i)
	}

	results := []indexedResult[R]{}

	for len(results) < limit {
		result := <-resultChan

		if result.reason != nil {
			return nil, result.reason
		} else {
			results = append(results, result)
		}
	}

	slices.SortStableFunc(results, func(a indexedResult[R], b indexedResult[R]) int {
		return a.index - b.index
	})

	return slicex.Map(results, func(item indexedResult[R], _ int) R {
		return item.value
	}), nil
}

// WaitAllSettled runs a series of functions in different goroutines and wait for all of them are
// settled, all the results are grouped in a single slice and ordered accordingly.
func WaitAllSettled[F func() (R, error), R any](fns ...F) []WaitResult[R] {
	resultChan := make(chan indexedResult[R])
	limit := len(fns)

	for i, fn := range fns {
		go func(fn F, i int) {
			res, err := fn()

			if err != nil {
				resultChan <- indexedResult[R]{
					index:  i,
					value:  *new(R),
					reason: err,
				}
			} else {
				resultChan <- indexedResult[R]{
					index:  i,
					value:  res,
					reason: nil,
				}
			}
		}(fn, i)
	}

	results := []indexedResult[R]{}

	for len(results) < limit {
		result := <-resultChan
		results = append(results, result)
	}

	slices.SortStableFunc(results, func(a indexedResult[R], b indexedResult[R]) int {
		return a.index - b.index
	})

	return slicex.Map(results, func(item indexedResult[R], _ int) WaitResult[R] {
		return WaitResult[R]{
			Value: item.value,
			Error: item.reason,
		}
	})
}

// Try runs a function in a safe context where if it or what's inside it panics, the panic reason
// can be caught and returned as a normal error.
//
// Deprecated: use goext.Try() instead.
func Try[R any](fn func() (R, error)) (res R, err error) {
	resChan := make(chan R)
	errChan := make(chan error)

	go func() {
		defer func() {
			if re := recover(); re != nil {
				if _err, ok := re.(error); ok {
					errChan <- _err
				} else if str, ok := re.(string); ok {
					errChan <- errors.New(str)
				} else {
					errChan <- errors.New(fmt.Sprint(re))
				}
			}
		}()

		res, err := fn()

		if err != nil {
			errChan <- err
		} else {
			resChan <- res
		}
	}()

	select {
	case res := <-resChan:
		return res, nil
	case err := <-errChan:
		return *new(R), err
	}
}
