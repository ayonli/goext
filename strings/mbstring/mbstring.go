package mbstring

import (
	"math"
	"strings"
	"unicode/utf8"

	"github.com/ayonli/goext/slices"
)

// Returns the character from the string according to the given index.
//
// If `i < 0`, it returns the character counting from the end of the string.
//
// If the given index doesn't contain a value (boundary exceeded), an empty string and `false` will
// be returned.
func At(str string, i int) (string, bool) {
	return slices.At(strings.Split(str, ""), i)
}

// Returns the index at which a given sub string can be found in the string, or -1 if it is not
// present.
func Index(str string, sub string) int {
	chars := strings.Split(str, "")
	subChars := strings.Split(sub, "")
	length := len(subChars)

	for i := range chars {
		if strings.Join(chars[i:i+length], "") == sub {
			return i
		}
	}

	return -1
}

// Returns the last index at which a given sub string can be found in the string, or -1 if it is not
// present. The string is searched backwards.
func LastIndex(str string, sub string) int {
	chars := strings.Split(str, "")
	subChars := strings.Split(sub, "")
	length := len(subChars)
	limit := len(chars)

	for i := len(chars) - 1; i >= 0; i-- {
		end := i + length

		if end > limit {
			end = limit
		}

		if strings.Join(chars[i:end], "") == sub {
			return i
		}
	}

	return -1
}

// Returns the number of the characters in the string.
func Length(str string) int {
	return utf8.RuneCountInString(str)
}

// Returns a section of the string selected from `start` to `end` (excluded).
//
// If `start < 0`, it will be calculated as `Length(str) + start`.
//
// If `end < 0`, it will be calculated as `Length(str) + end`.
func Slice(str string, start int, end int) string {
	chars := strings.Split(str, "")
	return strings.Join(slices.Slice(chars, start, end), "")
}

// Breaks the string into smaller chunks according to the given length.
func Chunk(str string, length int) []string {
	chars := strings.Split(str, "")
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

		chunks[idx] = strings.Join(chars[offset:end], "")
		offset += length
		idx++
	}

	return chunks
}

// Truncates the given string to the given length (including the ending `...`).
func Truncate(str string, length int) string {
	chars := strings.Split(str, "")
	limit := len(chars)

	if length <= 0 {
		return ""
	} else if length >= limit {
		return str
	} else {
		length -= 3

		return strings.Join(chars[0:length], "") + "..."
	}
}
