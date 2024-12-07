package main

import(
	// "os"
	"fmt"
	// "bufio"
	"strings"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

	input := string(data)
	inputs := strings.Split(input, "\n\n")

	orderingRules := strings.Split(inputs[0], "\n")
	updateStrings := strings.Split(inputs[1], "\n")

	orderingRuleMap := make(map[string][]string)
	for _, item := range orderingRules {
		parts := strings.Split(item, "|")
		before, after := parts[0], parts[1]

		if _, exists := orderingRuleMap[before]; exists {
			orderingRuleMap[before] = append(orderingRuleMap[before], after)
		}else{
			orderingRuleMap[before] = append([]string{}, after)
		}
	}

	outputPart1 := solvePart1(updateStrings, orderingRuleMap)
	outputPart2 := solvePart2(updateStrings, orderingRuleMap)
	fmt.Println("Solution to part1: ", outputPart1)
	fmt.Println("Solution to part2: ", outputPart2)
}
