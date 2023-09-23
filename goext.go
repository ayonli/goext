package goext

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ayonli/goext/async"
	"github.com/ayonli/goext/collections"
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

type throttleCache[R any] struct {
	key     string
	mut     *sync.Mutex
	expires time.Time
	result  *async.WaitResult[R]
	pending *async.AsyncTask[R]
}

var throttleCaches = &collections.Map[string, any]{}

// Creates a throttled function that will only be run once in a certain amount of time.
//
// If a subsequent call happens within the `duration`, the previous result will be returned and
// the `handler` function will not be invoked.
//
// If the `handler` function returns a promise, and two or more calls happen simultaneously,
// the later calls will try to resolve with the previous result immediately instead of waiting
// the pending call to complete.
//
// If `forKey` is provided, use the throttle strategy for the given key, this will keep the
// result in a global cache, binding new `handler` function for the same key will result in the
// same result as the previous, unless the duration has passed. This mechanism guarantees that both
// creating the throttled function in function scopes and overwriting the handler are possible.
func Throttle[A any, R any, Fn func(arg A) (R, error)](
	handler Fn,
	duration time.Duration,
	forKey string,
) Fn {
	handleCall := func(cache *throttleCache[R], arg A) (R, error) {
		cache.mut.Lock()
		defer cache.mut.Unlock()

		if cache.result != nil && (cache.pending != nil || cache.expires.After(time.Now())) {
			if cache.result.Error != nil {
				return *new(R), cache.result.Error
			} else {
				return cache.result.Value, nil
			}
		} else if cache.pending != nil {
			return cache.pending.Result()
		}

		cache.pending = &async.AsyncTask[R]{}

		go func() {
			res, err := handler(arg)

			if err != nil {
				cache.pending.Reject(err)
			} else {
				cache.pending.Resolve(res)
			}
		}()

		val, err := cache.pending.Result()
		cache.result = &async.WaitResult[R]{Value: val, Error: err}
		cache.expires = time.Now().Add(duration)
		cache.pending = nil

		return cache.result.Value, cache.result.Error
	}

	var cache *throttleCache[R]

	if forKey == "" {
		cache = &throttleCache[R]{key: "", mut: &sync.Mutex{}}
	} else {
		cache = (throttleCaches.EnsureGet(forKey, func() any {
			return any(&throttleCache[R]{key: forKey, mut: &sync.Mutex{}})
		})).(*throttleCache[R])
	}

	return func(arg A) (R, error) {
		return handleCall(cache, arg)
	}
}
