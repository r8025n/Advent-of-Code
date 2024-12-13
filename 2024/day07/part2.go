package main

import(
	"strconv"
)

func combineIntegers(a, b int) int {
	combined := strconv.Itoa(a) + strconv.Itoa(b)

	ab, _ := strconv.Atoi(combined)

	return ab
}

func recurSolvePart2(eqnValues []int, target, cumulativeVal, i int) bool {
	if i == len(eqnValues) {
		if cumulativeVal == target {
			return true
		}

		return false
	}

	if cumulativeVal > target {
		return false
	}

	opt1 := recurSolvePart2(eqnValues, target, cumulativeVal + eqnValues[i], i + 1)
 	opt2 := recurSolvePart2(eqnValues, target, cumulativeVal * eqnValues[i], i + 1)
  	opt3 := recurSolvePart2(eqnValues, target, combineIntegers(cumulativeVal, eqnValues[i]), i + 1)

	return opt1 || opt2 || opt3
}

func solvePart2(testValueSet []int, equationValueSet[][]int) int{
	output := 0

	for t := 0; t < len(testValueSet); t++ {
		testValue := testValueSet[t]
		equationValues := equationValueSet[t]

		if recurSolvePart2(equationValues, testValue, equationValues[0], 1) {
			output += testValue
		}
	}

	return output
}
