package main

type Move struct {
	Forward []int
	NintyDegree rune
}

var MoveMap = map[rune]Move{
	'^': {Forward: []int{-1, 0}, NintyDegree: '>'},
	'>': {Forward: []int{0, 1}, NintyDegree: 'v'},
	'v': {Forward: []int{1, 0}, NintyDegree: '<'},
	'<': {Forward: []int{0, -1}, NintyDegree: '^'},
}

func findGuardPosition(input [][]rune) (int, int) {
	var x, y int

	out:
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if (input[i][j] != '#' && input[i][j] != '.'){
				x, y = i, j
				break out
			}
		}
	}

	return x, y
}

func hasLeftMappedArea(x, y, X, Y int) bool {
	if x < 0 || y < 0 || x >= X || y >= Y {
		return true
	}

	return false
}

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
