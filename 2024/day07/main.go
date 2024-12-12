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

	// input := []string{}
	testValueSet := []int{}
	equationValueSet := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		splittedValues := strings.Split(line, ": ")
		remainingValues :=  strings.Fields(splittedValues[1])
		equationValues := []int{}
		for _, item := range remainingValues {
			itemValue, _ := strconv.Atoi(item)
			equationValues = append(equationValues, itemValue)
		}
		equationValueSet = append(equationValueSet, equationValues)
		testValue, _ := strconv.Atoi(splittedValues[0])
		testValueSet = append(testValueSet, testValue)
	}

	outputPart1 := solvePart1(testValueSet, equationValueSet)
	// outputPart2 := solvePart2(input2)
	fmt.Println("Solution to part1: ", outputPart1)
	// fmt.Println("Solution to part2: ", outputPart2)
}
