package main

import (
	"container/heap"
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("hello world . 你好 世界！")

	str := "lWoHvYe"

	TP("WHY")
	TP(str)
	TP(1)
	TP(true)
	TP(1.2)
	TP('y')

	fmt.Println("...go...")

	wg.Add(3)

	nextNumber := getSequence()
	go fmt.Println(nextNumber())
	go fmt.Println(nextNumber())
	go fmt.Println(nextNumber())

	runtime.Gosched() // 让出时间片
	wg.Wait()
	//time.Sleep(100 * time.Microsecond)

	list := []int{49, 38, 65, 97, 76, 13, 27, 49, 55, 04}
	//fmt.Println(ShellSort(list))
	//fmt.Println(ShellSort2(list))
	//fmt.Println(list)
	//QuickSort(list, 0, len(list)-1)
	//fmt.Println(list)

	h := &MinHeap{2, 1, 5, 3, 4}
	heap.Init(h)
	fmt.Println(h)
	heap.Push(h, 0)
	fmt.Println(h)
	fmt.Println(heap.Pop(h))

	fmt.Println(MergeSort(list))
}

// TP 泛型
func TP[T any](a T) {
	println(a)
}

// getSequence方法，返回一个 返回类型为int的func
func getSequence() func() int {
	i := 1
	return func() int {
		i <<= 2
		defer wg.Done() // 保证最后执行
		return i
	}
}

// ShellSort 希尔排序
func ShellSort(nums []int) []int {
	n := len(nums)
	gap := n / 2

	for gap > 0 {
		for k := 0; k < gap; k++ {
			for i := k + gap; i < n; i += gap {
				temp := nums[i]
				j := i
				for j >= gap && nums[j-gap] > temp {
					nums[j] = nums[j-gap]
					j -= gap
				}
				nums[j] = temp
			}
		}
		gap /= 2
	}

	return nums
}

func ShellSort2(nums []int) []int {
	n := len(nums)
	gap := 1

	// 计算初始步长
	for gap < n/3 {
		gap = 3*gap + 1
	}

	// 外层循环控制步长
	for gap > 0 {
		// 内层循环进行插入排序
		for i := gap; i < n; i++ {
			temp := nums[i]
			j := i - gap

			for j >= 0 && nums[j] > temp {
				nums[j+gap] = nums[j]
				j -= gap
			}

			nums[j+gap] = temp
		}

		// 更新步长
		gap = gap / 3
	}

	return nums
}

func QuickSort(nums []int, low, high int) {
	if low < high {
		pivot := MedianOfThree(nums, low, high) // 使用三数取中法选择枢轴
		pivots := Partition2(nums, low, high, pivot)
		QuickSort(nums, low, pivots-1)
		QuickSort(nums, pivots+1, high)
	}
}

func MedianOfThree(nums []int, low, high int) int {
	mid := low + (high-low)/2

	// 通过比较和交换确保 nums[low] <= nums[mid] <= nums[high]
	if nums[low] > nums[mid] {
		nums[low], nums[mid] = nums[mid], nums[low]
	}
	if nums[mid] > nums[high] {
		nums[mid], nums[high] = nums[high], nums[mid]
	}
	if nums[low] > nums[mid] {
		nums[low], nums[mid] = nums[mid], nums[low]
	}

	return mid // 返回中间元素的index
}

func Partition(nums []int, low, high, pivot int) int {
	pivotNum := nums[pivot]
	nums[pivot], nums[low] = nums[low], nums[pivot]

	for low < high {
		for low < high && nums[high] >= pivotNum {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= pivotNum {
			low++
		}
		nums[high] = nums[low]
	}

	nums[low] = pivotNum
	return low
}

func Partition2(nums []int, low, high, pivot int) int {
	// 将枢轴移到最右边
	nums[high], nums[pivot] = nums[pivot], nums[high]

	for j := low; j < high; j++ {
		if nums[j] <= nums[high] {
			nums[low], nums[j] = nums[j], nums[low]
			low++
		}
	}

	// 将枢轴放回正确的位置
	nums[low], nums[high] = nums[high], nums[low]

	return low // 返回枢轴的位置
}

// MinHeap 自定义类型
type MinHeap []int

// Len 实现 heap.Interface 接口的方法
func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // 末尾元素
	*h = old[0 : n-1] // 不含尾
	return x
}

// MergeSort 2路归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) >> 1
	left := MergeSort(arr[:mid]) // 切片
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// 合并两个有序数组
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	// 依次比较左右数组的元素，将较小的元素放入结果数组
	for ; i < len(left) && j < len(right); k++ {
		if left[i] <= right[j] { // 稳定性
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	// 将剩余的元素添加到结果数组
	for ; i < len(left); k++ {
		result[k] = left[i]
		i++
	}

	for ; j < len(right); k++ {
		result[k] = right[j]
		j++
	}

	return result
}
