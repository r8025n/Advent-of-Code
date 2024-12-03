package main

import (
	"math"
)

func checkSafety(report []int) bool {
	totalDiff := int(math.Abs(float64(report[0]) - float64(report[len(report) - 1])))
	sumOfDiff := 0

	for i := 1; i < len(report); i++ {
		lvlDiff := int(math.Abs(float64(report[i]) - float64(report[i - 1])))
		if lvlDiff < 1 || lvlDiff > 3 {
			return false
		}
		sumOfDiff += lvlDiff
	}
	if totalDiff != sumOfDiff {
		return false
	}

	return true
}
