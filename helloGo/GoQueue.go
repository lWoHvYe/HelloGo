package main

import "fmt"

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func main() {
	var queue Queue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue()) // 1, true
	fmt.Println(queue.Dequeue()) // 2, true
	fmt.Println(queue.Dequeue()) // 3, true
	fmt.Println(queue.Dequeue()) // 0, false (queue is empty)
}
