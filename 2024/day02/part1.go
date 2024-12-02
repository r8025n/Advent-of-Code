package main

import(
	"fmt"
)

func solvePart1(reports [][]int) int{
	safe := 0
	// for report, _ := range(reports){
	// 	fmt.Println(report)
	// }
	outer:
	for _, report := range(reports) {
		fmt.Println(report)
		if report[0] < report[len(report) - 1] {
			for i := 1; i < len(report); i++ {
				diff := report[i] - report[i - 1]
				if report[i] < report[i - 1] || diff > 3 || diff < 1 {
					continue outer
				}
			}
			safe += 1
		} else {
			for i := 1; i < len(report); i++ {
				diff := report[i - 1] - report[i]
				if report[i] > report[i - 1] || diff > 3 || diff < 1 {
					continue outer
				}
			}
			safe += 1
		}
	}

	return safe
}
