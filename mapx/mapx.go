// Additional functions for dealing with maps.
package mapx

import (
	"maps"
	"slices"

	"github.com/ayonli/goext/slicex"
)

// Copies one or more items from the source maps to the target map. The later key-value pairs
// override the existing ones or the ones before them.
//
// This function mutates the target map and returns it.
func Assign[M ~map[K]V, K comparable, V any](target M, sources ...M) M {
	for _, source := range sources {
		if source != nil {
			maps.Copy(target, source)
		}
	}

	return target
}

// Copies the key-value pairs that are presented in the source maps but are missing in the target
// map into the target map, later pairs are skipped if the same key already exists.
//
// This function mutates the target map and returns it.
func Patch[M ~map[K]V, K comparable, V any](target M, sources ...M) M {
	for _, source := range sources {
		for k, v := range source {
			_, ok := target[k]

			if !ok {
				target[k] = v
			}
		}
	}

	return target
}

// Returns the keys of the given map.
//
// Keys are sorted in ascending order if they are strings or integers.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	size := len(m)
	keys := make([]K, size)
	idx := 0

	for k := range m {
		keys[idx] = k
		idx++
	}

	slices.SortStableFunc(keys, slicex.CompareItems)
	return keys
}

// Returns the values of the given map.
//
// Values are ordered according to the keys' order returned by `Keys()`.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	keys := Keys(m)
	values := make([]V, len(keys))

	for i, k := range keys {
		values[i] = m[k]
	}

	return values
}

// Executes a provided function once for each key-value pair.
//
// This function adds a closure context around each item looped, may be useful for preventing
// variable pollution.
func ForEach[M ~map[K]V, K comparable, V any](m M, fn func(value V, key K)) {
	for key, value := range m {
		fn(value, key)
	}
}

// Creates a new map based on the original map but only contains the specified keys.
func Pick[M ~map[K]V, K comparable, V any](original M, keys []K) M {
	newMap := M{}

	for _, key := range keys {
		value, ok := original[key]

		if ok {
			newMap[key] = value
		}
	}

	return newMap
}

// Creates a new map based on the original map but without the specified keys.
func Omit[M ~map[K]V, K comparable, V any](original M, keys []K) M {
	allKeys := Keys(original)
	keptKeys := slicex.Diff(allKeys, keys)
	return Pick(original, keptKeys)
}
