package main

import(
	"regexp"
	"strings"
	"strconv"
)

func solvePart1(corruptedMemories []string) int{
	pattern := `mul\(\d+(,\d+)*\)`
	re := regexp.MustCompile(pattern)
	output := 0



	for _, corruptedMemory := range corruptedMemories {
		matches := re.FindAllStringSubmatch(corruptedMemory, -1)

		for _, match := range matches {
			data := match[0][4 : len(match[0]) - 1]
			numbers := strings.Split(data, ",")
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])

			output += (num1 * num2)
		}
	}

	return output
}
