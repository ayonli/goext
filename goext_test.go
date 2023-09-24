package goext_test

import (
	"testing"

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
