# GoExt

Extended functionalities that are missing in the Go standard library but frequently used in other
languages.

Especially for JavaScript developers, these packages should make us feel right at home.

## Install

```sh
go get github.com/ayonli/goext
```

## Functions

- [goext.ReadAll](#goextreadall)
- [goext.Ok](#goextok)
- [goext.Try](#goexttry)
- [goext.Queue](#goextqueue)
- [goext.Throttle](#goextthrottle)

### goext.ReadAll

```go
func ReadAll[T any](ch <-chan T) []T
```

ReadAll reads all values from the channel at once.

---

### goext.Ok

```go
func Ok[R any](res R, err error) R
```

Ok asserts a typical Golang function call which returns a result and an error is successful and
returns the result, it panics if the return error is not nil. This function should be composite
with the `goext.Try()` function, allowing the program to bubble the error and catch it from
outside.

**Example**

```go
_, err := goext.Try(func () int {
    res1 := goext.Ok(someCall())
    res2 := goext.Ok(anotherCall())
    return 0
})
```

---

### goext.Try

```go
func Try[R any](fn func() R) (res R, err error)
```

Try runs a function in a safe context where if it or what's inside it panics, the panic reason
can be caught and returned as a normal error.

**Example**

```go
_, err := goext.Try(func () int {
    res1 := goext.Ok(someCall())
    res2 := goext.Ok(anotherCall())
    return 0
})
```

---

### goext.Queue

```go
func Queue[T any](handler func(data T), bufferSize int) IQueue[T]
```

Queue processes data sequentially by the given `handler` function and prevents concurrency
conflicts, it returns a queue instance that we can push data into.

`bufferSize` is the maximum capacity of the underlying channel, once reached, the push
operation will block until there is new space available. Bu default, this option is not set and
use a non-buffered channel instead.

---

### goext.Throttle

```go
func Throttle[A any, R any, Fn func(arg A) (R, error)](
    handler Fn,
    duration time.Duration,
    forKey string,
    noWait bool,
) Fn
```

Creates a throttled function that will only be run once in a certain amount of time.

If a subsequent call happens within the `duration`, the previous result will be returned and
the `handler` function will not be invoked.

If `forKey` is provided, use the throttle strategy for the given key, this will keep the
result in a global cache, binding new `handler` function for the same key will result in the
same result as the previous, unless the duration has passed. This mechanism guarantees that both
creating the throttled function in function scopes and overwriting the handler are possible.

If `noWait` is turned on, respond with the last cache (if available) immediately, even if it has
expired, and update the cache in the background.

## Sub-packages

- **[async](https://pkg.go.dev/github.com/ayonli/goext/async)** (Since v0.2.0)
    Package async provides functions to run functions in other goroutines and wait for their results.
- **[mathx](https://pkg.go.dev/github.com/ayonli/goext/mathx)**
    Additional functions for math calculation that are missing in the standard library.
- **[stringx](https://pkg.go.dev/github.com/ayonli/goext/stringx)**
    Additional functions for string processing that are missing in the standard library.
    - **[mbstring](https://pkg.go.dev/github.com/ayonli/goext/stringx/mbstring)**
        Additional functions for processing strings in multi-byte sequence.
- **[slicex](https://pkg.go.dev/github.com/ayonli/goext/slicex)**
    Additional functions for playing with slices and reduce mistakes.
- **[mapx](https://pkg.go.dev/github.com/ayonli/goext/mapx)**
    Additional functions for dealing with maps.
- **[structx](https://pkg.go.dev/github.com/ayonli/goext/structx)** (Since v0.3.0)
    Functions used to manipulate structs.
- **[oop](https://pkg.go.dev/github.com/ayonli/goext/oop)**
    Object-oriented abstract wrappers for basic data structures.
    - `String` is an object-oriented abstract that works around multi-byte strings.
    - `List` is an objected-oriented abstract that works around the slice and acts as a dynamic array.
- **[collections](https://pkg.go.dev/github.com/ayonli/goext/collections)**
    Object-oriented abstract wrappers for basic types.
    - `Set` is an object-oriented collection that stores unique items and is thread-safe.
    - `Map` is an object-oriented collection of map with ordered keys and thread-safe by default.
    - `CiMap` Thread-safe case-insensitive map, keys are case-insensitive.
    - `BiMap` Thread-safe bi-directional map, keys and values are unique and map to each other.
