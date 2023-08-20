package mbstring_test

import (
	"testing"

	"github.com/ayonli/goext/stringx/mbstring"
	"github.com/stretchr/testify/assert"
)

func TestAt(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	char1 := mbstring.At(str1, 7)
	char2 := mbstring.At(str2, 3)
	char3 := mbstring.At(str3, 7)
	char4 := mbstring.At(str3, 10)

	assert.Equal(t, "W", char1)
	assert.Equal(t, "世", char2)
	assert.Equal(t, "世", char3)
	assert.Equal(t, "", char4)
}

func TestIndex(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	idx1 := mbstring.Index(str1, "World")
	idx2 := mbstring.Index(str2, "世界")
	idx3 := mbstring.Index(str3, "世界")

	assert.Equal(t, 7, idx1)
	assert.Equal(t, 3, idx2)
	assert.Equal(t, 7, idx3)
}

func TestLastIndex(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！Hi, 世界"
	str3 := "Hello, 世界！Hi, 世界"
	idx1 := mbstring.LastIndex(str1, "o")
	idx2 := mbstring.LastIndex(str2, "世界")
	idx3 := mbstring.LastIndex(str3, "H")

	assert.Equal(t, 8, idx1)
	assert.Equal(t, 10, idx2)
	assert.Equal(t, 10, idx3)
}

func TestLength(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	len1 := mbstring.Length(str1)
	len2 := mbstring.Length(str2)
	len3 := mbstring.Length(str3)

	assert.Equal(t, 13, len1)
	assert.Equal(t, 6, len2)
	assert.Equal(t, 10, len3)
}

func TestPadStart(t *testing.T) {
	str1 := "你好，世界！"
	str2 := mbstring.PadStart(str1, 8, " ")
	str3 := mbstring.PadStart(str1, 8, "*")
	str4 := mbstring.PadStart(str1, 8, "Hi")
	str5 := mbstring.PadStart(str1, 8, "Hola")

	assert.Equal(t, "  你好，世界！", str2)
	assert.Equal(t, "**你好，世界！", str3)
	assert.Equal(t, "Hi你好，世界！", str4)
	assert.Equal(t, "Ho你好，世界！", str5)
}

func TestPadEnd(t *testing.T) {
	str1 := "你好，世界！"
	str2 := mbstring.PadEnd(str1, 8, " ")
	str3 := mbstring.PadEnd(str1, 8, "*")
	str4 := mbstring.PadEnd(str1, 8, "Hi")
	str5 := mbstring.PadEnd(str1, 8, "Hola")

	assert.Equal(t, "你好，世界！  ", str2)
	assert.Equal(t, "你好，世界！**", str3)
	assert.Equal(t, "你好，世界！Hi", str4)
	assert.Equal(t, "你好，世界！Ho", str5)
}

func TestSlice(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	str4 := mbstring.Slice(str1, 0, 5)
	str5 := mbstring.Slice(str2, 3, -1)
	str6 := mbstring.Slice(str3, -3, -1)
	str7 := mbstring.Slice(str3, -3, -5)
	str8 := mbstring.Slice(str3, 5, 3)

	assert.Equal(t, "Hello", str4)
	assert.Equal(t, "世界", str5)
	assert.Equal(t, "世界", str6)
	assert.Equal(t, "", str7)
	assert.Equal(t, "", str8)
}

func TestSubstring(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "Hello, 世界！"
	str3 := mbstring.Substring(str1, 0, 5)
	str4 := mbstring.Substring(str2, 5, 3)

	assert.Equal(t, "Hello", str3)
	assert.Equal(t, "", str4)
}

func TestChunk(t *testing.T) {
	str1 := "foobar"
	str2 := "你好世界"
	chunks1 := mbstring.Chunk(str1, 2)
	chunks2 := mbstring.Chunk(str1, 4)
	chunks3 := mbstring.Chunk(str2, 2)
	chunks4 := mbstring.Chunk(str2, 3)

	assert.Equal(t, []string{"fo", "ob", "ar"}, chunks1)
	assert.Equal(t, []string{"foob", "ar"}, chunks2)
	assert.Equal(t, []string{"你好", "世界"}, chunks3)
	assert.Equal(t, []string{"你好世", "界"}, chunks4)
}

func TestTruncate(t *testing.T) {
	str1 := "Hello, World!"
	str2 := mbstring.Truncate(str1, 15)
	str3 := mbstring.Truncate(str1, 12)
	str4 := mbstring.Truncate(str1, 10)
	str5 := mbstring.Truncate(str1, -1)

	str6 := "你好，世界！Hallo 世界！"
	str7 := mbstring.Truncate(str6, 15)
	str8 := mbstring.Truncate(str6, 12)
	str9 := mbstring.Truncate(str6, 10)
	str10 := mbstring.Truncate(str6, -1)

	assert.Equal(t, "Hello, World!", str2)
	assert.Equal(t, "Hello, Wo...", str3)
	assert.Equal(t, "Hello, ...", str4)
	assert.Equal(t, "", str5)

	assert.Equal(t, "你好，世界！Hallo 世界！", str7)
	assert.Equal(t, "你好，世界！Hal...", str8)
	assert.Equal(t, "你好，世界！H...", str9)
	assert.Equal(t, "", str10)
}

func TestSearch(t *testing.T) {
	str := "你好，世界！"
	idx1 := mbstring.Search(str, "，\\S+")
	idx2 := mbstring.Search(str, "，\\s+")

	assert.Equal(t, 2, idx1)
	assert.Equal(t, -1, idx2)
}
