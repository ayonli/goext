package number_test

import (
	"fmt"

	"github.com/ayonli/goext"
	"github.com/ayonli/goext/number"
)

func ExampleRandom() {
	fmt.Println(number.Random(0, 0))

	num := number.Random(0, 10)
	fmt.Println(num >= 0 && num <= 10)
	// Output:
	// 0
	// true
}

func ExampleSequence() {
	seq := number.Sequence(0, 9, 1, false)
	fmt.Println(goext.ReadAll(seq))
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleSequence_step() {
	seq := number.Sequence(0, 9, 2, false)
	fmt.Println(goext.ReadAll(seq))
	// Output:
	// [0 2 4 6 8]
}

func ExampleSequence_loop() {
	seq := number.Sequence(0, 2, 1, true)
	numbers := []int{}

	for num := range seq {
		numbers = append(numbers, num)

		if len(numbers) == 9 {
			break
		}
	}

	fmt.Println(numbers)
	// Output:
	// [0 1 2 0 1 2 0 1 2]
}

func ExampleIsNumeric() {
	fmt.Println(number.IsNumeric(123))
	fmt.Println(number.IsNumeric("123"))
	fmt.Println(number.IsNumeric("1.23"))
	fmt.Println(number.IsNumeric("0b01"))
	fmt.Println(number.IsNumeric("0o123"))
	fmt.Println(number.IsNumeric("0x12ab"))
	fmt.Println(number.IsNumeric("abc"))
	fmt.Println(number.IsNumeric([]int{}))
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
}

func ExampleParse() {
	fmt.Println(number.Parse("123"))
	fmt.Println(number.Parse("1.23"))
	fmt.Println(number.Parse("0b01"))
	fmt.Println(number.Parse("0o123"))
	fmt.Println(number.Parse("0x12ab"))
	fmt.Println(number.Parse("abc"))
	// Output:
	// 123 true
	// 1.23 true
	// 1 true
	// 83 true
	// 4779 true
	// 0 false
}
