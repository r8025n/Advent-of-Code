package main

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
