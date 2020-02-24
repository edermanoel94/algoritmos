package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_gapBinary(t *testing.T) {

	biggestGap := gapBinary(529)

	assert.Equal(t, 4, biggestGap)
}

func Test_cyclicRotation(t *testing.T) {

	arr := cyclicRotation([]int{3, 8, 9, 7, 6}, 3)

	assert.Equal(t, []int{9, 7, 6, 3, 8}, arr)
}

func Test_oddOccurrencesWithXor(t *testing.T) {

	lonely := oddOccurrencesWithXor([]int{9,3,9,3,9,7,9})

	assert.Equal(t, 7, lonely)
}

func Test_countMinimalJumps(t *testing.T) {

	jumps := countMinimalJumps(1, 2, 3)

	assert.Equal(t, 2, jumps)
}
