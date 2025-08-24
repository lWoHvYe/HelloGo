package main

// 创建最小堆（适用于 int 类型）
func IntMinHeap() *Hp[int] {
	return &Hp[int]{
		sortSlice: []int{},
		lessFunc:  func(a, b int) bool { return a < b }, // 小顶堆
	}
}

// 创建最大堆（适用于 int 类型）
func IntMaxHeap() *Hp[int] {
	return &Hp[int]{
		sortSlice: []int{},
		lessFunc:  func(a, b int) bool { return a > b }, // 大顶堆
	}
}

// 定义泛型最小堆
type Hp[T any] struct {
	sortSlice []T
	lessFunc  func(a, b T) bool // 用于比较元素的函数
}

// 实现 heap.Interface 方法
func (h Hp[T]) Len() int           { return len(h.sortSlice) }
func (h Hp[T]) Less(i, j int) bool { return h.lessFunc(h.sortSlice[i], h.sortSlice[j]) }
func (h Hp[T]) Swap(i, j int)      { h.sortSlice[i], h.sortSlice[j] = h.sortSlice[j], h.sortSlice[i] }

// Push 方法
func (h *Hp[T]) Push(v any) {
	h.sortSlice = append(h.sortSlice, v.(T))
}

// Pop 方法
func (h *Hp[T]) Pop() any {
	n := len(h.sortSlice)
	v := h.sortSlice[n-1]
	h.sortSlice = h.sortSlice[:n-1]
	return v
}

// Peek 方法：获取堆顶元素但不移除
func (h *Hp[T]) Peek() T {
	if len(h.sortSlice) == 0 {
		var zero T
		return zero // 返回 T 类型的零值
	}
	return h.sortSlice[0]
}
