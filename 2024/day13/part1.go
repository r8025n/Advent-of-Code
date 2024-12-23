package main

import(
	// "fmt"
	"math"
)

func recurSolvePart1(data PrizeData, memory [][]int, x, y, aCount, bCount int) int {
	if x > data.Target[0] || y > data.Target[1] {
		return math.MaxInt64
	}

	if aCount > 100 || bCount > 100 {
		return math.MaxInt64
	}


	tokenCount := aCount * 3 + bCount

	if x == data.Target[0] && y == data.Target[1] {
		return tokenCount
	}


	if memory[aCount][bCount] != 0 {
		return  memory[aCount][bCount]
	}

	pressA := recurSolvePart1(data, memory, x + data.A[0], y + data.A[1], aCount + 1, bCount)
	pressB := recurSolvePart1(data, memory, x + data.B[0], y + data.B[1], aCount, bCount + 1)

	memory[aCount][bCount] = min(pressA, pressB)

	return memory[aCount][bCount]
}

func solvePart1(input []PrizeData) int {
	output := 0

	for _, data := range input {
		memory := make([][]int, 105)
		for i := range memory {
			memory[i] = make([]int, 105)
		}

		tokenCount := recurSolvePart1(data, memory, 0, 0, 0, 0)
		if tokenCount != math.MaxInt64 {
			output += tokenCount
		}
	}

	return output
}
