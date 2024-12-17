package main

import(
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		inputLine := []int{}

		for _, c := range line{
			inputVal, _ := strconv.Atoi(string(c))
			inputLine = append(inputLine, inputVal)
		}

		input = append(input, inputLine)
	}

	outputPart1 := solvePart1(input)
	outputPart2 := solvePart2(input)
	fmt.Println("Solution to part1: ", outputPart1)
	fmt.Println("Solution to part2: ", outputPart2)
}
