package main

import(
	"strconv"
)

var solutionMap = make(map[string]int)

func recurSolve(input string, len, iter int) int {
	if input == "" {
		return 0
	}
	if iter == 0 {
		return len
	}

	mapKey := input+"#"+strconv.Itoa(iter)

	if _, exists := solutionMap[mapKey]; exists {
		return solutionMap[mapKey]
	}

	out1, out2 := transformNumber(input)
	opt1 := recurSolve(out1, len, iter - 1)
	opt2 := recurSolve(out2, len, iter - 1)

	solutionMap[mapKey] = opt1 + opt2
	return solutionMap[mapKey]
}

func solvePart2(input []string) int {
	output := 0
	iter := 75

	for _, item := range input {
		output += recurSolve(item, 1, iter)
	}

	return output
}
