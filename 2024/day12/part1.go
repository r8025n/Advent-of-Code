package main

import(
	// "fmt"
)

var XX = []int{0, 0, 1, -1}
var YY = []int{1, -1, 0, 0}

func bfsSolve(input [][]rune, visited [][]int, i, j int) int {
	q := Queue{}
	target := input[i][j]
	area, perimeter := 1, 0
	input[i][j] = 0
	visited[i][j] = 1
	q.Enqueue([]int{i, j})

	for !q.Empty() {
		topItem := q.Dequeue()
		x, y := topItem[0], topItem[1]

		for k := 0; k < 4; k++ {
			xx := x + XX[k]
			yy := y + YY[k]

			if xx < 0 || yy < 0 || xx >= len(input) || yy >= len(input[0]) {
				perimeter++
				continue
			}

			if visited[xx][yy] == 1 {
				continue
			}

			if input[xx][yy] == target {
				q.Enqueue([]int{xx, yy})
				visited[xx][yy] = 1
				input[xx][yy] = 0
				area++
			} else {
				perimeter++
			}
		}
	}

	return area * perimeter
}

func solvePart1(input [][]rune) int {
	output := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i][j] != 0 {
				visited := make([][]int, len(input))

				for i := 0; i < len(visited); i++ {
					visited[i] = make([]int, len(input[0]))
				}

				output += bfsSolve(input, visited, i, j)
			}
		}
	}

	return output
}
