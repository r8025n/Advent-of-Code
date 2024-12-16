package main

import(
	"strconv"
)

func solvePart2(input string) int{
	output := 0
	inputs := []rune(input)
	expandedInput := []int{}
	id := 0
	emptyBlockIdVsBlockSize := [][]int{}
	idVsBlockSize := [][]int{}
	expandedInputIndex := 0
	for i := 0; i < len(inputs); i++ {
		blockSize, _ := strconv.Atoi(string(inputs[i]))
		if i % 2 == 0 {
			idVsBlockSize = append(idVsBlockSize, []int{expandedInputIndex, blockSize})
			for j := 0; j < blockSize; j++ {
				expandedInput = append(expandedInput, id)
				expandedInputIndex++
			}
			id++
		}else{
			emptyBlockIdVsBlockSize = append(emptyBlockIdVsBlockSize, []int{expandedInputIndex, blockSize})
			for j := 0; j < blockSize; j++ {
				expandedInput = append(expandedInput, -1)
				expandedInputIndex++
			}
		}
	}

	id--

	for id >= 0 {
		for i := 0; i < len(emptyBlockIdVsBlockSize); i++ {
			if idVsBlockSize[id][0] < emptyBlockIdVsBlockSize[i][0]{
				break
			}
			if idVsBlockSize[id][1] <= emptyBlockIdVsBlockSize[i][1] {
				for j := emptyBlockIdVsBlockSize[i][0]; j < emptyBlockIdVsBlockSize[i][0] + idVsBlockSize[id][1]; j++ {
					expandedInput[j] = id
				}

				for j := idVsBlockSize[id][0]; j < idVsBlockSize[id][0] + idVsBlockSize[id][1]; j++ {
					expandedInput[j] = -1
				}

				emptyBlockIdVsBlockSize[i][0] += idVsBlockSize[id][1]
				emptyBlockIdVsBlockSize[i][1] -= idVsBlockSize[id][1]
				break;
			}
		}
		id--
	}

	for i := 0; i < len(expandedInput); i++ {
		if expandedInput[i] >= 0 {
			output += expandedInput[i] * i
		}
	}

	return output
}
