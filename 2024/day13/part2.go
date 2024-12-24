package main

import(
	// "fmt"
	// "math"
)

func algebricSolvePart2(data PrizeData) int {
	// data.A[0]*X + data.B[0].Y = data.Target[0]
	// data.A[1]*X + data.B[1].Y = data.Target[1]
	tokenCount := 0

	bPressCount := float64(data.A[0] * data.Target[1] - data.A[1] * data.Target[0]) / float64(data.A[0] * data.B[1] - data.A[1] * data.B[0])
	aPressCount := float64(data.Target[0] - data.B[0] * int(bPressCount)) / float64(data.A[0])

	if aPressCount == float64(int(aPressCount)) && bPressCount == float64(int(bPressCount)) {
		tokenCount = int(3 * aPressCount) + int(bPressCount)
	}

	return tokenCount
}

func solvePart2(input []PrizeData) int {
	output := 0

	for _, data := range input {
		data.Target[0] = data.Target[0] + 10000000000000
		data.Target[1] = data.Target[1] + 10000000000000

		output += algebricSolvePart2(data)
	}

	return output
}
