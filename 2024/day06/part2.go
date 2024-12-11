package main

// import(
// 	"fmt"
// )

func solvePart2(input [][]rune) int{
	loopCreatingObstacle := 0
	guardX, guardY := findGuardPosition(input)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			visitedDirectionsMap := generateVisitedDirectionsMap(len(input), len(input[0]))
			currentDirection := input[guardX][guardY]
			if input[i][j] == '#' || input[i][j] == currentDirection {
				continue
			}

			input[i][j] = '#'
			x, y := guardX, guardY

			for {
				if hasLeftMappedArea(x, y, len(input), len(input[0])){
					break
				}
				if  input[x][y] == '#' {
					if visitedDirectionsMap[x][y][currentDirection] {
						loopCreatingObstacle += 1
						break
					}

					visitedDirectionsMap[x][y][currentDirection] = true

					x -= MoveMap[currentDirection].Forward[0]
					y -= MoveMap[currentDirection].Forward[1]
				 	currentDirection = MoveMap[currentDirection].NintyDegree
					x += MoveMap[currentDirection].Forward[0]
					y += MoveMap[currentDirection].Forward[1]
					continue
				}

				x += MoveMap[currentDirection].Forward[0]
				y += MoveMap[currentDirection].Forward[1]
			}
			input[i][j] = '.'
		}
	}

	return loopCreatingObstacle
}
