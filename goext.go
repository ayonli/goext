package goext

import (
	"errors"
	"fmt"
)

// Try runs a function in a safe context where if it or what's inside it panics, the panic reason
// can be caught and returned as a normal error.
func Try[R any](fn func() (R, error)) (res R, err error) {
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

	_res, _err := fn()

	if _err != nil {
		err = _err
	} else {
		res = _res
	}

	return res, err
}
