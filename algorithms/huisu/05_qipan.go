package main

import "strings"

/*
	按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

	n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

	给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

	每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

*/
// solveNQueens 51. N 皇后
func solveNQueens(n int) [][]string {
	var res [][]string
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		temp := make([]string,n)
		for j := 0;j < n;j++{
			temp[j] = "."
		}
		chessboard[i] = temp
	}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				backtrack(row + 1)
				chessboard[row][i] = "."
			}
		}
	}
	backtrack(0)
	return res
}

func isValid(n, row, col int, chessboard [][]string) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}




