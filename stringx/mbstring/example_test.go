package mbstring_test

import (
	"fmt"
	"strconv"

	"github.com/ayonli/goext/stringx/mbstring"
)

func ExampleAt() {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	char1 := mbstring.At(str1, 7)
	char2 := mbstring.At(str2, 3)
	char3 := mbstring.At(str3, 7)
	char4 := mbstring.At(str3, 10)
	char5 := mbstring.At("", 0) // an empty string always returns an empty string

	fmt.Println(char1)
	fmt.Println(char2)
	fmt.Println(char3)
	fmt.Println(strconv.Quote(char4))
	fmt.Println(strconv.Quote(char5))
	// Output:
	// W
	// 世
	// 世
	// ""
	// ""
}

func ExampleIndex() {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	idx1 := mbstring.Index(str1, "World")
	idx2 := mbstring.Index(str2, "世界")
	idx3 := mbstring.Index(str3, "世界")
	idx4 := mbstring.Index(str3, "你好")
	idx5 := mbstring.Index("", "")
	idx6 := mbstring.Index("", "你好")

	fmt.Println(idx1)
	fmt.Println(idx2)
	fmt.Println(idx3)
	fmt.Println(idx4)
	fmt.Println(idx5)
	fmt.Println(idx6)
	// Output:
	// 7
	// 3
	// 7
	// -1
	// 0
	// -1
}

func ExampleLastIndex() {
	str1 := "Hello, World!"
	str2 := "你好，世界！Hi, 世界"
	str3 := "Hello, 世界！Hi, 世界"
	idx1 := mbstring.LastIndex(str1, "o")
	idx2 := mbstring.LastIndex(str2, "世界")
	idx3 := mbstring.LastIndex(str3, "H")
	idx4 := mbstring.LastIndex(str3, "你好")
	idx5 := mbstring.LastIndex("", "")
	idx6 := mbstring.LastIndex("", "你好")

	fmt.Println(idx1)
	fmt.Println(idx2)
	fmt.Println(idx3)
	fmt.Println(idx4)
	fmt.Println(idx5)
	fmt.Println(idx6)
	// Output:
	// 8
	// 10
	// 10
	// -1
	// 0
	// -1
}

func ExampleLength() {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	len1 := mbstring.Length(str1)
	len2 := mbstring.Length(str2)
	len3 := mbstring.Length(str3)

	fmt.Println(len1)
	fmt.Println(len2)
	fmt.Println(len3)
	// Output:
	// 13
	// 6
	// 10
}

func ExamplePadStart() {
	str1 := "你好，世界！"
	str2 := mbstring.PadStart(str1, 8, " ")
	str3 := mbstring.PadStart(str1, 8, "*")
	str4 := mbstring.PadStart(str1, 8, "Hi")
	str5 := mbstring.PadStart(str1, 8, "Hola")
	str6 := mbstring.PadStart(str1, 5, "**")

	fmt.Println(strconv.Quote(str2))
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	// Output:
	// "  你好，世界！"
	// **你好，世界！
	// Hi你好，世界！
	// Ho你好，世界！
	// 你好，世界！
}

func ExamplePadEnd() {
	str1 := "你好，世界！"
	str2 := mbstring.PadEnd(str1, 8, " ")
	str3 := mbstring.PadEnd(str1, 8, "*")
	str4 := mbstring.PadEnd(str1, 8, "Hi")
	str5 := mbstring.PadEnd(str1, 8, "Hola")
	str6 := mbstring.PadEnd(str1, 5, "**")

	fmt.Println(strconv.Quote(str2))
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	// Output:
	// "你好，世界！  "
	// 你好，世界！**
	// 你好，世界！Hi
	// 你好，世界！Ho
	// 你好，世界！
}

func ExampleSlice() {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := "Hello, 世界！"
	str4 := mbstring.Slice(str1, 0, 5)
	str5 := mbstring.Slice(str2, 3, -1)
	str6 := mbstring.Slice(str3, -3, -1)
	str7 := mbstring.Slice(str3, -3, -5)
	str8 := mbstring.Slice(str3, 5, 3)

	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(strconv.Quote(str7))
	fmt.Println(strconv.Quote(str8))
	// Output:
	// Hello
	// 世界
	// 世界
	// ""
	// ""
}

func ExampleSubstring() {
	str1 := "Hello, World!"
	str2 := "你好，世界！"
	str3 := mbstring.Substring(str1, 0, 5)
	str4 := mbstring.Substring(str2, 5, 3)
	str5 := mbstring.Substring(str2, 3, 20)
	str6 := mbstring.Substring(str2, -1, 5) // negative index will be reset to 0
	str7 := mbstring.Substring(str2, 7, -1)

	fmt.Println(str3)
	fmt.Println(strconv.Quote(str4))
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(strconv.Quote(str7))
	// Output:
	// Hello
	// ""
	// 世界！
	// 你好，世界
	// ""
}

func ExampleChunk() {
	str1 := "foobar"
	str2 := "你好世界"
	chunks1 := mbstring.Chunk(str1, 2)
	chunks2 := mbstring.Chunk(str1, 4)
	chunks3 := mbstring.Chunk(str2, 2)
	chunks4 := mbstring.Chunk(str2, 3)

	fmt.Println(chunks1)
	fmt.Println(chunks2)
	fmt.Println(chunks3)
	fmt.Println(chunks4)
	// Output:
	// [fo ob ar]
	// [foob ar]
	// [你好 世界]
	// [你好世 界]
}

func ExampleTruncate() {
	str1 := "Hello, World!"
	str2 := mbstring.Truncate(str1, 15)
	str3 := mbstring.Truncate(str1, 12)
	str4 := mbstring.Truncate(str1, 10)
	str5 := mbstring.Truncate(str1, -1) // 0 or negative length return an empty string

	str6 := "你好，世界！Hallo 世界！"
	str7 := mbstring.Truncate(str6, 15)
	str8 := mbstring.Truncate(str6, 12)
	str9 := mbstring.Truncate(str6, 10)
	str10 := mbstring.Truncate(str6, -1)

	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(strconv.Quote(str5))
	fmt.Println(str7)
	fmt.Println(str8)
	fmt.Println(str9)
	fmt.Println(strconv.Quote(str10))
	// Output:
	// Hello, World!
	// Hello, Wo...
	// Hello, ...
	// ""
	// 你好，世界！Hallo 世界！
	// 你好，世界！Hal...
	// 你好，世界！H...
	// ""
}

func ExampleSearch() {
	str := "你好，世界！"
	idx1 := mbstring.Search(str, "，\\S+")
	idx2 := mbstring.Search(str, "，\\s+")

	fmt.Println(idx1)
	fmt.Println(idx2)
	// Output:
	// 2
	// -1
}
