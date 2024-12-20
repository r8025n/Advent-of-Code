package main

type Queue [][]int

func (q *Queue) Enqueue(xy []int) {
	*q = append(*q, xy)
}

func (q *Queue) Dequeue() []int {
	if len(*q) == 0 {
		return nil
	}

	item := (*q)[0]
	*q = (*q)[1:]

	return item
}

// func (q *Queue) Peek() []int {
// 	if len(*q) == 0 {
// 		return nil
// 	}

// 	return (*q)[0]
// }

func (q *Queue) Empty() bool {
	if len(*q) == 0 {
		return true
	}

	return false
}
