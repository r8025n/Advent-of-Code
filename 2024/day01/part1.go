package main

import(
	"sort"
	"math"
)

func solvePart1(list1, list2 []int) int{
	sort.Ints(list1)
	sort.Ints(list2)

	output := 0
	for i := 0; i < len(list1); i++ {
		output += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	return output
}
