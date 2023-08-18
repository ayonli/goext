package strings

import (
	"math"
	"regexp"
	gStrings "strings"
)

var wordRegex = regexp.MustCompile("\\w+")
var spaceRegex = regexp.MustCompile("\\s+")
var spaceSepRegex = regexp.MustCompile("\\S\\s+\\S")

// Checks if the given string starts with the specified sub string.
func StartsWith(str string, sub string) bool {
	return str[0:len(sub)] == sub
}

// Checks if the given string ends with the specified sub string.
func EndsWith(str string, sub string) bool {
	return str[len(str)-len(sub):] == sub
}

// Pads the given string with another string (multiple times, if needed) until the resulting string
// reaches the final length. The padding is applied from the start of the string.
func PadStart(str string, finalLength int, padStr string) string {
	leftLength := finalLength - len(str)

	if leftLength <= 0 {
		return str
	}

	if len(padStr) > leftLength {
		padStr = padStr[0:leftLength]
	}

	for len(str) < finalLength {
		str = padStr + str
	}

	return str
}

// Pads the given string with another string (multiple times, if needed) until the resulting string
// reaches the final length. The padding is applied from the end of the string.
func PadEnd(str string, finalLength int, padStr string) string {
	leftLength := finalLength - len(str)

	if leftLength <= 0 {
		return str
	}

	if len(padStr) > leftLength {
		padStr = padStr[0:leftLength]
	}

	for len(str) < finalLength {
		str += padStr
	}

	return str
}

// Capitalizes the given string, if `all` is true, all words are capitalized, otherwise only the
// first word will be capitalized.
func Capitalize(str string, all bool) string {
	if all {
		return wordRegex.ReplaceAllStringFunc(str, func(s string) string {
			return gStrings.ToUpper(s[:1]) + gStrings.ToLower(s[1:])
		})
	} else {
		loc := wordRegex.FindStringIndex(str)

		if loc == nil {
			return str
		}

		idx := loc[0]
		return str[0:idx] + gStrings.ToUpper(str[idx:idx+1]) + str[idx+1:]
	}
}

// Replaces the spaces of the given string with hyphens (`-`).
func Hyphenate(str string) string {
	return spaceSepRegex.ReplaceAllStringFunc(str, func(s string) string {
		return spaceRegex.ReplaceAllString(s, "-")
	})
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
