package main

import(
	"strings"
	"strconv"
)

func solvePart1(updates []string, orderingRuleMap map[string][]string) int{
	output := 0
	middlePages := []string{}
	var updateMap map[string]bool

	outerLoop:
	for _, update := range updates {
		updateMap = make(map[string]bool)
		pages := strings.Split(update, ",")

		for _, page := range pages {
			if _, orderRuleExists := orderingRuleMap[page]; orderRuleExists {
				for _, item := range orderingRuleMap[page] {
					if _, pageUpdated := updateMap[item]; pageUpdated {
						continue outerLoop;
					}
				}
			}
			updateMap[page] = true
		}

		middlePages = append(middlePages, pages[(len(pages) + 1)/2 - 1])
	}

	for _, page := range middlePages {
		pageVal, _ := strconv.Atoi(page)
		output += pageVal
	}

	return output
}
