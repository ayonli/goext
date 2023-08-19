package mbstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAt(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	char1 := At(str1, 7)
	char2 := At(str2, 3)
	char3 := At(str3, 7)
	char4 := At(str3, 10)

	assert.Equal(t, "W", char1)
	assert.Equal(t, "世", char2)
	assert.Equal(t, "世", char3)
	assert.Equal(t, "", char4)
}

func TestIndex(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	idx1 := Index(str1, "World")
	idx2 := Index(str2, "世界")
	idx3 := Index(str3, "世界")

	assert.Equal(t, 7, idx1)
	assert.Equal(t, 3, idx2)
	assert.Equal(t, 7, idx3)
}

func TestLastIndex(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！Hi, 世界"
	str3 := "Hello, 世界！Hi, 世界"
	idx1 := LastIndex(str1, "o")
	idx2 := LastIndex(str2, "世界")
	idx3 := LastIndex(str3, "H")

	assert.Equal(t, 8, idx1)
	assert.Equal(t, 10, idx2)
	assert.Equal(t, 10, idx3)
}

func TestLength(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	len1 := Length(str1)
	len2 := Length(str2)
	len3 := Length(str3)

	assert.Equal(t, 13, len1)
	assert.Equal(t, 6, len2)
	assert.Equal(t, 10, len3)
}

func TestPadStart(t *testing.T) {
	str1 := "你好，世界！"
	str2 := PadStart(str1, 8, " ")
	str3 := PadStart(str1, 8, "*")
	str4 := PadStart(str1, 8, "Hi")
	str5 := PadStart(str1, 8, "Hola")

	assert.Equal(t, "  你好，世界！", str2)
	assert.Equal(t, "**你好，世界！", str3)
	assert.Equal(t, "Hi你好，世界！", str4)
	assert.Equal(t, "Ho你好，世界！", str5)
}

func TestPadEnd(t *testing.T) {
	str1 := "你好，世界！"
	str2 := PadEnd(str1, 8, " ")
	str3 := PadEnd(str1, 8, "*")
	str4 := PadEnd(str1, 8, "Hi")
	str5 := PadEnd(str1, 8, "Hola")

	assert.Equal(t, "你好，世界！  ", str2)
	assert.Equal(t, "你好，世界！**", str3)
	assert.Equal(t, "你好，世界！Hi", str4)
	assert.Equal(t, "你好，世界！Ho", str5)
}

func TestSlice(t *testing.T) {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	str4 := Slice(str1, 0, 5)
	str5 := Slice(str2, 3, -1)
	str6 := Slice(str3, -3, -1)
	str7 := Slice(str3, -3, -5)
	str8 := Slice(str3, 5, 3)

	assert.Equal(t, "Hello", str4)
	assert.Equal(t, "世界", str5)
	assert.Equal(t, "世界", str6)
	assert.Equal(t, "", str7)
	assert.Equal(t, "", str8)
}

func TestChunk(t *testing.T) {
	str1 := "foobar"
	str2 := "你好世界"
	chunks1 := Chunk(str1, 2)
	chunks2 := Chunk(str1, 4)
	chunks3 := Chunk(str2, 2)
	chunks4 := Chunk(str2, 3)

	assert.Equal(t, []string{"fo", "ob", "ar"}, chunks1)
	assert.Equal(t, []string{"foob", "ar"}, chunks2)
	assert.Equal(t, []string{"你好", "世界"}, chunks3)
	assert.Equal(t, []string{"你好世", "界"}, chunks4)
}

func TestTruncate(t *testing.T) {
	str1 := "Hello, World!"
	str2 := Truncate(str1, 15)
	str3 := Truncate(str1, 12)
	str4 := Truncate(str1, 10)
	str5 := Truncate(str1, -1)

	str6 := "你好，世界！Hallo 世界！"
	str7 := Truncate(str6, 15)
	str8 := Truncate(str6, 12)
	str9 := Truncate(str6, 10)
	str10 := Truncate(str6, -1)

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
	idx1 := Search(str, "，\\S+")
	idx2 := Search(str, "，\\s+")

	assert.Equal(t, 2, idx1)
	assert.Equal(t, -1, idx2)
}
