package main

import(
	"unicode"
)

func calculateAntiNodesForPart2(nodeA, nodeB []int, X, Y int) [][]int {
	antiNodes := [][]int{}
	xDiff := nodeB[0] - nodeA[0]
	yDiff := nodeB[1] - nodeA[1]

	for x, y, i := nodeA[0], nodeA[1], 0; isAntiNodeValid([]int{x, y}, X, Y); i++ {
		antiNodes = append(antiNodes, []int{x, y})
		x = nodeA[0] + xDiff * i
		y = nodeA[1] + yDiff * i
	}

	for x, y, i := nodeA[0], nodeA[1], 0; isAntiNodeValid([]int{x, y}, X, Y); i++ {
		antiNodes = append(antiNodes, []int{x, y})
		x = nodeA[0] - xDiff * i
		y = nodeA[1] - yDiff * i
	}

	return antiNodes
}

func solvePart2(input []string) int{
	output := 0
	antennaMap := make(map[string][][]int)
	antiNodeMap := make([][]int, len(input))
	for i := range antiNodeMap {
		antiNodeMap[i] = make([]int, len(input[0]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if !unicode.IsLetter(rune(input[i][j])) && !unicode.IsDigit(rune(input[i][j])) {
				continue
			}

	  		if _, exists := antennaMap[string(input[i][j])]; exists {
				antennaMap[string(input[i][j])] = append(antennaMap[string(input[i][j])], []int{i, j})
			}else{
				antennaMap[string(input[i][j])] = append([][]int{}, []int{i, j})
			}
		}
	}

	for _, value := range antennaMap {
		for i := 0; i < len(value); i++ {
			for j := i + 1; j < len(value); j++ {
				antiNodes := calculateAntiNodesForPart2(value[i], value[j], len(input), len(input[0]))
				for _, antiNode := range antiNodes {
					if isAntiNodeUnique(antiNode, antiNodeMap) {
						antiNodeMap[antiNode[0]][antiNode[1]] = 1
						output += 1
					}
				}
			}
		}
	}

	return output
}
