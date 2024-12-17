package main

var XX = [4]int{0, 0, 1, -1}
var YY = [4]int{1, -1, 0, 0}

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
