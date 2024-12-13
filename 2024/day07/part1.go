package main

func recurSolvePart1(eqnValues []int, target, cumulativeVal, i int) bool {
	if i == len(eqnValues) {
		if cumulativeVal == target {
			return true
		}

		return false
	}

	if cumulativeVal > target {
		return false
	}

	opt1 := recurSolvePart1(eqnValues, target, cumulativeVal + eqnValues[i], i + 1)
 	opt2 := recurSolvePart1(eqnValues, target, cumulativeVal * eqnValues[i], i + 1)

	return opt1 || opt2
}

func solvePart1(testValueSet []int, equationValueSet[][]int) int{
	output := 0

	for t := 0; t < len(testValueSet); t++ {
		testValue := testValueSet[t]
		equationValues := equationValueSet[t]

		if recurSolvePart1(equationValues, testValue, equationValues[0], 1) {
			output += testValue
		}
	}

	return output
}
