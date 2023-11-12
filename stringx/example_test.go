package stringx_test

import (
	"fmt"

	"github.com/ayonli/goext/stringx"
)

func ExampleRandom() {
	str := stringx.Random(4)
	matches := stringx.Match(str, "[0-9a-zA-Z]{4}")

	fmt.Println(matches[0] == str)
	// Output:
	// true
}

func ExampleStartsWith() {
	str := "Hello, World"

	fmt.Println(stringx.StartsWith(str, "Hello"))
	fmt.Println(stringx.StartsWith(str, "World"))
	// Output:
	// true
	// false
}

func ExampleEndsWith() {
	str := "Hello, World"

	fmt.Println(stringx.EndsWith(str, "World"))
	fmt.Println(stringx.EndsWith(str, "Hello"))
	// Output:
	// true
	// false
}

func ExamplePadStart() {
	str1 := "Hello, World!"
	str2 := stringx.PadStart(str1, 15, " ")
	str3 := stringx.PadStart(str1, 15, "*")
	str4 := stringx.PadStart(str1, 15, "Hi")
	str5 := stringx.PadStart(str1, 15, "Hola")
	str6 := stringx.PadStart(str1, 12, "**")
	str7 := stringx.PadStart(str1, 18, "Hola")

	fmt.Printf("%#v\n", str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(str7)
	// Output:
	// "  Hello, World!"
	// **Hello, World!
	// HiHello, World!
	// HoHello, World!
	// Hello, World!
	// HolaHHello, World!
}

func ExamplePadEnd() {
	str1 := "Hello, World!"
	str2 := stringx.PadEnd(str1, 15, " ")
	str3 := stringx.PadEnd(str1, 15, "*")
	str4 := stringx.PadEnd(str1, 15, "Hi")
	str5 := stringx.PadEnd(str1, 15, "Hola")
	str6 := stringx.PadEnd(str1, 12, "**")
	str7 := stringx.PadEnd(str1, 18, "Hola")

	fmt.Printf("%#v\n", str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(str7)
	// Output:
	// "Hello, World!  "
	// Hello, World!**
	// Hello, World!Hi
	// Hello, World!Ho
	// Hello, World!
	// Hello, World!HolaH
}

func ExampleCapitalize() {
	str1 := "hello, world"
	str2 := stringx.Capitalize(str1, false)
	str3 := stringx.Capitalize(str1, true)
	str4 := stringx.Capitalize(" hello world", false)
	str5 := stringx.Capitalize("  hello    world", true)
	str6 := stringx.Capitalize("你好，世界！", false) // this doesn't effect since it contains no latin characters

	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Printf("%#v\n", str4)
	fmt.Printf("%#v\n", str5)
	fmt.Println(str6)
	// Output:
	// Hello, world
	// Hello, World
	// " Hello world"
	// "  Hello    World"
	// 你好，世界！
}

func ExampleHyphenate() {
	str1 := stringx.Hyphenate("hello world")
	str2 := stringx.Hyphenate(" hello  world   ")

	fmt.Println(str1)
	fmt.Printf("%#v\v", str2)
	// Output:
	// hello-world
	// " hello-world   "
}

func ExampleSlice() {
	str1 := "Hello, World!"
	str2 := stringx.Slice(str1, 0, 2)
	str3 := stringx.Slice(str1, 0, -2)
	str4 := stringx.Slice(str1, -6, -1)
	str5 := stringx.Slice(str1, 0, 20)
	str6 := stringx.Slice(str1, 20, 25) // exceeding index return an empty string

	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Printf("%#v\n", str6)
	// Output:
	// He
	// Hello, Worl
	// World
	// Hello, World!
	// ""
}

func ExampleSubstring() {
	str1 := "Hello, World!"
	str2 := stringx.Substring(str1, 0, 2)
	str3 := stringx.Substring(str1, 5, 3)
	str4 := stringx.Substring(str1, 7, 20)
	str5 := stringx.Substring(str1, -1, 5) // negative index will be reset to 0
	str6 := stringx.Substring(str1, 7, -1)

	fmt.Println(str2)
	fmt.Printf("%#v\n", str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Printf("%#v\n", str6)
	// Output:
	// He
	// ""
	// World!
	// Hello
	// ""
}

func ExampleWords() {
	str := "Hello, World!"
	words := stringx.Words(str)

	fmt.Println(words)
	// Output:
	// [Hello World]
}

func ExampleChunk() {
	str := "foobar"
	chunks1 := stringx.Chunk(str, 2)
	chunks2 := stringx.Chunk(str, 4)

	fmt.Println(chunks1)
	fmt.Println(chunks2)
	// Output:
	// [fo ob ar]
	// [foob ar]
}

func ExampleTruncate() {
	str1 := "Hello, World!"
	str2 := stringx.Truncate(str1, 15)
	str3 := stringx.Truncate(str1, 12)
	str4 := stringx.Truncate(str1, 10)
	str5 := stringx.Truncate(str1, -1) // negative indexing isn't supported

	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Printf("%#v\n", str5)
	// Output:
	// Hello, World!
	// Hello, Wo...
	// Hello, ...
	// ""
}

func ExampleSearch() {
	str := "Hello, World!"
	idx1 := stringx.Search(str, "l{2,}")
	idx2 := stringx.Search(str, "o{2,}")
	idx3 := stringx.Search(str, "[a") // invalid regex pattern returns -1

	fmt.Println(idx1)
	fmt.Println(idx2)
	fmt.Println(idx3)
	// Output:
	// 2
	// -1
	// -1
}

func ExampleMatch() {
	str := "Hello, World!"
	match1 := stringx.Match(str, "l{2,}")
	match2 := stringx.Match(str, "o{2,}")
	match3 := stringx.Match(str, "[a") // invalid regex pattern returns an empty string

	fmt.Println(match1)
	fmt.Printf("%#v\n", match2)
	fmt.Printf("%#v\n", match3)
	// Output:
	// [ll]
	// []string(nil)
	// []string(nil)
}

func ExampleMatchAll() {
	str := "Hello, World!"
	match1 := stringx.MatchAll(str, "l{2,}")
	match2 := stringx.MatchAll(str, "o{2,}")
	match3 := stringx.MatchAll(str, "[a") // invalid regex pattern returns an empty slice

	fmt.Println(match1)
	fmt.Printf("%#v\n", match2)
	fmt.Printf("%#v\n", match3)
	// Output:
	// [[ll]]
	// [][]string(nil)
	// [][]string(nil)
}
