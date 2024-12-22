package main

import(
	// "fmt"
)

func bfsSolvePart1(input [][]rune, visited [][]int, i, j int) int {
	q := Queue{}
	target := input[i][j]
	areas, perimeters := 1, 0
	visited[i][j] = 1
	input[i][j] = 0
	q.Enqueue([]int{i, j})

	for !q.Empty() {
		topItem := q.Dequeue()
		x, y := topItem[0], topItem[1]

		for k := 0; k < 4; k++ {
			xx := x + XX[k]
			yy := y + YY[k]

			if xx < 0 || yy < 0 || xx >= len(input) || yy >= len(input[0]) {
				perimeters++
				continue
			}

			if visited[xx][yy] == 1 {
				continue
			}

			if input[xx][yy] == target {
				q.Enqueue([]int{xx, yy})
				visited[xx][yy] = 1
				areas++
				input[xx][yy] = 0
			} else {
				perimeters++
			}
		}
	}

	return areas * perimeters
}

func solvePart1(input [][]rune) int {
	output := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			visited := make([][]int, len(input))
			for i := 0; i < len(visited); i++ {
				visited[i] = make([]int, len(input[0]))
			}
			if input[i][j] != 0 {
				output += bfsSolvePart1(input, visited, i, j)
			}
		}
	}

	return output
}
