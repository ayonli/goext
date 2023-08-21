package oop

import (
	"strings"

	"github.com/ayonli/goext/slicex"
	"github.com/ayonli/goext/stringx"
	"github.com/ayonli/goext/stringx/mbstring"
)

// String is an object-oriented abstract that works around multi-byte strings.
type String string

func (str String) At(i int) String {
	return String(mbstring.At(string(str), i))
}

func (str String) IndexOf(sub string) int {
	return mbstring.Index(string(str), sub)
}

func (str String) LastIndexOf(sub string) int {
	return mbstring.LastIndex(string(str), sub)
}

func (str String) Length() int {
	return mbstring.Length(string(str))
}

func (str String) String() string {
	return string(str)
}

func (str String) Clone() String {
	return String(strings.Clone(string(str)))
}

func (str String) Compare(another string) int {
	return strings.Compare(string(str), another)
}

func (str String) Contains(sub string) bool {
	return strings.Contains(string(str), sub)
}

func (str String) Count(sub string) int {
	return strings.Count(string(str), sub)
}

func (str String) StartsWith(sub string) bool {
	return stringx.StartsWith(string(str), sub)
}

func (str String) EndsWith(sub string) bool {
	return stringx.EndsWith(string(str), sub)
}

func (str String) PadStart(finalLength int, padStr string) String {
	return String(mbstring.PadStart(string(str), finalLength, padStr))
}

func (str String) PadEnd(finalLength int, padStr string) String {
	return String(mbstring.PadEnd(string(str), finalLength, padStr))
}

func (str String) Trim(chars string) String {
	return String(strings.Trim(string(str), chars))
}

func (str String) TrimLeft(chars string) String {
	return String(strings.TrimLeft(string(str), chars))
}

func (str String) TrimRight(chars string) String {
	return String(strings.TrimRight(string(str), chars))
}

func (str String) ToUpperCase() String {
	return String(strings.ToUpper(string(str)))
}

func (str String) ToLowerCase() String {
	return String(strings.ToLower(string(str)))
}

func (str String) Capitalize(all bool) String {
	return String(stringx.Capitalize(string(str), all))
}

func (str String) Hyphenate() String {
	return String(stringx.Hyphenate(string(str)))
}

func (str String) Slice(start int, end int) String {
	return String(mbstring.Slice(string(str), start, end))
}

func (str String) Substring(start int, end int) String {
	return String(mbstring.Substring(string(str), start, end))
}

func (str String) Split(sep string) []String {
	return slicex.Map(strings.Split(string(str), sep), func(str string, _ int) String {
		return String(str)
	})
}

func (str String) Chunk(length int) []String {
	return slicex.Map(mbstring.Chunk(string(str), length), func(str string, _ int) String {
		return String(str)
	})
}

func (str String) Truncate(length int) String {
	return String(mbstring.Truncate(string(str), length))
}

func (str String) Repeat(count int) String {
	return String(strings.Repeat(string(str), count))
}

func (str String) Replace(old string, rep string) String {
	return String(strings.Replace(string(str), old, rep, 1))
}

func (str String) ReplaceAll(old string, rep string) String {
	return String(strings.ReplaceAll(string(str), old, rep))
}

func (str String) Search(pattern string) int {
	return mbstring.Search(string(str), pattern)
}

func (str String) Match(pattern string) String {
	return String(stringx.Match(string(str), pattern))
}

func (str String) MatchAll(pattern string) []String {
	return slicex.Map(stringx.MatchAll(string(str), pattern), func(str string, _ int) String {
		return String(str)
	})
}
