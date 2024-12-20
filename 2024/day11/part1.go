package main

import(
	"strconv"
)

func simulateSingleBlink(input []string) []string{
	inputLength := len(input)

	for i := 0; i < inputLength; i++ {
		if input[i] == "0" {
			input[i] = "1"
		}else if len(input[i]) % 2 == 0 {
			new1 := removeLeadingZeros(input[i][:len(input[i])/2])
			new2 := removeLeadingZeros(input[i][len(input[i])/2:])
			left := append(input[:i], new1)
			right := append([]string{new2}, input[i+1:]...)
			input = append(left, right...)
			i++
			inputLength += 1
		} else{
			stoneVal, _ := strconv.Atoi(input[i])
			stoneVal = stoneVal * 2024;
			input[i] = strconv.Itoa(stoneVal)
		}
	}

	return input
}

func solvePart1(input []string) int {
	iter := 25

	for i := 0; i < iter; i++ {
		input = simulateSingleBlink(input)
	}

	return len(input)
}
