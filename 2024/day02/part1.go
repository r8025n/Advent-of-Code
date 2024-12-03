package main

import(
	// "fmt"
	"math"
)

func solvePart1(reports [][]int) int{
	safe := 0
	outer:
	for _, report := range(reports) {
		totalDiff := int(math.Abs(float64(report[0]) - float64(report[len(report) - 1])))
		sumOfDiff := 0

		for i := 1; i < len(report); i++ {
			lvlDiff := int(math.Abs(float64(report[i]) - float64(report[i - 1])))
			if lvlDiff < 1 || lvlDiff > 3 {
				continue outer
			}
			sumOfDiff += lvlDiff
		}
		if totalDiff == sumOfDiff {
			safe += 1
		}
	}

	return safe
}
