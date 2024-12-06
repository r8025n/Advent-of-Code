package main

import(

)

func solvePart2(input []string) int{
	output := 0
	target := []string{"MAS", "SAM"}

	for i := 1; i < len(input) - 1; i++ {
		for j := 1; j < len(input[0]) - 1; j++ {
			if string(input[i][j]) == "A" {
				leftToRight := string(input[i - 1][j - 1]) + string(input[i][j]) + string(input[i + 1][j + 1])
				rightToLeft := string(input[i - 1][j + 1]) + string(input[i][j]) + string(input[i + 1][j - 1])

			 	if (leftToRight == target[0] || leftToRight == target[1]) && (rightToLeft == target[0] || rightToLeft == target[1]){
					output += 1
				}
			}
		}
	}

	return output
}
