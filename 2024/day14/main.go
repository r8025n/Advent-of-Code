package main

import(
	"fmt"
	// "strings"
	"os"
	"bufio"
)

type RobotData struct{
	Position []int
	Velocity []int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []RobotData

	for scanner.Scan() {
		line := scanner.Text()
		extractedData := extractIntegerData(line)
		robotData := RobotData{
			Position: []int{extractedData[1], extractedData[0]},
			Velocity: []int{extractedData[3], extractedData[2]},
		}

		input = append(input, robotData)
	}

	outputPart1 := solvePart1(input)
	// outputPart2 := solvePart2(input)
	fmt.Println("Solution to part1: ", outputPart1)
	// fmt.Println("Solution to part2: ", outputPart2)
}
