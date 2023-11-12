// Additional functions for processing strings in multi-byte sequence.
package mbstring

import (
	"math"
	"slices"
	"unicode/utf8"

	"github.com/ayonli/goext/slicex"
	"github.com/ayonli/goext/stringx"
)

// Returns the character from the string according to the given index.
//
// If `i < 0`, it returns the character counting from the end of the string.
//
// If the given index doesn't contain a value (boundary exceeded), an empty string will be returned.
func At(str string, i int) string {
	if str == "" {
		return ""
	}

	code, ok := slicex.At([]rune(str), i)

	if ok {
		return string(code)
	} else {
		return ""
	}
}

// Returns the index at which a given sub string can be found in the string, or -1 if it is not
// present.
func Index(str string, sub string) int {
	if str == "" {
		if sub == "" {
			return 0
		} else {
			return -1
		}
	}

	chars := []rune(str)
	subChars := []rune(sub)
	length := len(subChars)
	limit := len(chars)

	for i := range chars {
		end := i + length

		if end > limit {
			break
		} else if slices.Equal(chars[i:i+length], subChars) {
			return i
		}
	}

	return -1
}

// Returns the last index at which a given sub string can be found in the string, or -1 if it is not
// present. The string is searched backwards.
func LastIndex(str string, sub string) int {
	if str == "" {
		if sub == "" {
			return 0
		} else {
			return -1
		}
	}

	chars := []rune(str)
	subChars := []rune(sub)
	length := len(subChars)
	limit := len(chars)

	for i := limit - length; i >= 0; i-- {
		end := i + length

		if slices.Equal(chars[i:end], subChars) {
			return i
		}
	}

	return -1
}

// Returns the number of the characters in the string.
func Length(str string) int {
	return utf8.RuneCountInString(str)
}

// Pads the given string with another string (multiple times, if needed) until the resulting string
// reaches the final length. The padding is applied from the start of the string.
func PadStart(str string, finalLength int, padStr string) string {
	padLength := finalLength - Length(str)

	if padLength <= 0 {
		return str
	}

	padChars := []rune(padStr)

	if len(padChars) > padLength {
		padChars = slicex.Slice(padChars, 0, padLength)
	}

	padCharsLimit := len(padChars) - 1
	startChars := make([]rune, padLength)

	for i, j := 0, 0; i < padLength; i++ {
		startChars[i] = padChars[j]

		if j < padCharsLimit {
			j++
		} else {
			j = 0
		}
	}

	return string(startChars) + str
}

// Pads the given string with another string (multiple times, if needed) until the resulting string
// reaches the final length. The padding is applied from the end of the string.
func PadEnd(str string, finalLength int, padStr string) string {
	padLength := finalLength - Length(str)

	if padLength <= 0 {
		return str
	}

	padChars := []rune(padStr)

	if len(padChars) > padLength {
		padChars = slicex.Slice(padChars, 0, padLength)
	}

	padCharsLimit := len(padChars) - 1
	endChars := make([]rune, padLength)

	for i, j := 0, 0; i < padLength; i++ {
		endChars[i] = padChars[j]

		if j < padCharsLimit {
			j++
		} else {
			j = 0
		}
	}

	return str + string(endChars)
}

// Returns a section of the string selected from `start` to `end` (excluded).
//
// If `start < 0`, it will be calculated as `Length(str) + start`.
//
// If `end < 0`, it will be calculated as `Length(str) + end`.
func Slice(str string, start int, end int) string {
	chars := []rune(str)
	return string(slicex.Slice(chars, start, end))
}

// Returns a section of the string selected from `start` to `end` (excluded).
//
// This function is similar to the `Slice()`, except it doesn't accept negative positions.
func Substring(str string, start int, end int) string {
	chars := []rune(str)
	limit := len(chars)

	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = 0
	}

	if end >= limit {
		end = limit
	}

	if start >= end || start >= limit {
		return "" // return an empty string directly
	}

	return string(chars[start:end])
}

// Breaks the string into smaller chunks according to the given length.
func Chunk(str string, length int) []string {
	chars := []rune(str)
	limit := len(chars)
	size := int(math.Ceil(float64(limit) / float64(length)))
	chunks := make([]string, size)
	offset := 0
	idx := 0

	for offset < limit {
		end := offset + length

		if end > limit {
			end = limit
		}

		chunks[idx] = string(chars[offset:end])
		offset += length
		idx++
	}

	return chunks
}

// Truncates the given string to the given length (including the ending `...`).
func Truncate(str string, length int) string {
	chars := []rune(str)
	limit := len(chars)

	if length <= 0 {
		return ""
	} else if length >= limit {
		return str
	} else {
		length -= 3

		return string(chars[0:length]) + "..."
	}
}

// Executes a search for a match between a regular expression and the string, returning the index of
// the first match in the string.
func Search(str string, pattern string) int {
	match := stringx.Match(str, pattern)

	if match != nil {
		return Index(str, match[0])
	}

	return -1
}
