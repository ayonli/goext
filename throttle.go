package goext

import (
	"sync"
	"time"

	"github.com/ayonli/goext/async"
	"github.com/ayonli/goext/collections"
)

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
