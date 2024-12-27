package main

func min(vals []int) int {
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min
}

func findTarget(input []string, target byte) (int, int) {
	x, y := 0, 0
	outerLoop:
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == target {
				x, y = i, j
				break outerLoop
			}
		}
	}

	return x, y
}
