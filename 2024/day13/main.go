package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"regexp"
)

type PrizeData struct {
	A []int
	B []int
	Target []int
}

func extractIntegerData(str string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str, -1)

	data := []int{}
	for _, item := range matches {
		num, _ := strconv.Atoi(item)
		data = append(data, num)
	}

	return data
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	stringifiedData := string(data)
	splittedData := strings.Split(stringifiedData, "\n\n")
	var input []PrizeData

	for _, item := range splittedData {
		lines := strings.Split(item, "\n")
		var prizeData PrizeData

		prizeData.A = extractIntegerData(lines[0])
		prizeData.B = extractIntegerData(lines[1])
		prizeData.Target = extractIntegerData(lines[2])

		input = append(input, prizeData)
	}

	outputPart1 := solvePart1(input)
	// outputPart2 := solvePart2(input)
	fmt.Println("Solution to part1: ", outputPart1)
	// fmt.Println("Solution to part2: ", outputPart2)
}
