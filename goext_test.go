package goext_test

import (
	"errors"
	"testing"
	"time"

	"github.com/ayonli/goext"
	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	texture := func(good bool) string {
		if !good {
			panic("something went wrong")
		}

		return "everything looks fine"
	}

	call := goext.Wrap(func(args ...any) string {
		text := texture(args[0].(bool))
		return text
	})

	res, err := call(true)

	assert.Equal(t, "everything looks fine", res)
	assert.Nil(t, err)
}

func TestThrottle(suit *testing.T) {
	suit.Run("failedWithoutKey", func(t *testing.T) {
		fn := goext.Throttle[int](func(arg int) (int, error) {
			if arg == 1 {
				return 0, errors.New("something went wrong")
			} else {
				return 0, errors.New("something else went wrong")
			}
		}, time.Millisecond*5, "")

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
		}, time.Millisecond*5, "bar")(1)
		_, err2 := goext.Throttle[int](func(arg int) (int, error) {
			if arg == 1 {
				return 0, errors.New("something went wrong")
			} else {
				return 0, errors.New("something else went wrong")
			}
		}, time.Millisecond*5, "bar")(2)

		assert.Equal(t, errors.New("something went wrong"), err1)
		assert.Equal(t, err1, err2)

		time.Sleep(time.Millisecond * 6)
		_, err3 := goext.Throttle[int](func(arg int) (int, error) {
			if arg == 1 {
				return 0, errors.New("something went wrong")
			} else {
				return 0, errors.New("something else went wrong")
			}
		}, time.Millisecond*5, "bar")(2)
		assert.Equal(t, errors.New("something else went wrong"), err3)
	})
}
