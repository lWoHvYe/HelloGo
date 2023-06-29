package main

import "math"

func main() {
	println(fibonacci(2))
	println(fibonacci2(4))
	println(fibonacciByFormula(6))
}

// 使用矩阵乘法的方法来计算斐波那契数列，以减少计算步骤的方式。这种方法基于斐波那契数列的递推关系和矩阵乘法的性质。
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	// 定义斐波那契数列的初始矩阵
	matrix := [][]int{{1, 1}, {1, 0}}

	// 计算矩阵的 n 次方
	result := matrixPower(matrix, n-1)

	// 返回斐波那契数列的第 n 个数
	return result[0][0]
}

// 矩阵的 n 次方
func matrixPower(matrix [][]int, n int) [][]int {
	if n == 0 {
		// 单位矩阵
		return [][]int{{1, 0}, {0, 1}}
	}

	if n == 1 {
		return matrix
	}

	result := matrixPower(matrix, n/2)

	// x^2
	result = matrixMultiply(result, result)

	// 奇数
	if n%2 == 1 {
		result = matrixMultiply(result, matrix)
	}

	return result
}

// 矩阵乘法
func matrixMultiply(a, b [][]int) [][]int {
	rowsA := len(a)    // 矩阵 a 的行数
	colsA := len(a[0]) // 矩阵 a 的列数
	colsB := len(b[0]) // 矩阵 b 的列数

	// 创建结果矩阵，行数为矩阵 a 的行数，列数为矩阵 b 的列数
	result := make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		result[i] = make([]int, colsB)
	}

	// 逐个计算结果矩阵的每个元素，这个可以挤一挤
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j] // 矩阵元素的乘法累加，A逐行，一行中每个元素 * B一列中每个元素
			}
		}
	}

	return result
}

// 迭代法
func fibonacci2(n int) int {
	if n <= 1 {
		return n
	}

	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

// 斐波那契数列的通项公式如下：
//
// Fn = (1/sqrt(5)) * (((1+sqrt(5))/2)^n - ((1-sqrt(5))/2)^n)
//
// 其中，Fn 表示第 n 个斐波那契数。
// 请注意，由于使用浮点数进行计算，当需要计算大型斐波那契数时，可能会出现精度问题。此方法适用于小型斐波那契数的计算。
// 如果需要计算大型斐波那契数，请考虑其他更适合的算法或数据结构。
func fibonacciByFormula(n int) int {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2

	fn := (math.Pow(phi, float64(n)) - math.Pow(1-phi, float64(n))) / sqrt5

	return int(fn)
}
