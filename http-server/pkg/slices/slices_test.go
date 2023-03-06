package slices_test

import (
	"testing"

	"github.com/MedvedewEM/pow/pkg/slices"
	"github.com/stretchr/testify/assert"
)

func TestIntSlicesEqual(t *testing.T) {
	x := []int{}
	y := []int{}
	assert.True(t, slices.Equal(x, y))

	x = nil
	y = nil
	assert.True(t, slices.Equal(x, y))

	x = []int{}
	y = nil
	assert.True(t, slices.Equal(x, y))

	x = nil
	y = []int{}
	assert.True(t, slices.Equal(x, y))

	x = []int{1, 2, 3}
	y = nil
	assert.False(t, slices.Equal(x, y))

	x = nil
	y = []int{1, 2, 3}
	assert.False(t, slices.Equal(x, y))

	x = []int{1, 2, 3}
	y = []int{}
	assert.False(t, slices.Equal(x, y))

	x = []int{}
	y = []int{1, 2, 3}
	assert.False(t, slices.Equal(x, y))

	x = []int{1, 2}
	y = []int{1, 2, 3}
	assert.False(t, slices.Equal(x, y))

	x = []int{1, 2, 3}
	y = []int{1, 2}
	assert.False(t, slices.Equal(x, y))

	x = []int{1, 2, 3}
	y = []int{1, 2, 3}
	assert.True(t, slices.Equal(x, y))
}
