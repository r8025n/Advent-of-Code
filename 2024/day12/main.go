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

	input1 := [][]rune{}
	input2 := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		input1 = append(input1, []rune(line))
		input2 = append(input2, []rune(line))
	}

	outputPart1 := solvePart1(input1)
	outputPart2 := solvePart2(input2)
	fmt.Println("Solution to part1: ", outputPart1)
	fmt.Println("Solution to part2: ", outputPart2)
}
