package main

import(
	"strings"
	"strconv"
	"fmt"
)

func solvePart2(updates []string, orderingRuleMap map[string][]string) int{
	output := 0
	var updateMap map[string]int

	for _, item := range updates {
		updateMap = make(map[string]int)
		update := strings.Split(item, ",")
		isIncorrect := false

		for i := 0; i < len(update); i++ {
			updateMap[update[i]] = i
		}

		// custom ordering rule based bubble sort
		for j := 0; j < len(update); j++ {
			for k := 0; k < len(update); k++ {
				currentPage := update[k]
				afterPages := orderingRuleMap[currentPage]

				for _, afterPage := range afterPages {
					if _, exists := updateMap[afterPage]; exists {
						if updateMap[afterPage] < k {
							isIncorrect = true
							update[updateMap[afterPage]], update[k] = currentPage, afterPage
							updateMap[afterPage], updateMap[currentPage] = k,  updateMap[afterPage]
							break
						}
					}
				}
			}
		}

		if isIncorrect {
			middlePage := update[(len(update) + 1)/2 - 1]
			middlePageVal, _ := strconv.Atoi(middlePage)
			output += middlePageVal
		}
	}

	return output
}
