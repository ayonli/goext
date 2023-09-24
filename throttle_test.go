package goext_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ayonli/goext"
	"github.com/stretchr/testify/assert"
)

func ExampleThrottle_withoutKey() {
	fn := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "", false)

	fmt.Println(fn(1))
	fmt.Println(fn(2))

	time.Sleep(time.Millisecond * 6)
	fmt.Println(fn(3))

	// Output:
	// 2 <nil>
	// 2 <nil>
	// 6 <nil>
}

func ExampleThrottle_withKey() {
	res1, err1 := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "foo", false)(1)
	res2, err2 := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "foo", false)(2)

	fmt.Println(res1, err1)
	fmt.Println(res2, err2)

	time.Sleep(time.Millisecond * 6)
	res3, err3 := goext.Throttle[int](func(arg int) (int, error) {
		return arg * 2, nil
	}, time.Millisecond*5, "foo", false)(3)
	fmt.Println(res3, err3)

	// Output:
	// 2 <nil>
	// 2 <nil>
	// 6 <nil>
}

func TestThrottle(suit *testing.T) {
	suit.Run("failedWithoutKey", func(t *testing.T) {
		fn := goext.Throttle[int](func(arg int) (int, error) {
			if arg == 1 {
				return 0, errors.New("something went wrong")
			} else {
				return 0, errors.New("something else went wrong")
			}
		}, time.Millisecond*5, "", false)

		_, err1 := fn(1)
		_, err2 := fn(2)
		assert.Equal(t, errors.New("something went wrong"), err1)
		assert.Equal(t, err1, err2)

		time.Sleep(time.Millisecond * 6)
		_, err3 := fn(2)
		assert.Equal(t, errors.New("something else went wrong"), err3)
	})

	suit.Run("failedWithKey", func(t *testing.T) {
		_, err1 := goext.Throttle[int](func(arg int) (int, error) {
			if arg == 1 {
				return 0, errors.New("something went wrong")
			} else {
				return 0, errors.New("something else went wrong")
			}
		}, time.Millisecond*5, "bar", false)(1)
		_, err2 := goext.Throttle[int](func(arg int) (int, error) {
			if arg == 1 {
				return 0, errors.New("something went wrong")
			} else {
				return 0, errors.New("something else went wrong")
			}
		}, time.Millisecond*5, "bar", false)(2)

		assert.Equal(t, errors.New("something went wrong"), err1)
		assert.Equal(t, err1, err2)

		time.Sleep(time.Millisecond * 6)
		_, err3 := goext.Throttle[int](func(arg int) (int, error) {
			if arg == 1 {
				return 0, errors.New("something went wrong")
			} else {
				return 0, errors.New("something else went wrong")
			}
		}, time.Millisecond*5, "bar", false)(2)
		assert.Equal(t, errors.New("something else went wrong"), err3)
	})

	suit.Run("noWait", func(t *testing.T) {
		fn := goext.Throttle[int](func(arg int) (int, error) {
			return arg * 2, nil
		}, time.Millisecond*5, "", true)

		res1, _ := fn(1)

		time.Sleep(time.Millisecond * 6)
		res2, _ := fn(2)
		assert.Equal(t, res2, res1)

		time.Sleep(time.Millisecond)
		res3, _ := fn(3)
		assert.Equal(t, res3, 4)
	})
}
