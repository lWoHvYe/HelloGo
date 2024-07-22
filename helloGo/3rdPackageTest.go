package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/stacks/arraystack"
)

func main() {
	// 使用GoDS的栈
	stack := arraystack.New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	value, _ := stack.Pop()
	fmt.Println(value) // 3

	// 使用GoDS的队列
	queue := linkedlistqueue.New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	value, _ = queue.Dequeue()
	fmt.Println(value) // 1
}
