package mathx_test

import (
	"fmt"

	"github.com/ayonli/goext/mathx"
)

func ExampleMax() {
	fmt.Println(mathx.Max(1))
	fmt.Println(mathx.Max(1, 2))
	fmt.Println(mathx.Max(1, 2, 0.5))
	// Output:
	// 1
	// 2
	// 2
}

func ExampleMin() {
	fmt.Println(mathx.Min(1))
	fmt.Println(mathx.Min(1, 2))
	fmt.Println(mathx.Min(1, 2, 0.5))
	// Output:
	// 1
	// 1
	// 0.5
}

func ExampleSum() {
	fmt.Println(mathx.Sum(1))
	fmt.Println(mathx.Sum(1, 2))
	fmt.Println(mathx.Sum(1, 2, 0.5))
	// Output:
	// 1
	// 3
	// 3.5
}

func ExampleAvg() {
	fmt.Println(mathx.Avg(1))
	fmt.Println(mathx.Avg(1, 2))
	fmt.Println(mathx.Avg(1, 2, 3))
	// Output:
	// 1
	// 1.5
	// 2
}

func ExampleProduct() {
	fmt.Println(mathx.Product(1))
	fmt.Println(mathx.Product(1, 2))
	fmt.Println(mathx.Product(1, 2, 0.5))
	// Output:
	// 1
	// 2
	// 1
}
