package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/trees/redblacktree"
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

	tree := redblacktree.NewWithIntComparator() // 创建一个红黑树
	tree.Put(1, "a")                            // 插入元素
	tree.Put(2, "b")
	tree.Put(3, "c")

	// 查找
	value, found := tree.Get(2) // 通过 key 查找值
	if found {
		fmt.Println("Found:", value)
	}

	// 删除
	tree.Remove(2) // 删除元素
}

type iterator struct {
	redblacktree.Iterator
	count int
}
