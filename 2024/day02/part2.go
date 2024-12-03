package main

import(
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

func solvePart2(reports [][]int) int{
	safe := 0

	for _, report := range(reports) {
		isSafe := checkSafety(report)

		if isSafe {
			safe += 1
		}else{
			for i, _ := range report {
				reportWithoutOneLvl := append([]int{}, report[:i]...)
				reportWithoutOneLvl = append(reportWithoutOneLvl, report[i+1:]...)
				isSafe := checkSafety(reportWithoutOneLvl)
				if isSafe {
					safe += 1
					break
				}
			}
		}
	}

	return safe
}
