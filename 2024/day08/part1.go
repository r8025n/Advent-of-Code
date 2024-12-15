package main

import(
	// "fmt"
	"unicode"
	"math"
)

func calculateAntiNodesForPart1(nodeA, nodeB []int) [][]int {
	xDiff := int(math.Abs(float64(nodeA[0] - nodeB[0])))
	yDiff := int(math.Abs(float64(nodeA[1] - nodeB[1])))

	if xDiff == 0 {
		return [][]int{[]int{nodeA[0], nodeA[1] - yDiff}, []int{nodeB[0], nodeB[1] + yDiff}}
	}else if yDiff == 0 {
		return [][]int{[]int{nodeA[0] - xDiff, nodeA[1]}, []int{nodeB[0] + xDiff, nodeB[1]}}
	}else if nodeA[1] > nodeB[1] {
		return [][]int{[]int{nodeA[0] - xDiff, nodeA[1] + yDiff}, []int{nodeB[0] + xDiff, nodeB[1] - yDiff}}
	}else{
		return [][]int{[]int{nodeA[0] - xDiff, nodeA[1] - yDiff}, []int{nodeB[0] + xDiff, nodeB[1] + yDiff}}
	}
}

// func isAntiNodeValidAndUnique(node []int, X, Y int, antiNodeMap[][]int) bool {
// 	if node[0] < 0 || node[1] < 0 || node[0] >= X || node[1] >= Y {
// 		return false
// 	}

// 	if antiNodeMap[node[0]][node[1]] == 1 {
// 		return false
// 	}

// 	return true
// }

func solvePart1(input []string) int{
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
				antiNodes := calculateAntiNodesForPart1(value[i], value[j])

				for _, antiNode := range antiNodes {
					if isAntiNodeValid(antiNode, len(input), len(input[0])) && isAntiNodeUnique(antiNode, antiNodeMap)  {
						antiNodeMap[antiNode[0]][antiNode[1]] = 1
						output += 1
					}
				}
			}
		}
	}

	return output
}
