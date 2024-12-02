package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
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

	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		line = strings.Trim(line, " ")
		splittedLine := strings.Fields(line)
		var report []int
		for _, lvl := range (splittedLine) {
			lvlValue, _ := strconv.Atoi(lvl)
			report = append(report, lvlValue)
		}

		reports = append(reports, report)
	}

	fmt.Println(reports)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	outputPart1 := solvePart1(reports)
	// outputPart2 := solvePart2(list1, list2)
	fmt.Println("Solution to part1: ", outputPart1)
	// fmt.Println("Solution to part2: ", outputPart2)
}
