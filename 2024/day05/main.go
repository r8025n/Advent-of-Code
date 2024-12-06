package main

import(
	// "os"
	"fmt"
	// "bufio"
	"strings"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

	input := string(data)
	inputs := strings.Split(input, "\n\n")

	orderingRules := strings.Split(inputs[0], "\n")
	updateStrings := strings.Split(inputs[1], "\n")


	outputPart1 := solvePart1(orderingRules, updateStrings)
	// outputPart2 := solvePart2(input)
	fmt.Println("Solution to part1: ", outputPart1)
	// fmt.Println("Solution to part2: ", outputPart2)
}
