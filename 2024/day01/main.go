package main

import(
	"os"
	"fmt"
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

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	outputPart1 := solvePart1(lines)
	fmt.Println("Solution to part1: ", outputPart1)
}
