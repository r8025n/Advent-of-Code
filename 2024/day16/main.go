package main

import(
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line,"\n")

		input = append(input, line)
	}

	outputPart1 := solvePart1(input)
	// outputPart2 := solvePart2(input)
	fmt.Println("Solution to part1: ", outputPart1)
	// fmt.Println("Solution to part2: ", outputPart2)
}
