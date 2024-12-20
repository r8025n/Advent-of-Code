package main

import(
	"fmt"
	"os"
	"strings"
	"bufio"
	// "strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		input = append(input, []rune(line))
	}

	outputPart1 := solvePart1(input)
	// outputPart2 := solvePart2(input[0])
	fmt.Println("Solution to part1: ", outputPart1)
	// fmt.Println("Solution to part2: ", outputPart2)
}
