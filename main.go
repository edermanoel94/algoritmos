package main

import (
	"strconv"
)

const (
	ZERO = "0"
	ONE  = "1"
)

// TODO: refactor
func gapBinary(N int) int {

	binaries := strconv.FormatInt(int64(N), 2)

	var firstIndex, currentGap, biggestGap int

	for index, binary := range binaries {
		currentBinary := string(binary)
		if currentBinary == ONE {
			firstIndex = index
			break
		}
	}

	for i := firstIndex; i < len(binaries); i++ {
		currentBinary := string(binaries[i])
		if currentBinary == ZERO {
			currentGap += 1
		} else {
			if currentGap > biggestGap {
				biggestGap = currentGap
			}
			currentGap = 0
		}
	}

	return biggestGap
}

// TODO: make a gapBinary recursively

func cyclicRotation(A []int, K int) []int {

	// create a new array
	total := len(A)

	arr := make([]int, total)

	// iterate over array
	for i := 0; i < total; i++ {

		// i is the current index of element of array A
		// K is the number of many shifted will occurs
		// i = 0 and k = 4, 4 % 4 = 0
		// i = 1 and k = 4, 5 % 4 = 1
		// i = 2 and k = 4, 6 % 4 = 2
		// i = 3 and k = 4, 7 % 4 = 3
		cyclic := (i + K) % total

		// change position of all elements
		arr[cyclic] = A[i]
	}

	return arr
}

// TODO: make a oddOccurrences with map

func oddOccurrencesWithXor(A []int) int {
	result := 0
	for _, elem := range A {
		// why xor operator?
		// when we use xor operator in many integers,
		// the number which repeat is ignored and only who stay alone
		result ^= elem
	}
	return result
}

func countMinimalJumps(X, Y, D int) int {

	// the frog is located at position X and wants to get a posisition
	// greater than or equal to Y. the frog jumps a fixed distance D




	return 0
}


