package main

import(
	"strings"
	"strconv"
	"sort"
	"math"
)

func solvePart1(input []string) int{
	var list1, list2 []int
	for _, line := range(input) {
		line = strings.Trim(line, " ")
		locIds := strings.Fields(line)
		locId1, _ := strconv.Atoi(locIds[0])
		locId2, _ := strconv.Atoi(locIds[1])
		list1 = append(list1, locId1)
		list2 = append(list2, locId2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	output := 0
	for i := 0; i < len(list1); i++ {
		output += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	return output
}
