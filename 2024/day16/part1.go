package main

import(
	// "fmt"
	"math"
)

var directionChangeToCostMap = map[string]int{
	"_>": 0,
	">^": 1001,
	">v": 1001,
	"^>": 1001,
	"^<": 1001,
	"v>": 1001,
	"v<": 1001,
	"<^": 1001,
	"<v": 1001,
	">>": 1,
	"<<": 1,
	"^^": 1,
	"vv": 1,
}

func dfsSolve(input []string, visited [][]int, x, y, cost int, dirChange string) {
	if _, exists := directionChangeToCostMap[dirChange]; !exists {
		// fmt.Println(dirChange)
		return
	}

	if x < 0 || y < 0 || x >= len(input) || y >= len(input[0]) {
		return
	}

	if input[x][y] == '#' {
		return
	}


	newCost := cost + directionChangeToCostMap[dirChange]
	if visited[x][y] != 0 {
		if newCost >= visited[x][y] {
			return
		}
	}
	visited[x][y] = newCost

	if input[x][y] == 'E' {
		return
	}

	// opt1 := directionChangeToCostMap[dirChange] + dfsSolve(input, visited, x + 1, y, string(dirChange[1]) + "v")
	// opt2 := directionChangeToCostMap[dirChange] + dfsSolve(input, visited, x - 1, y, string(dirChange[1]) + "^")
	// opt3 := directionChangeToCostMap[dirChange] + dfsSolve(input, visited, x, y + 1, string(dirChange[1]) + ">")
	// opt4 := directionChangeToCostMap[dirChange] + dfsSolve(input, visited, x, y - 1, string(dirChange[1]) + "<")
	//
	dfsSolve(input, visited, x + 1, y, newCost, string(dirChange[1]) + "v")
	dfsSolve(input, visited, x - 1, y, newCost, string(dirChange[1]) + "^")
	dfsSolve(input, visited, x, y + 1, newCost, string(dirChange[1]) + ">")
	dfsSolve(input, visited, x, y - 1, newCost, string(dirChange[1]) + "<")

	return
	// return min([]int{opt1, opt2, opt3, opt4})
}

func solvePart1(input []string) int {
	// output := 0
	visited := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		visited[i] = make([]int, len(input[0]))
		for j := 0; j < len(visited[i]); j++ {
			visited[i][j] = math.MaxInt
		}
	}

	// for _, line := range input {
	// 	fmt.Println(line)
	// }
	startX, startY := findTarget(input, 'S')
	finishX, finishY := findTarget(input, 'E')
	dfsSolve(input, visited, startX, startY, 0, "_>")

	// for _, row := range visited {
	// 	fmt.Println(row)
	// }

	return visited[finishX][finishY]
}
