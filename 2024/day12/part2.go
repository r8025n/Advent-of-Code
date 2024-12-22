package main

import(
	// "fmt"
)

func findAndCountCorners(input [][]rune, i, j int) int {
	corners := 0
	plot := input[i][j]
	var up, down, left, right rune
	var upperLeft, lowerLeft, upperRight, lowerRight rune

	if  j - 1 >= 0 { left = input[i][j - 1] }
	if j + 1 < len(input[0]) { right = input[i][j + 1] }
	if i - 1 >= 0 { up = input[i - 1][j] }
	if i + 1 < len(input) { down = input[i + 1][j] }
	if i - 1 >=0 && j - 1 >= 0 { upperLeft = input[i - 1][j - 1] }
	if i + 1 < len(input) && j - 1 >= 0 { lowerLeft = input[i + 1][j - 1] }
	if i - 1 >= 0 && j + 1 < len(input[0]) { upperRight = input[i - 1][j + 1] }
	if i + 1 <len(input) && j + 1 < len(input[0]) { lowerRight = input[i + 1][j + 1] }

	if ! (plot == left && plot == right) && !(plot == up && plot == down) {
		// outward corners
		if up != plot && left != plot {
			corners++
		}
		if up != plot && right != plot {
			corners++
		}
		if down != plot && right != plot {
			corners++
		}
		if down != plot && left != plot {
			corners++
		}
	}

	// inward corners
	if plot != left && plot == upperLeft && up == plot {
		corners++
	}
	if plot != left && plot == lowerLeft && down == plot {
		corners++
	}
	if plot != right && plot == upperRight && up == plot {
		corners++
	}
	if plot != right && plot == lowerRight && down == plot {
		corners++
	}

	return corners
}

func bfsSolvePart2(input [][]rune, visited [][]int, i, j int) int {
	q := Queue{}
	target := input[i][j]
	area, sides := 1, 0
	sides += findAndCountCorners(input, i, j)
	visited[i][j] = 1

	q.Enqueue([]int{i, j})
	for !q.Empty() {
		topItem := q.Dequeue()
		x, y := topItem[0], topItem[1]

		for k := 0; k < 4; k++ {
			xx := x + XX[k]
			yy := y + YY[k]

			if xx < 0 || yy < 0 || xx >= len(input) || yy >= len(input[0]) {
				continue
			}

			if visited[xx][yy] == 1 {
				continue
			}

			if input[xx][yy] == target {
				q.Enqueue([]int{xx, yy})
				visited[xx][yy] = 1
				sides += findAndCountCorners(input, xx, yy)
				area++

			}
		}
	}

	return area * sides
}

func solvePart2(input [][]rune) int {
	output := 0
	visited := make([][]int, len(input))

	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(input[0]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if visited[i][j] == 0 {
				output += bfsSolvePart2(input, visited, i, j)
			}
		}
	}

	return output
}
