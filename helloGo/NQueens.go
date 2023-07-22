package main

import (
	"fmt"
	"math/bits"
)

// 回溯

var solutions [][]string

func solveNQueens(boardSize int) [][]string {
	solutions = [][]string{}
	queens := make([]int, boardSize)
	for i := 0; i < boardSize; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}                                // 列
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{} // 两对角，分别满足差相等、和相等
	backtrack(queens, boardSize, 0, columns, diagonals1, diagonals2)
	return solutions
}

// 干活
func backtrack(queens []int, boardSize, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == boardSize {
		board := generateBoard(queens, boardSize)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < boardSize; i++ { // i 是列，row是行。每行放一个，有的是按列放的，注意区分
		if columns[i] { // 列冲突
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] { // 左上 -> 右下 冲突
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] { // 左下 -> 右上 冲突
			continue
		}
		// 至此，可以放
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack(queens, boardSize, row+1, columns, diagonals1, diagonals2) // 开始放下一个皇后
		// After the recursive call returns, the code backtracks by removing the queen from the current row and resetting the maps to their previous state
		//  (removing the queen from the column and diagonals) to explore other possible placements.
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	var board []string
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

// 基于位运算回溯
func solveNQueens2(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(queens, n, 0, 0, 0, 0)
	return solutions
}

// x & (-x) 可以获得x的二进制表示中的最低位的1的位置
// x & (x - 1) 可以将x的二进制表示中的最低位的1置成0
func solve(queens []int, boardSize, row, columns, diagonals1, diagonals2 int) {
	if row == boardSize {
		board := generateBoard(queens, boardSize)
		solutions = append(solutions, board)
		return
	}
	// 得到可放置皇后的位置
	availablePositions := ((1 << boardSize) - 1) & (^(columns | diagonals1 | diagonals2))
	for availablePositions != 0 {
		position := availablePositions & (-availablePositions)             // 最低位1
		availablePositions = availablePositions & (availablePositions - 1) // 将最低位1置0，表示放了旗子
		column := bits.OnesCount(uint(position - 1))                       // 1的个数
		queens[row] = column
		// 进入下一行时， columns 的值保持不变，diagonals1 左移一位，diagonals2 右移一位
		solve(queens, boardSize, row+1, columns|position, (diagonals1|position)>>1, (diagonals2|position)<<1)
		queens[row] = -1
	}
}

func main() {
	res := solveNQueens(8)
	fmt.Println(len(res))
	res = solveNQueens2(8)
	fmt.Println(len(res))
}
