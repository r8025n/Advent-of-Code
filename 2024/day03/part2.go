package main

import(
	"regexp"
	"strings"
	"strconv"
)

func solvePart2(corruptedMemories []string) int{
	pattern := `do\(\)|don't\(\)|mul\(\d+(,\d+)*\)`
	re := regexp.MustCompile(pattern)
	output := 0
	doMul := true

	for _, corruptedMemory := range corruptedMemories {
		matches := re.FindAllStringSubmatch(corruptedMemory, -1)

		for _, match := range matches {
			data := match[0]
			if strings.Contains(data, "do()"){
				doMul = true
			}else if strings.Contains(data, "don't()"){
				doMul = false
			}else{
				if doMul == true{
					numbers := strings.Split(data[4 : len(match[0]) - 1], ",")
					num1, _ := strconv.Atoi(numbers[0])
					num2, _ := strconv.Atoi(numbers[1])

					output += (num1 * num2)
				}
			}
		}
	}

	return output
}
