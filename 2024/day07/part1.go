package main

func recur(eqnValues []int, target, cumulativeVal, i int) bool {
	if i == len(eqnValues) {
		if cumulativeVal == target {
			return true
		}

		return false
	}

	opt1 := recur(eqnValues, target, cumulativeVal + eqnValues[i], i + 1)
 	opt2 := recur(eqnValues, target, cumulativeVal * eqnValues[i], i + 1)

	return opt1 || opt2
}

func solvePart1(testValueSet []int, equationValueSet[][]int) int{
	output := 0

	for t := 0; t < len(testValueSet); t++ {
		testValue := testValueSet[t]
		equationValues := equationValueSet[t]

		if recur(equationValues, testValue, equationValues[0], 1) {
			output += testValue
		}
	}

	return output
}
