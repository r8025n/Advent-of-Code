package main

import(
	// "fmt"
)

const seconds, lenX, lenY = 100, 103, 101

func calculateFuturePosition(data RobotData) []int {
	x, y := data.Position[0], data.Position[1]
	xx, yy := data.Velocity[0], data.Velocity[1]

	totalMoveX, totalMoveY := x + (xx * seconds), y + (yy * seconds)
	finalPosX, finalPosY := totalMoveX % lenX, totalMoveY % lenY

	if finalPosX < 0 {
		finalPosX = lenX + finalPosX
	}
	if finalPosY < 0 {
		finalPosY = lenY + finalPosY
	}

	return []int{finalPosX, finalPosY}
}

func solvePart1(input []RobotData) int {
	futurePositions := [][]int{}

	for _, data := range input {
		futurePositions = append(futurePositions, calculateFuturePosition(data))
	}

	midX, midY := lenX/2, lenY/2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, item := range futurePositions {
		switch {
			case item[0] < midX && item[1] < midY:
				q1++
			case item[0] < midX && item[1] > midY:
				q2++
			case item[0] > midX && item[1] < midY:
				q3++
			case item[0] >midX && item[1] > midY:
				q4++
		}
	}

	return q1 * q2 * q3 * q4
}
