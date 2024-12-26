package main

import(
	"fmt"
)

func simulateSingleSecond(input []RobotData, currentPositions [][]int) [][]int {
	newPositions := [][]int{}

	for i := 0; i < len(input); i++ {
		currentPosition := currentPositions[i]
		robotData := input[i]

		newPosition := []int{currentPosition[0] + robotData.Velocity[0], currentPosition[1] + robotData.Velocity[1]}
		newPosition[0] = newPosition[0] % lenX
		newPosition[1] = newPosition[1] % lenY

		if newPosition[0] < 0 {
			newPosition[0] = lenX + newPosition[0]
		}

		if newPosition[1] < 0 {
			newPosition[1] = lenY + newPosition[1]
		}

		newPositions = append(newPositions, newPosition)
	}

	return newPositions
}

func generateGrid(positions [][]int) [lenX][lenY]rune{
	var grid [lenX][lenY]rune

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			grid[i][j] = '.'
		}
	}

	for _, item := range positions {
		grid[item[0]][item[1]] = '*'
	}

	return grid
}

func detectChristmasTree(grid [lenX][lenY]rune) bool {
	for i := 0; i < lenX - 11; i++ {
		outerLoop1:
		for j := 0; j < lenY - 11; j++ {
			if grid[i][j] == '*' {
				for k := j; k < j + 10; k++ {
					if grid[i][k] != '*' {
						break outerLoop1
					}
				}
				return true
			}
		}
	}

	return false
}

func solvePart2(input []RobotData) int {
	futurePositions := [][]int{}

	for _, data := range input {
		futurePositions = append(futurePositions, data.Position)
	}

	seconds := 0
	for {
		seconds++
		futurePositions = simulateSingleSecond(input, futurePositions)
		grid := generateGrid(futurePositions)
		if detectChristmasTree(grid) {
			for _, row := range grid {
				fmt.Println(string(row[:]))
			}

			break
		}
	}
	return seconds
}
