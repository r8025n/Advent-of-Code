package main

import(
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

	splittedData := strings.Split(strings.Trim(string(data), "\n"), "\n")
	input1 := make([][]rune, len(splittedData))
	input2 := make([][]rune, len(splittedData))

	for i, item := range splittedData {
		input1[i] = []rune(item)
		input2[i] = []rune(item)
	}

	outputPart1 := solvePart1(input1)
	outputPart2 := solvePart2(input2)
	fmt.Println("Solution to part1: ", outputPart1)
	fmt.Println("Solution to part2: ", outputPart2)
}
