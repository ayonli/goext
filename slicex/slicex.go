// Additional functions for playing with slices and reduce mistakes.
package slicex

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strconv"
	"strings"

	"github.com/ayonli/goext/mathx"
)

// Returns the item from the slice according to the given index.
//
// If `i < 0`, it returns the item counting from the end of the slice.
//
// If the given index doesn't contain a value (boundary exceeded), the zero-value of the given type
// and `false` will be returned.
func At[S ~[]T, T any](s S, i int) (T, bool) {
	if i < 0 {
		i = len(s) + i
	}

	if i < len(s) {
		return s[i], true
	}

	var empty T
	return empty, false
}

// Returns the last index at which a given item can be found in the slice, or -1 if it is not
// present. The slice is searched backwards.
func LastIndex[S ~[]T, T comparable](s S, item T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == item {
			return i
		}
	}

	return -1
}

// Counts the occurrence of the item against the given slice.
func Count[S ~[]T, T comparable](s S, item T) int {
	count := 0

	for _, v := range s {
		if v == item {
			count++
		}
	}

	return count
}

// Merges two or more slices into one and returns the new slice.
func Concat[S ~[]T, T any](sources ...S) S {
	sources = Filter(sources, func(item S, _ int) bool { return item != nil })
	lengths := Map(sources, func(item S, _ int) float64 { return float64(len(item)) })
	length := mathx.Sum(lengths...)
	list := make(S, int(length))
	idx := 0

	for _, items := range sources {
		for _, value := range items {
			list[idx] = value
			idx++
		}
	}

	return list
}

// Creates a new slice with all sub-slice items concatenated into it.
func Flat[S ~[]T, T any](original []S) S {
	original = Filter(original, func(item S, _ int) bool { return item != nil })
	lengths := Map(original, func(item S, _ int) float64 { return float64(len(item)) })
	length := mathx.Sum(lengths...)
	list := make(S, int(length))
	idx := 0

	for _, items := range original {
		for _, item := range items {
			list[idx] = item
			idx++
		}
	}

	return list
}

// Returns a shallow copy of a portion of the slice into a new slice selected from `start` to `end`
// (excluded).
//
// If `start < 0`, it will be calculated as `len(original) + start`.
//
// If `end < 0`, it will be calculated as `len(original) + end`.
//
// Unlike the `[:]` syntax which creates a new slice that shares the same underlying array with the
// old slice, this function creates a new slice with new underlying array and copies data from the
// old one to the new one that prevent side effect when modifying them.
func Slice[S ~[]T, T any](original S, start int, end int) S {
	limit := len(original)

	if start < 0 {
		start = limit + start
	}

	if end < 0 {
		end = limit + end
	}

	if end > limit {
		end = limit
	}

	if start >= end || start >= limit {
		return make(S, 0) // return an empty slice directly
	}

	// Create a new slice with a new underlying array, so that when the new slice are modified,
	// the old slice (and its underlying array) remains untouched.
	// And vice versa.
	part := make(S, end-start)
	copy(part, original[start:end])

	return part
}

// Breaks the original slice into smaller chunks according to the given delimiter.
func Split[S ~[]T, T comparable](original S, delimiter T) []S {
	chunks := []S{}
	limit := len(original)
	offset := 0

	for i := 0; i < limit; i++ {
		if original[i] == delimiter {
			chunks = append(chunks, original[offset:i])
			offset = i + 1
		}
	}

	if offset < limit {
		chunks = append(chunks, original[offset:limit])
	}

	return chunks
}

// Breaks the original slice into smaller chunks according to the given length.
func Chunk[S ~[]T, T any](original S, length int) []S {
	limit := len(original)
	size := int(math.Ceil(float64(limit) / float64(length)))
	chunks := make([]S, size)
	offset := 0
	idx := 0

	for offset < limit {
		chunks[idx] = Slice(original, offset, offset+length)
		offset += length
		idx++
	}

	return chunks
}

// Creates and returns a string by concatenating all of the items in the slice, separated by the
// specified separator string.
func Join[S ~[]T, T any](s S, sep string) string {
	return strings.Join(Map(s, func(item T, _ int) string {
		if value, ok := any(item).(int); ok {
			return strconv.Itoa(value)
		} else if value, ok := any(item).(string); ok {
			return value
		} else {
			return fmt.Sprint(item)
		}
	}), sep)
}

// Tests whether all items in the slice pass the test implemented by the provided function.
func Every[S ~[]T, T any](s S, fn func(item T, idx int) bool) bool {
	for idx, item := range s {
		ok := fn(item, idx)

		if !ok {
			return false
		}
	}

	return true
}

