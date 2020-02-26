package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_gapBinary(t *testing.T) {

	testCases := []struct {
		description string
		input       int
		expected    int
	}{
		{"should get a gap of 4", 529, 4},
		{"should get a gap of 1", 9, 2},
		{"should get a gap of 0", 15, 0},
		{"should get a gap of 0", 32, 0},
	}

	for _, tc := range testCases {

		t.Run(tc.description, func(t *testing.T) {

			biggestGap := gapBinary(tc.input)

			assert.Equal(t, tc.expected, biggestGap)
		})
	}
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

func Test_tapeEquilibrium(t *testing.T) {

	//missing := tapeEquilibrium([]int{3, 1, 2, 4, 3})
	missing := tapeEquilibrium([]int{5, 6, 2, 4, 1})

	assert.Equal(t, 4, missing)
}

func Test_frogRiverOne(t *testing.T) {
	one := frogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4})

	assert.Equal(t, 6, one)
}

func Test_log(t *testing.T) {

	i := log(1000)

	assert.Equal(t, int(math.Log2(1000)), i)
}
