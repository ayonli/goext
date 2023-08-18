package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	v1 := Max(1.0)
	v2 := Max(2.0, 1.0)
	v3 := Max(2.0, 1.0, 3.0)

	assert.Equal(t, 1.0, v1)
	assert.Equal(t, 2.0, v2)
	assert.Equal(t, 3.0, v3)
}

func TestMin(t *testing.T) {
	v1 := Min(1.0)
	v2 := Min(2.0, 1.0)
	v3 := Min(2.0, 1.0, 3.0)

	assert.Equal(t, 1.0, v1)
	assert.Equal(t, 1.0, v2)
	assert.Equal(t, 1.0, v3)
}

func TestSum(t *testing.T) {
	sum1 := Sum(1.0)
	sum2 := Sum(1.0, 2.0)
	sum3 := Sum(1.0, 2.0, 3.0)

	assert.Equal(t, 1.0, sum1)
	assert.Equal(t, 3.0, sum2)
	assert.Equal(t, 6.0, sum3)
}

func TestProduct(t *testing.T) {
	sum1 := Product(1.0)
	sum2 := Product(1.0, 2.0)
	sum3 := Product(1.0, 2.0, 3.0)

	assert.Equal(t, 1.0, sum1)
	assert.Equal(t, 2.0, sum2)
	assert.Equal(t, 6.0, sum3)
}
