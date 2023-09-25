// Functions for dealing with numbers.
package number

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// Returns a random integer ranged from `min` to `max` (inclusive).
func Random(min int, max int) int {
	return min + int(math.Floor(rand.Float64()*float64(max-min+1)))
}

// Creates a generator that produces sequential numbers from `min` to `max` (inclusive).
func Sequence(min int, max int, step int, loop bool) <-chan int {
	channel := make(chan int)

	go func() {
		id := min

		for {
			channel <- id
			id += step

			if id > max {
				if loop {
					id = min
				} else {
					break
				}
			}
		}

		close(channel)
	}()

	return channel
}

// Returns `true` if the given value is a numeric value, `false` otherwise. A numeric value is an
// int or float family value, or a string that can be converted to a number.
func IsNumeric(value any) bool {
	switch val := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	case string:
		_, ok := Parse(val)
		return ok
	default:
		return false
	}
}

// Parses the given numeric string to a number.
func Parse(numeric string) (float64, bool) {
	numeric = strings.TrimSpace(numeric)
	num, err := strconv.ParseFloat(numeric, 64)

	if err == nil {
		return num, true
	}

	_num, _err := strconv.ParseInt(numeric, 0, 64)

	if _err == nil {
		return float64(_num), true
	} else {
		return 0, false
	}
}
