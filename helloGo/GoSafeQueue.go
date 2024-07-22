package main

import (
	"fmt"
	"sync"
)

type SafeQueue struct {
	queue []int
	lock  sync.Mutex
}

func (sq *SafeQueue) Enqueue(v int) {
	sq.lock.Lock()
	defer sq.lock.Unlock() // 会在之后执行
	sq.queue = append(sq.queue, v)
}

func (sq *SafeQueue) Dequeue() (int, bool) {
	sq.lock.Lock()
	defer sq.lock.Unlock()
	if len(sq.queue) == 0 {
		return 0, false
	}
	element := sq.queue[0]
	sq.queue = sq.queue[1:]
	return element, true
}

func main() {
	var queue SafeQueue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue()) // 1, true
	fmt.Println(queue.Dequeue()) // 2, true
	fmt.Println(queue.Dequeue()) // 3, true
	fmt.Println(queue.Dequeue()) // 0, false (queue is empty)
}
