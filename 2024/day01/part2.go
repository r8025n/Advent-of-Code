package main

import(

)

func solvePart2(list1, list2 []int) int{
	m := make(map[int]int)

	for _, locId := range(list1) {
		m[locId] = 0
	}

	for _, locId := range(list2) {
		m[locId] += 1
	}

	output := 0
	for _, locId := range(list1) {
		output += m[locId] * locId
	}

	return output
}