// Tests whether at least one item in the slice passes the test implemented by the provided
// function.
func Some[S ~[]T, T any](s S, fn func(item T, idx int) bool) bool {
	for idx, item := range s {
		if fn(item, idx) {
			return true
		}
	}

	return false
}

// Returns the first item in the provided slice that satisfies the provided testing function.
//
// If no value satisfies the testing function, the zero-value of the given type and `false` will be
// returned.
func Find[S ~[]T, T any](s S, fn func(item T, idx int) bool) (T, bool) {
	for idx, item := range s {
		if fn(item, idx) {
			return item, true
		}
	}

	var empty T
	return empty, false
}

// Iterates the slice in reverse order and returns the value of the first item that satisfies the
// provided testing function.
//
// If no item satisfies the testing function, the zero-value of the given type and `false` will be
// returned.
func FindLast[S ~[]T, T any](s S, fn func(item T, idx int) bool) (T, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		item := s[i]

		if fn(item, i) {
			return item, true
		}
	}

	var empty T
	return empty, false
}

// Returns the index of the first item in the slice that satisfies the provided testing function.
// If no item satisfies the testing function, -1 is returned.
//
// This function is similar to the `slices.IndexFunc()` from the standard library.
func FindIndex[S ~[]T, T any](s S, fn func(item T, idx int) bool) int {
	for idx, item := range s {
		ok := fn(item, idx)

		if ok {
			return idx
		}
	}

	return -1
}

// Iterates the slice in reverse order and returns the index of the first item that satisfies the
// provided testing function. If no item satisfies the testing function, -1 is returned.
func FindLastIndex[S ~[]T, T any](s S, fn func(item T, idx int) bool) int {
	for idx := len(s) - 1; idx >= 0; idx-- {
		item := s[idx]
		ok := fn(item, idx)

		if ok {
			return idx
		}
	}

	return -1
}

// Creates a shallow copy of a portion of a given slice, filtered down to just the items from the
// given slice that pass the test implemented by the provided function.
func Filter[S ~[]T, T any](original S, fn func(item T, idx int) bool) S {
	items := S{}

	for idx, item := range original {
		if fn(item, idx) {
			items = append(items, item)
		}
	}

	return items
}

// Creates a new slice based on the original slice and removes all the duplicated items.
func Uniq[S ~[]E, E comparable](original S) S {
	items := S{}

	for _, item := range original {
		if !slices.Contains(items, item) {
			items = append(items, item)
		}
	}

	return items
}

// Creates a new slice based on the original slice and remove all the duplicated items identified
// by the given key.
func UniqBy[S ~[]M, M ~map[K]V, K comparable, V comparable](original S, key K) S {
	ids := []V{}
	items := S{}

	for _, item := range original {
		if item == nil {
			continue
		}

		id, ok := item[key]

		if !ok {
			continue
		}

		if !slices.Contains(ids, id) {
			ids = append(ids, id)
			items = append(items, item)
		}
	}

	return items
}

// Executes a provided function once for each slice item.
//
// This function adds a closure context around each item looped, may be useful for preventing
// variable pollution.
func ForEach[S ~[]T, T any](s S, fn func(item T, idx int)) {
	for idx, item := range s {
		fn(item, idx)
	}
}

// Creates a new slice populated with the results of calling a provided function on every item in
// the original slice.
func Map[S ~[]T, T any, R any](original S, fn func(item T, idx int) R) []R {
	items := make([]R, len(original))

	for idx, item := range original {
		items[idx] = fn(item, idx)
	}

	return items
}

// Executes a user-supplied "reducer" callback function on each item of the slice, in order,
// passing in the return value from the calculation on the preceding item.
//
// The reducer walks through the slice item-by-item, at each step adding the current item to the
// result from the previous step (this result is the running sum of all the previous steps) — until
// there are no more items to add.
func Reduce[S ~[]T, T any, R any](s S, fn func(acc R, item T, idx int) R, acc R) R {
	for idx, item := range s {
		acc = fn(acc, item, idx)
	}

	return acc
}

// Removes the last item from the slice and returns that item.
//
// This function mutates the input slice.
func Pop[S ~[]T, T any](s *S) T {
	end := len(*s) - 1
	item := (*s)[end]
	*s = (*s)[0:end]
	return item
}

// Adds the specified items to the end of the slice and returns the new length of the slice.
//
// This function mutates the input slice.
func Push[S ~[]T, T any](s *S, items ...T) int {
	*s = append(*s, items...)
	return len(*s)
}

// Removes the first item from the slice and returns that item.
//
// This function mutates the input slice.
func Shift[S ~[]T, T any](s *S) T {
	item := (*s)[0]
	*s = (*s)[1:]
	return item
}

