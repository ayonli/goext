# GoExt

Extended functionalities that are missing in the Go standard library but frequently used in other
languages.

Especially for JavaScript developers, these packages should make us feel right at home.

## Install

```sh
go get github.com/ayonli/goext
```

## Functions

- `func Try[R any](fn func() (R, error)) (res R, err error)`

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
    - `Set` is an object-oriented collection that stores unique items.
    - `Map` is an object-oriented collection of map with ordered keys.
    - `CiMap` Case-insensitive map, keys are case-insensitive.
    - `BiMap` Bi-directional map, keys and values are unique and map to each other.
