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

	lonely := oddOccurrencesWithXor([]int{9, 3, 9, 3, 9, 7, 9})

	assert.Equal(t, 7, lonely)
}

func Test_countMinimalJumps(t *testing.T) {

	// TODO: add a table driven tests
	//jumps := countMinimalJumps(10, 85, 30)
	jumps := countMinimalJumps(1, 5, 2)

	assert.Equal(t, 2, jumps)
}

func Test_countMinimalJumpsWithoutLoop(t *testing.T) {

	// TODO: add a table driven tests
	// jumps := countMinimalJumpsWithoutLoop(10, 85, 30)
	jumps := countMinimalJumpsWithoutLoop(10, 135, 35)

	assert.Equal(t, 4, jumps)
}

func Test_permMissingElem(t *testing.T) {

	missing := permMissingElem([]int{2, 3, 1, 5})

	assert.Equal(t, 4, missing)
}
