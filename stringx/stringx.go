// Additional functions for string processing that are missing in the standard library.
package stringx

import (
	"math"
	"math/rand"
	"regexp"
	"strings"

	"github.com/ayonli/goext/slicex"
)

var wordRegex = regexp.MustCompile(`\w+`)
var spaceRegex = regexp.MustCompile(`\s+`)
var spaceSepRegex = regexp.MustCompile(`\S\s+\S`)

// Returns a random string, the charset matches `/[0-9a-zA-Z]/`.
func Random(length int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	limit := float64(len(chars))
	bytes := []byte{}

	for length > 0 {
		i := int(math.Floor(rand.Float64() * limit))
		bytes = append(bytes, chars[i])
		length--
	}

	return string(bytes)
}

// Checks if the given string starts with the specified sub string.
//
// This function is the same as `strings.HasPrefix()`.
func StartsWith(str string, sub string) bool {
	return len(str) >= len(sub) && str[0:len(sub)] == sub
}

// Checks if the given string ends with the specified sub string.
//
// This function is the same as `strings.HasSuffix()`.
func EndsWith(str string, sub string) bool {
	return len(str) >= len(sub) && str[len(str)-len(sub):] == sub
}

// Pads the given string with another string (multiple times, if needed) until the resulting string
// reaches the final length. The padding is applied from the start of the string.
func PadStart(str string, finalLength int, padStr string) string {
	padLength := finalLength - len(str)

	if padLength <= 0 {
		return str
	}

	padChars := []byte(padStr)

	if len(padChars) > padLength {
		padChars = slicex.Slice(padChars, 0, padLength)
	}

	padCharsLimit := len(padChars) - 1
	startChars := make([]byte, padLength)

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
	padLength := finalLength - len(str)

	if padLength <= 0 {
		return str
	}

	padChars := []byte(padStr)

	if len(padChars) > padLength {
		padChars = slicex.Slice(padChars, 0, padLength)
	}

	padCharsLimit := len(padChars) - 1
	endChars := make([]byte, padLength)

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

// Capitalizes the given string, if `all` is true, all words are capitalized, otherwise only the
// first word will be capitalized.
func Capitalize(str string, all bool) string {
	if all {
		return wordRegex.ReplaceAllStringFunc(str, func(s string) string {
			return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
		})
	} else {
		loc := wordRegex.FindStringIndex(str)

		if loc == nil {
			return str
		}

		idx := loc[0]
		return str[0:idx] + strings.ToUpper(str[idx:idx+1]) + str[idx+1:]
	}
}

// Replaces the spaces between non-empty characters of the given string with hyphens (`-`).
func Hyphenate(str string) string {
	return spaceSepRegex.ReplaceAllStringFunc(str, func(s string) string {
		return spaceRegex.ReplaceAllString(s, "-")
	})
}

// Returns a section of the string selected from `start` to `end` (excluded).
//
// If `start < 0`, it will be calculated as `Length(str) + start`.
//
// If `end < 0`, it will be calculated as `Length(str) + end`.
func Slice(str string, start int, end int) string {
	limit := len(str)

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
		return "" // return an empty string directly
	}

	return str[start:end]
}

// Returns a section of the string selected from `start` to `end` (excluded).
//
// This function is similar to the `Slice()`, except it doesn't accept negative positions.
func Substring(str string, start int, end int) string {
	limit := len(str)

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

	return str[start:end]
}

// Extracts words (in latin characters) from the given string.
func Words(str string) []string {
	return wordRegex.FindAllString(str, -1)
}

// Breaks the string into smaller chunks according to the given length.
func Chunk(str string, length int) []string {
	limit := len(str)
	size := int(math.Ceil(float64(limit) / float64(length)))
	chunks := make([]string, size)
	offset := 0
	idx := 0

	for offset < limit {
		end := offset + length

		if end > limit {
			end = limit
		}

		chunks[idx] = str[offset:end]
		offset += length
		idx++
	}

	return chunks
}

// Truncates the given string to the given length (including the ending `...`).
func Truncate(str string, length int) string {
	limit := len(str)

	if length <= 0 {
		return ""
	} else if length >= limit {
		return str
	} else {
		length -= 3

		return str[0:length] + "..."
	}
}

// Executes a search for a match between a regular expression and the string, returning the index of
// the first match in the string.
func Search(str string, pattern string) int {
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return -1
	}

	loc := regex.FindStringIndex(str)

	if loc == nil {
		return -1
	}

	return loc[0]
}

// Retrieves the first result (with sub matches) of matching the string against a regular expression.
// If no match, this function returns nil.
func Match(str string, patten string) []string {
	regex, err := regexp.Compile(patten)

	if err != nil {
		return nil
	}

	return regex.FindStringSubmatch(str)
}

// Retrieves all results (with sub matches) of matching the string against a regular expression.
// If no match, this function returns an empty slice.
func MatchAll(str string, pattern string) [][]string {
	regexp, err := regexp.Compile(pattern)

	if err != nil {
		return [][]string(nil)
	}

	return regexp.FindAllStringSubmatch(str, -1)
}
