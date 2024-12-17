package main

import(
	// "fmt"
)

func recurSolvePart2(input [][]int, x, y, prev int) int {
	if x < 0 || y < 0 || x >= len(input) || y >= len(input[0]) {
		return 0
	}

	current := input[x][y]

	if current <= prev || current - prev != 1 {
		return 0
	}

	if current == 9 {
		return 1
	}

	output := 0
	for i := 0; i < 4; i++ {
		output += recurSolvePart2(input, x + XX[i], y + YY[i], current)
	}

	return output
}

func solvePart2(input [][]int) int {
	output := 0

	trailHeads := findTrailHeads(input)

	for _, trailHead := range trailHeads{
		output += recurSolvePart2(input, trailHead[0], trailHead[1], -1)
	}


	return output
}
