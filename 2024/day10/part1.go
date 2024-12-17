package main

import(
	// "fmt"
)

var XX = [4]int{0, 0, 1, -1}
var YY = [4]int{1, -1, 0, 0}
var count = 0
func findTrailHeads(input [][]int) [][]int {
	trailHeads := [][]int{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == 0 {
				trailHeads = append(trailHeads, []int{i, j})
			}
		}
	}

	return trailHeads
}

func recurSolve(input [][]int, x, y, prev int) {
	if x < 0 || y < 0 || x >= len(input) || y >= len(input[0]) {
		return
	}

	current := input[x][y]

	if current <= prev || current - prev != 1 {
		return
	}

	if current == 9 {
		input[x][y] = -1
		return
	}

	for i := 0; i < 4; i++ {
		recurSolve(input, x + XX[i], y + YY[i], current)
	}

	return
}

func solvePart1(input [][]int) int {
	output := 0

	trailHeads := findTrailHeads(input)

	for _, trailHead := range trailHeads{
		successCount := 0

		recurSolve(input, trailHead[0], trailHead[1], -1)

		for i:= 0; i < len(input); i++ {
			for j := 0; j < len(input[0]); j++ {
				if input[i][j] == -1 {
					successCount += 1
					input[i][j] = 9
				}
			}
		}
		output += successCount
	}


	return output
}