// Adds the specified items to the head of the slice and returns the new length of the slice.
//
// This function mutates the input slice.
func Unshift[S ~[]T, T any](s *S, items ...T) int {
	*s = append(items, (*s)...)
	return len(*s)
}

// Reorganizes the items in the given slice in random order.
//
// This function mutate the input slice.
func Shuffle[S ~[]T, T any](s S) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// This function is used as the compare function for sorting functions.
func CompareItems[T any](a, b T) int {
	switch _a := any(a).(type) {
	case int:
		_b := any(b).(int)
		return _a - _b
	case int16:
		_b := any(b).(int16)
		return int(_a - _b)
	case int32:
		_b := any(b).(int32)
		return int(_a - _b)
	case int64:
		_b := any(b).(int64)
		return int(_a - _b)
	case uint:
		_b := any(b).(uint)
		return int(_a) - int(_b)
	case uint16:
		_b := any(b).(uint16)
		return int(_a) - int(_b)
	case uint32:
		_b := any(b).(uint32)
		return int(_a) - int(_b)
	case uint64:
		_b := any(b).(uint64)
		return int(_a) - int(_b)
	case uintptr:
		_b := any(b).(uintptr)
		return int(_a) - int(_b)
	case float32:
		_b := any(b).(float32)
		if _a < _b {
			return -1
		} else if _a == _b {
			return 0
		} else {
			return 1
		}
	case float64:
		_b := any(b).(float64)
		if _a < _b {
			return -1
		} else if _a == _b {
			return 0
		} else {
			return 1
		}
	case string:
		_b := any(b).(string)
		if _a < _b {
			return -1
		} else if _a == _b {
			return 0
		} else {
			return 1
		}
	default:
		return -1 // When used in sorting functions, return -1 means remaining the original order.
	}
}

// Orders the items of the given slice according to the specified comparable `key` (whose value must
// either be of type int or string). `order` can be either `asc` or `desc`.
//
// This function does not mutate the original slice but create a new one. However, if the `order` is
// malformed (not `asc` or `desc`), the function simply returns the original slice without ordering.
func OrderBy[S ~[]T, T ~map[K]V, K comparable, V any](original S, key K, order string) S {
	if order != "asc" && order != "desc" {
		return original
	}

	items := slices.Clone(original)
	slices.SortStableFunc(items, func(a, b T) int {
		_a, aOk := a[key]
		_b, bOk := b[key]

		if !aOk || !bOk {
			return -1 // remain the original order
		}

		return CompareItems(_a, _b)
	})

	if order == "desc" {
		slices.Reverse(items)
	}

	return items
}

// Groups the items of the given slice according to the comparable values returned
// by a provided callback function.
// The returned map has separate properties for each group, containing slices with the items in the
// group.
func GroupBy[S ~[]T, T ~map[K1]V, K1 comparable, V any, K2 comparable](
	items S,
	fn func(item T, idx int) K2,
) map[K2]S {
	groups := map[K2]S{}

	for idx, item := range items {
		key := fn(item, idx)
		list, ok := groups[key]

		if ok {
			groups[key] = append(list, item)
		} else {
			groups[key] = []T{item}
		}
	}

	return groups
}

// Finds all the items presented in the first slice but are missing in the others.
func Diff[S ~[]T, T comparable](first S, others ...S) S {
	others = Filter(others, func(item S, _ int) bool { return item != nil })
	flatted := Flat(others)
	items := S{}

	for _, item := range first {
		if !slices.Contains(flatted, item) {
			items = append(items, item)
		}
	}

	return items
}

// Creates a slice of unique values that is the symmetric difference of the given sources.
func Xor[S ~[]T, T comparable](sources ...S) S {
	items := S{}
	sources = Filter(sources, func(item S, _ int) bool { return item != nil })
	intersection := Intersect(sources...)

	for _, source := range sources {
		for _, item := range source {
			if !slices.Contains(intersection, item) {
				items = append(items, item)
			}
		}
	}

	return items
}

// Creates a new slice containing all the items from the given sources and remove the duplicated
// ones.
func Union[S ~[]T, T comparable](sources ...S) S {
	items := S{}

	for _, source := range sources {
		if source == nil {
			continue
		}

		for _, item := range source {
			if !slices.Contains(items, item) {
				items = append(items, item)
			}
		}
	}

	return items
}

// Creates a slice containing the items which are presented in all the sources, duplicated items are
// removed.
func Intersect[S ~[]T, T comparable](sources ...S) S {
	items := S{}
	union := Union(sources...)

	for _, item := range union {
		if Every(sources, func(source S, idx int) bool { return slices.Contains(source, item) }) {
			items = append(items, item)
		}
	}

	return items
}
