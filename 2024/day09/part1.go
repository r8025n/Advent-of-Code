package main

import(
	// "fmt"
	"strconv"
)

func solvePart1(input string) int{
	output := 0
	inputs := []rune(input)

	expandedInput := []int{}
	id := 0
	for i := 0; i < len(inputs); i++ {
		blockSize, _ := strconv.Atoi(string(inputs[i]))
		if i % 2 == 0 {
			for j := 0; j < blockSize; j++ {
				expandedInput = append(expandedInput, id)
			}
			id++
		}else{
			for j := 0; j < blockSize; j++ {
				expandedInput = append(expandedInput, -1)
			}
		}
	}

	k := len(expandedInput) - 1
	outer:
	for i := 0; i < len(expandedInput); i++ {
		if expandedInput[i] != -1 {
			output += i * expandedInput[i]
		}else{
			for expandedInput[k] == -1 {
				k--
				if k <= i {
					break outer
				}
			}
			output += i * expandedInput[k]
			k--
		}

		if k <= i {
			break
		}
	}

	return output
}
