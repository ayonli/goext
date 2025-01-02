package result_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ayonli/goext/result"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
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

func withOtherError() (string, error) {
	return "", errors.WithStack(errors.New("something went wrong"))
}

func TestWrap(t *testing.T) {
	mathAdd := func(input1 string, input2 string) (int, error) {
		return result.Wrap(func() (int, error) {
			num1 := result.Unwrap(strconv.Atoi(input1))
			num2 := result.Unwrap(strconv.Atoi(input2))
			return num1 + num2, nil
		})
	}

	res, err := mathAdd("10", "b")
	assert.Equal(t, 0, res)
	assert.Equal(t, "strconv.Atoi: parsing \"b\": invalid syntax", err.Error())

	mathAdd2 := func(input1 string, input2 string) (int, error) {
		return result.Wrap(func() (int, error) {
			num1 := result.Unwrap(strconv.Atoi(input1))
			num2 := result.Unwrap(strconv.Atoi(input2))
			result.Unwrap(withOtherError())
			return num1 + num2, nil
		})
	}

	_, err = mathAdd2("10", "20")
	assert.Equal(t, "something went wrong", err.Error())
	assert.Contains(t, fmt.Sprintf("%+v", err), "result_test.withOtherError")
}
