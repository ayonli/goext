package oop_test

import (
	"fmt"
	"strconv"

	"github.com/ayonli/goext/oop"
)

func ExampleString() {
	str := oop.String("你好，世界！")

	fmt.Println(str)
	// Output:
	// 你好，世界！
}

func ExampleString_At() {
	str := oop.String("你好，世界！")

	fmt.Println(str.At(0))  // do not use str[0] as it refers to a single byte
	fmt.Println(str.At(1))  // do not use str[1] as it refers to a single byte
	fmt.Println(str.At(-1)) // negative index counts backwards
	// Output:
	// 你
	// 好
	// ！
}

func ExampleString_IndexOf() {
	str := oop.String("你好，世界！")

	fmt.Println(str.IndexOf("世界"))
	fmt.Println(str.IndexOf("哈啰"))
	// Output:
	// 3
	// -1
}

func ExampleString_LastIndexOf() {
	str := oop.String("你好，世界！嗨，世界！")

	fmt.Println(str.LastIndexOf("世界"))
	fmt.Println(str.LastIndexOf("哈啰"))
	// Output:
	// 8
	// -1
}

func ExampleString_Length() {
	str := oop.String("你好，世界！")

	fmt.Println(str.Length())
	// Output:
	// 6
}

func ExampleString_Clone() {
	str1 := oop.String("你好，世界！")
	str2 := str1.Clone()

	fmt.Println(str2)
	// Output:
	// 你好，世界！
}

func ExampleString_Compare() {
	str1 := oop.String("你好，世界！")
	str2 := oop.String("你好，世界！")
	str3 := oop.String("Hello，世界！")
	str4 := oop.String("你好，祖国！")

	fmt.Println(str1.Compare(str2))
	fmt.Println(str1.Compare(str3))
	fmt.Println(str1.Compare(str4))
	// Output:
	// 0
	// 1
	// -1
}

func ExampleString_Contains() {
	str := oop.String("你好，世界！")

	fmt.Println(str.Contains("你好"))
	fmt.Println(str.Contains("Hello"))
	// Output:
	// true
	// false
}

func ExampleString_StartsWith() {
	str := oop.String("你好，世界！")

	fmt.Println(str.StartsWith("你好"))
	fmt.Println(str.StartsWith("Hello"))
	// Output:
	// true
	// false
}

func ExampleString_EndsWith() {
	str := oop.String("你好，世界！")

	fmt.Println(str.EndsWith("世界！"))
	fmt.Println(str.EndsWith("World!"))
	// Output:
	// true
	// false
}

func ExampleString_PadStart() {
	str := oop.String("你好，世界！")

	fmt.Println(str.PadStart(10, "*"))
	// Output:
	// ****你好，世界！
}

func ExampleString_PadEnd() {
	str := oop.String("你好，世界！")

	fmt.Println(str.PadEnd(10, "*"))
	// Output:
	// 你好，世界！****
}

func ExampleString_Trim() {
	str := oop.String("  你好，世界！  ")

	fmt.Println(strconv.Quote(string(*str.Trim(" "))))
	// Output:
	// "你好，世界！"
}

func ExampleString_TrimLeft() {
	str := oop.String("  你好，世界！  ")

	fmt.Println(strconv.Quote(string(*str.TrimLeft(" "))))
	// Output:
	// "你好，世界！  "
}

func ExampleString_TrimRight() {
	str := oop.String("  你好，世界！  ")

	fmt.Println(strconv.Quote(string(*str.TrimRight(" "))))
	// Output:
	// "  你好，世界！"
}

func ExampleString_ToUpperCase() {
	str := oop.String("Hello, World!")

	fmt.Println(str.ToUpperCase())
	// Output:
	// HELLO, WORLD!
}

func ExampleString_ToLowerCase() {
	str := oop.String("Hello, World!")

	fmt.Println(str.ToLowerCase())
	// Output:
	// hello, world!
}

func ExampleString_Capitalize() {
	str := oop.String("hello, world!")

	fmt.Println(str.Capitalize(false))
	fmt.Println(str.Capitalize(true))
	// Output:
	// Hello, world!
	// Hello, World!
}

func ExampleString_Hyphenate() {
	str := oop.String("hello world")

	fmt.Println(str.Hyphenate())
	// Output:
	// hello-world
}

func ExampleString_Slice() {
	str := oop.String("你好，世界！")

	fmt.Println(str.Slice(0, 2))   // don't use str[0:2] as it refers to only 2 bytes
	fmt.Println(str.Slice(-3, -1)) // negative index counts backwards
	// Output:
	// 你好
	// 世界
}

func ExampleString_Substring() {
	str := oop.String("你好，世界！")

	fmt.Println(str.Substring(0, 2))   // don't use str[0:2] as it refers to only 2 bytes
	fmt.Println(str.Substring(-3, -1)) // negative index are not supported, returns an empty string
	// Output:
	// 你好
	//
}

func ExampleString_Split() {
	str := oop.String("你好，世界")

	fmt.Println(str.Split("，"))
	// Output:
	// [你好 世界]
}

func ExampleString_Chunk() {
	str := oop.String("你好世界")

	fmt.Println(str.Chunk(2))
	// Output:
	// [你好 世界]
}

func ExampleString_Truncate() {
	str := oop.String("你好，世界！Hallo 世界！")

	fmt.Println(str.Truncate(10))
	// Output:
	// 你好，世界！H...
}

func ExampleString_Repeat() {
	str := oop.String("你好，世界！")

	fmt.Println(str.Repeat(2))
	// Output:
	// 你好，世界！你好，世界！
}

func ExampleString_Replace() {
	str := oop.String("你好，世界！")

	fmt.Println(str.Replace("你好", "Hello"))
	// Output:
	// Hello，世界！
}

func ExampleString_ReplaceAll() {
	str := oop.String("你好，世界！Hello, 世界！")

	fmt.Println(str.ReplaceAll("世界", "World"))
	// Output:
	// 你好，World！Hello, World！
}

func ExampleString_Search() {
	str := oop.String("你好，世界！Hello, 世界！")

	fmt.Println(str.Search("l{2,}"))
	// Output:
	// 8
}

func ExampleString_Match() {
	str := oop.String("你好，世界！Hello, 世界！")

	fmt.Println(str.Match("l{2,}"))
	// Output:
	// ll
}

func ExampleString_MatchAll() {
	str := oop.String("Hello，World！Hi, 世界！")

	fmt.Println(str.MatchAll("(?i)h{1,}"))
	// Output:
	// [H H]
}
