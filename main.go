package main

import (
	"math"
	"strconv"
)

const (
	ZERO = "0"
	ONE  = "1"
)

// TODO: refactor
// TODO: make a gapBinary recursively
// Complexity O(n+m)
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

// Worst performance, with O(y-x)
func countMinimalJumps(X, Y, D int) int {
	// the frog is located at position X and wants to get a posisition
	// greater than or equal to Y. the frog jumps a fixed distance D
	// X = 10
	// Y = 85
	// D = 30
	// X + 30 + 30 + 30 >= Y
	jumps := 0
	for X < Y {
		X += D
		jumps += 1
	}
	return jumps
}

// Better performance, with O(1)
func countMinimalJumpsWithoutLoop(X, Y, D int) int {
	// Check if is divisible
	if (Y-X)%D == 0 {
		return (Y - X) / D
	}
	return (Y-X)/D + 1
}

// TODO: using with XOR
func permMissingElem(A []int) int {

	// [2, 3, 1, 5]
	// n = 4 + 1 (missing element)
	// (n * (n + 1)) / 2; = 15
	// 2+3+1+5 = 11

	n := len(A) + 1

	totalSumAllElements := (n * (1 + n)) / 2

	sum := 0

	for _, elem := range A {
		sum += elem
	}

	return totalSumAllElements - sum
}

func tapeEquilibrium(A []int) int {

	arraySum := func(arr []int) int {
		sum := 0
		for _, elem := range arr {
			sum += elem
		}
		return sum
	}
	minusAbs := func(x, y int) int {
		return int(math.Abs(float64(x) - float64(y)))
	}

	left := A[0]
	right := arraySum(A[1:])
	minimumDistance := minusAbs(left, right)

	for i := 1; i < len(A)-1; i++ {

		left += A[i]
		right -= A[i]

		distance := minusAbs(left, right)

		if distance < minimumDistance {
			minimumDistance = distance
		}
	}

	return minimumDistance
}

func frogRiverOne(X int, A []int) int {

	for _, elem := range A {
		if X == elem {
			return len(A) - 1
		}
	}

	return 0
}

func log(N int) int {
	result := 0
	for N > 1 {
		N = N / 2
		result += 1
	}
	return result
}
