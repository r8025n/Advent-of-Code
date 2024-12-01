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

	var list1, list2 []int
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		line = strings.Trim(line, " ")
		locIds := strings.Fields(line)
		locId1, _ := strconv.Atoi(locIds[0])
		locId2, _ := strconv.Atoi(locIds[1])
		list1 = append(list1, locId1)
		list2 = append(list2, locId2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	outputPart1 := solvePart1(list1, list2)
	outputPart2 := solvePart2(list1, list2)
	fmt.Println("Solution to part1: ", outputPart1)
	fmt.Println("Solution to part2: ", outputPart2)
}
