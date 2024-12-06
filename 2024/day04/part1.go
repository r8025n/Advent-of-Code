package main

import(
)

func countOverlapping(text, substring string) int {
	count := 0
	for i := 0; i <= (len(text) - len(substring)); i++ {
		if text[i : i + len(substring)] == substring {
			count += 1
		}
	}

	return count
}

func solvePart1(input []string) int{
	output := 0
	substringsToFind := []string{"XMAS", "SAMX"}
	possibleCombinations := append([]string{}, input...)

	for j := 0; j < len(input[0]); j++ {
		tempString := ""
		for i := 0; i < len(input); i++ {
			tempString = tempString + string(input[i][j])
		}
		possibleCombinations = append(possibleCombinations, tempString)
	}

	for k :=0; k < len(input); k++ {
		tempString := ""
		for i, j := k, 0; i < len(input) && j < len(input[0]); i, j = i+1, j+1 {
			tempString = tempString + string(input[i][j])
		}
		possibleCombinations = append(possibleCombinations, tempString)
	}

	for k := 1; k < len(input[0]); k++ {
		tempString := ""
		for i, j := 0, k; i < len(input) && j < len(input[0]); i,j = i+1, j+1 {
			tempString = tempString + string(input[i][j])
		}
		possibleCombinations = append(possibleCombinations, tempString)
	}

	for k := len(input) - 1; k >= 0; k-- {
		tempString := ""
		for i, j := k, 0; i >= 0 && j < len(input[0]); i, j = i-1, j+1 {
			tempString = tempString + string(input[i][j])
		}
		possibleCombinations = append(possibleCombinations, tempString)
	}

	for k := 0; k < len(input[0]); k++ {
		tempString := ""
		for i, j := len(input) - 1, k + 1; i >= 0 && j < len(input[0]); i,j = i-1, j+1 {
			tempString = tempString + string(input[i][j])
		}
		possibleCombinations = append(possibleCombinations, tempString)
	}

	for _, text := range possibleCombinations {
		for _, stringToFind := range substringsToFind {
			output += countOverlapping(text, stringToFind)
		}
	}

	return output
}
