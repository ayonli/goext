package oop

import (
	"strings"

	"github.com/ayonli/goext/slicex"
	"github.com/ayonli/goext/stringx"
	"github.com/ayonli/goext/stringx/mbstring"
)

// String is an object-oriented abstract that works around multi-byte strings.
type String string

func (self *String) At(i int) *String {
	char := String(mbstring.At(string(*self), i))
	return &char
}

func (self *String) Index(sub string) int {
	return mbstring.Index(string(*self), sub)
}

func (self *String) LastIndex(sub string) int {
	return mbstring.LastIndex(string(*self), sub)
}

func (self *String) Length() int {
	return mbstring.Length(string(*self))
}

func (self *String) String() string {
	return string(*self)
}

func (self *String) Clone() *String {
	str := String(strings.Clone(string(*self)))
	return &str
}

func (self *String) Compare(another String) int {
	return strings.Compare(string(*self), string(another))
}

func (self *String) Contains(sub string) bool {
	return strings.Contains(string(*self), sub)
}

func (self *String) StartsWith(sub string) bool {
	return stringx.StartsWith(string(*self), sub)
}

func (self *String) EndsWith(sub string) bool {
	return stringx.EndsWith(string(*self), sub)
}

func (self *String) PadStart(finalLength int, padStr string) *String {
	str := String(mbstring.PadStart(string(*self), finalLength, padStr))
	return &str
}

func (self *String) PadEnd(finalLength int, padStr string) *String {
	str := String(mbstring.PadEnd(string(*self), finalLength, padStr))
	return &str
}

func (self *String) Trim(chars string) *String {
	str := String(strings.Trim(string(*self), chars))
	return &str
}

func (self *String) TrimLeft(chars string) *String {
	str := String(strings.TrimLeft(string(*self), chars))
	return &str
}

func (self *String) TrimRight(chars string) *String {
	str := String(strings.TrimRight(string(*self), chars))
	return &str
}

func (self *String) ToUpper() *String {
	str := String(strings.ToUpper(string(*self)))
	return &str
}

func (self *String) ToLower() *String {
	str := String(strings.ToLower(string(*self)))
	return &str
}

func (self *String) Capitalize(all bool) *String {
	str := String(stringx.Capitalize(string(*self), all))
	return &str
}

func (self *String) Hyphenate() *String {
	str := String(stringx.Hyphenate(string(*self)))
	return &str
}

func (self *String) Slice(start int, end int) *String {
	str := String(mbstring.Slice(string(*self), start, end))
	return &str
}

func (self *String) Substring(start int, end int) *String {
	str := String(mbstring.Substring(string(*self), start, end))
	return &str
}

func (self *String) Split(sep string) []String {
	return slicex.Map(strings.Split(string(*self), sep), func(str string, _ int) String {
		return String(str)
	})
}

func (self *String) Chunk(length int) []String {
	return slicex.Map(mbstring.Chunk(string(*self), length), func(str string, _ int) String {
		return String(str)
	})
}

func (self *String) Truncate(length int) *String {
	str := String(mbstring.Truncate(string(*self), length))
	return &str
}

func (self *String) Repeat(count int) *String {
	str := String(strings.Repeat(string(*self), count))
	return &str
}

func (self *String) Replace(old string, rep string) *String {
	str := String(strings.Replace(string(*self), old, rep, 1))
	return &str
}

func (self *String) ReplaceAll(old string, rep string) *String {
	str := String(strings.ReplaceAll(string(*self), old, rep))
	return &str
}

func (self *String) Search(pattern string) int {
	return mbstring.Search(string(*self), pattern)
}

func (self *String) Match(pattern string) *String {
	str := String(stringx.Match(string(*self), pattern))
	return &str
}

func (self *String) MatchAll(pattern string) []String {
	return slicex.Map(stringx.MatchAll(string(*self), pattern), func(str string, _ int) String {
		return String(str)
	})
}
