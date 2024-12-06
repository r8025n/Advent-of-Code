package main

import(
	"strings"
	"strconv"
)

func solvePart1(orderingRules, updates []string) int{
	output := 0
	middlePages := []string{}
	orderingRuleMap := make(map[string][]string)
	var updateMap map[string]bool

	for _, item := range orderingRules {
		parts := strings.Split(item, "|")
		before, after := parts[0], parts[1]

		if _, exists := orderingRuleMap[before]; exists {
			orderingRuleMap[before] = append(orderingRuleMap[before], after)
		}else{
			orderingRuleMap[before] = append([]string{}, after)
		}
	}

	continueToOuterLoop:
	for _, update := range updates {
		updateMap = make(map[string]bool)
		pages := strings.Split(update, ",")

		for _, page := range pages {
			if _, orderRuleExists := orderingRuleMap[page]; orderRuleExists {
				for _, item := range orderingRuleMap[page] {
					if _, pageUpdated := updateMap[item]; pageUpdated {
						continue continueToOuterLoop;
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
