package main

func solvePart1(reports [][]int) int{
	safe := 0

	for _, report := range(reports) {
		isSafe := checkSafety(report)

		if isSafe {
			safe += 1
		}
	}

	return safe
}
