package main

func solvePart1(input [][]rune) int{
	output := 0

	guardX, guardY := findGuardPosition(input)
	currentDirection := input[guardX][guardY]
	i, j := guardX, guardY

	for {
		if hasLeftMappedArea(i, j, len(input), len(input[0])){
			break
		}
		if  input[i][j] == '#' {
			i -= MoveMap[currentDirection].Forward[0]
			j -= MoveMap[currentDirection].Forward[1]
		 	currentDirection = MoveMap[currentDirection].NintyDegree
			i += MoveMap[currentDirection].Forward[0]
			j += MoveMap[currentDirection].Forward[1]
			continue
		}

		if input[i][j] != '*' {
			output += 1
		}

		input[i][j] = '*'
		i += MoveMap[currentDirection].Forward[0]
		j += MoveMap[currentDirection].Forward[1]
	}

	return output
}
