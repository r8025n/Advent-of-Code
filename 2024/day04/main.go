package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
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
		line = strings.Trim(line, " ")

		input = append(input, line)
	}


	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	outputPart1 := solvePart1(input)
	outputPart2 := solvePart2(input)
	fmt.Println("Solution to part1: ", outputPart1)
	fmt.Println("Solution to part2: ", outputPart2)
}
