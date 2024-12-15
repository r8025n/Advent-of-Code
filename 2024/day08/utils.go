package main

func isAntiNodeValid(node []int, X, Y int) bool {
	if node[0] >= 0 && node[1] >= 0 && node[0] < X && node[1] < Y {
		return true
	}

	return false
}

func isAntiNodeUnique(node []int, antiNodeMap[][]int) bool {
	if antiNodeMap[node[0]][node[1]] == 1 {
		return false
	}

	return true
}
