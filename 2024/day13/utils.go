package main

import(
	"regexp"
	"strconv"
)

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

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
