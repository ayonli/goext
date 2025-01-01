package result_test

import (
	"fmt"
	"strconv"

	"github.com/ayonli/goext/result"
)

func ExampleWrap() {
	mathAdd := func(input1 string, input2 string) (int, error) {
		return result.Wrap(func() (int, error) {
			num1 := result.Unwrap(strconv.Atoi(input1))
			num2 := result.Unwrap(strconv.Atoi(input2))
			return num1 + num2, nil
		})
	}

	res, err := mathAdd("10", "20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	// Output:
	// 30
}

func ExampleUnwrap() {
	mathAdd := func(input1 string, input2 string) (int, error) {
		return result.Wrap(func() (int, error) {
			num1 := result.Unwrap(strconv.Atoi(input1))
			num2 := result.Unwrap(strconv.Atoi(input2))
			return num1 + num2, nil
		})
	}

	res, err := mathAdd("10", "20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	// Output:
	// 30
}
