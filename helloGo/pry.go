package main

import (
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

	//wg.Add(3)

	nextNumber := func(in int) func(string) int {
		i := 1
		return func(s string) int { // 返回的是一个 func(string) int 的 func
			i <<= in
			fmt.Printf("in: %v -> arg: %s , ret: %v \n", in, s, i)
			defer wg.Done() // 保证最后执行
			return i * in
		}
	}(2) // 最后这个 (10) 进行了func call，得到一个 func(string) int，下面再对 nextNumber call ( nextNumber("str") ) 得到结果

	wg.Add(1)
	go fmt.Printf("for first, ret -> %v \n", nextNumber("first"))
	wg.Add(1)
	go fmt.Printf("for second, ret -> %v \n", nextNumber("second"))
	wg.Add(1)
	go fmt.Printf("for third, ret -> %v \n", nextNumber("third"))

	runtime.Gosched() // 让出时间片
	wg.Wait()
	//time.Sleep(time.Microsecond) // output同样需要时间

	// 创建一个有缓冲的 Channel，缓冲区大小为 2
	ch := make(chan int, 2)

	// 启动一个并发的 Goroutine 发送数据到 Channel
	go func() {
		for i := 0; i < 8; i++ {
			fmt.Println("Sending -> channel:", i)
			ch <- 1 << i // 发送数据到 Channel
		}
		close(ch) // 关闭 Channel
	}()

	// 主 Goroutine 从 Channel 接收数据
	for num := range ch {
		fmt.Println("Receiving <- channel:", num)
	}

	fmt.Println("Done")
}

// TP 泛型
func TP[T any](a T) {
	println(a)
}

// getSequence方法，返回一个 返回类型为int的func
func getSequence(_ bool) func(string) int {
	i := 1
	return func(str string) int {
		i <<= 2
		fmt.Printf("arg: %s , ret: %v \n", str, i)
		defer wg.Done() // 保证最后执行
		return i
	}
}
