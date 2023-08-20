package stringx_test

import (
	"testing"

	"github.com/ayonli/goext/stringx"
	"github.com/stretchr/testify/assert"
)

func TestStartsWith(t *testing.T) {
	str := "Hello, Wolrd"

	assert.Equal(t, true, stringx.StartsWith(str, "Hello"))
	assert.Equal(t, false, stringx.StartsWith(str, "Wolrd"))
}

func TestEndsWith(t *testing.T) {
	str := "Hello, Wolrd"

	assert.Equal(t, true, stringx.EndsWith(str, "Wolrd"))
	assert.Equal(t, false, stringx.EndsWith(str, "Hello"))
}

func TestPadStart(t *testing.T) {
	str1 := "Hello, World!"
	str2 := stringx.PadStart(str1, 15, " ")
	str3 := stringx.PadStart(str1, 15, "*")
	str4 := stringx.PadStart(str1, 15, "Hi")
	str5 := stringx.PadStart(str1, 15, "Hola")

	assert.Equal(t, "  Hello, World!", str2)
	assert.Equal(t, "**Hello, World!", str3)
	assert.Equal(t, "HiHello, World!", str4)
	assert.Equal(t, "HoHello, World!", str5)
}

func TestPadEnd(t *testing.T) {
	str1 := "Hello, World!"
	str2 := stringx.PadEnd(str1, 15, " ")
	str3 := stringx.PadEnd(str1, 15, "*")
	str4 := stringx.PadEnd(str1, 15, "Hi")
	str5 := stringx.PadEnd(str1, 15, "Hola")

	assert.Equal(t, "Hello, World!  ", str2)
	assert.Equal(t, "Hello, World!**", str3)
	assert.Equal(t, "Hello, World!Hi", str4)
	assert.Equal(t, "Hello, World!Ho", str5)
}

func TestCapitalize(t *testing.T) {
	str1 := "hello, world"
	str2 := stringx.Capitalize(str1, false)
	str3 := stringx.Capitalize(str1, true)
	str4 := stringx.Capitalize(" hello world", false)
	str5 := stringx.Capitalize("  hello    world", true)

	assert.Equal(t, "Hello, world", str2)
	assert.Equal(t, "Hello, World", str3)
	assert.Equal(t, " Hello world", str4)
	assert.Equal(t, "  Hello    World", str5)
}

func TestHyphenate(t *testing.T) {
	str1 := stringx.Hyphenate("hello world")
	str2 := stringx.Hyphenate(" hello  world   ")

	assert.Equal(t, "hello-world", str1)
	assert.Equal(t, " hello-world   ", str2)
}

func TestSlice(t *testing.T) {
	str1 := "Hello, World!"
	str2 := stringx.Slice(str1, 0, 2)
	str3 := stringx.Slice(str1, 0, -2)

	assert.Equal(t, "He", str2)
	assert.Equal(t, "Hello, Worl", str3)
}

func TestS(t *testing.T) {
	str1 := "Hello, World!"
	str2 := stringx.Substring(str1, 0, 2)
	str3 := stringx.Substring(str1, 5, 3)

	assert.Equal(t, "He", str2)
	assert.Equal(t, "", str3)
}

func TestWords(t *testing.T) {
	str := "Hello, World!"
	words := stringx.Words(str)

	assert.Equal(t, []string{"Hello", "World"}, words)
}

func TestChunk(t *testing.T) {
	str := "foobar"
	chunks1 := stringx.Chunk(str, 2)
	chunks2 := stringx.Chunk(str, 4)

	assert.Equal(t, []string{"fo", "ob", "ar"}, chunks1)
	assert.Equal(t, []string{"foob", "ar"}, chunks2)
}

func TestTruncate(t *testing.T) {
	str1 := "Hello, World!"
	str2 := stringx.Truncate(str1, 15)
	str3 := stringx.Truncate(str1, 12)
	str4 := stringx.Truncate(str1, 10)
	str5 := stringx.Truncate(str1, -1)

	assert.Equal(t, "Hello, World!", str2)
	assert.Equal(t, "Hello, Wo...", str3)
	assert.Equal(t, "Hello, ...", str4)
	assert.Equal(t, "", str5)
}

func TestSearch(t *testing.T) {
	str := "Hello, World!"
	idx1 := stringx.Search(str, "l{2,}")
	idx2 := stringx.Search(str, "o{2,}")

	assert.Equal(t, 2, idx1)
	assert.Equal(t, -1, idx2)
}

func TestMatch(t *testing.T) {
	str := "Hello, World!"
	match1 := stringx.Match(str, "l{2,}")
	match2 := stringx.Match(str, "o{2,}")

	assert.Equal(t, "ll", match1)
	assert.Equal(t, "", match2)
}

func TestMatchAll(t *testing.T) {
	str := "Hello, World!"
	match1 := stringx.MatchAll(str, "l{2,}")
	match2 := stringx.MatchAll(str, "o{2,}")

	assert.Equal(t, []string{"ll"}, match1)
	assert.Equal(t, []string(nil), match2)
}
