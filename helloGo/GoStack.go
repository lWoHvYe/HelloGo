package main

import "fmt"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func main() {
	var stack Stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop()) // 3, true
	fmt.Println(stack.Pop()) // 2, true
	fmt.Println(stack.Pop()) // 1, true
	fmt.Println(stack.Pop()) // 0, false (stack is empty)
}
