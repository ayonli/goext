package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartsWith(t *testing.T) {
	str := "Hello, Wolrd"

	assert.Equal(t, true, StartsWith(str, "Hello"))
	assert.Equal(t, false, StartsWith(str, "Wolrd"))
}

func TestEndsWith(t *testing.T) {
	str := "Hello, Wolrd"

	assert.Equal(t, true, EndsWith(str, "Wolrd"))
	assert.Equal(t, false, EndsWith(str, "Hello"))
}

func TestPadStart(t *testing.T) {
	str1 := "Hello, World!"
	str2 := PadStart(str1, 15, " ")
	str3 := PadStart(str1, 15, "*")
	str4 := PadStart(str1, 15, "Hi")
	str5 := PadStart(str1, 15, "Hola")

	assert.Equal(t, "  Hello, World!", str2)
	assert.Equal(t, "**Hello, World!", str3)
	assert.Equal(t, "HiHello, World!", str4)
	assert.Equal(t, "HoHello, World!", str5)
}

func TestPadEnd(t *testing.T) {
	str1 := "Hello, World!"
	str2 := PadEnd(str1, 15, " ")
	str3 := PadEnd(str1, 15, "*")
	str4 := PadEnd(str1, 15, "Hi")
	str5 := PadEnd(str1, 15, "Hola")

	assert.Equal(t, "Hello, World!  ", str2)
	assert.Equal(t, "Hello, World!**", str3)
	assert.Equal(t, "Hello, World!Hi", str4)
	assert.Equal(t, "Hello, World!Ho", str5)
}

func TestCapitalize(t *testing.T) {
	str1 := "hello, world"
	str2 := Capitalize(str1, false)
	str3 := Capitalize(str1, true)
	str4 := Capitalize(" hello world", false)
	str5 := Capitalize("  hello    world", true)

	assert.Equal(t, "Hello, world", str2)
	assert.Equal(t, "Hello, World", str3)
	assert.Equal(t, " Hello world", str4)
	assert.Equal(t, "  Hello    World", str5)
}

func TestHyphenate(t *testing.T) {
	str1 := Hyphenate("hello world")
	str2 := Hyphenate(" hello  world   ")

	assert.Equal(t, "hello-world", str1)
	assert.Equal(t, " hello-world   ", str2)
}

func TestWords(t *testing.T) {
	str := "Hello, World!"
	words := Words(str)

	assert.Equal(t, []string{"Hello", "World"}, words)
}

func TestChunk(t *testing.T) {
	str := "foobar"
	chunks1 := Chunk(str, 2)
	chunks2 := Chunk(str, 4)

	assert.Equal(t, []string{"fo", "ob", "ar"}, chunks1)
	assert.Equal(t, []string{"foob", "ar"}, chunks2)
}

func TestTruncate(t *testing.T) {
	str1 := "Hello, World!"
	str2 := Truncate(str1, 15)
	str3 := Truncate(str1, 12)
	str4 := Truncate(str1, 10)
	str5 := Truncate(str1, -1)

	assert.Equal(t, "Hello, World!", str2)
	assert.Equal(t, "Hello, Wo...", str3)
	assert.Equal(t, "Hello, ...", str4)
	assert.Equal(t, "", str5)
}

func TestSearch(t *testing.T) {
	str := "Hello, World!"
	idx1 := Search(str, "l{2,}")
	idx2 := Search(str, "o{2,}")

	assert.Equal(t, 2, idx1)
	assert.Equal(t, -1, idx2)
}

func TestMatch(t *testing.T) {
	str := "Hello, World!"
	match1 := Match(str, "l{2,}")
	match2 := Match(str, "o{2,}")

	assert.Equal(t, "ll", match1)
	assert.Equal(t, "", match2)
}

func TestMatchAll(t *testing.T) {
	str := "Hello, World!"
	match1 := MatchAll(str, "l{2,}")
	match2 := MatchAll(str, "o{2,}")

	assert.Equal(t, []string{"ll"}, match1)
	assert.Equal(t, []string(nil), match2)
}
