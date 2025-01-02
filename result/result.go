package result

import (
	"errors"
	"fmt"
)

func Wrap[T any](fn func() (value T, err error)) (value T, err error) {
	defer func() {
		if re := recover(); re != nil {
			if _err, ok := re.(error); ok {
				err = _err
			} else if str, ok := re.(string); ok {
				err = errors.New(str)
			} else {
				err = errors.New(fmt.Sprint(re))
			}
		}
	}()

	return fn()
}

func Unwrap[R any](value R, err error) R {
	if err != nil {
		panic(err)
	}

	return value
}
