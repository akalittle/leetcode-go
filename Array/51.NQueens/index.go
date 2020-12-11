package Array

import (
	"strings"
)

//The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
//
//Given an integer n, return all distinct solutions to the n-queens puzzle.
//
//Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
//
//
//Input: n = 4
//Output: [[".Q..", "...Q", "Q...", "..Q."], ["..Q.", "Q...", "...Q", ".Q.."]]
//Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

func solveNQueens(n int) [][]string {
	board := make([][]string, n)

	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}

	res := [][]string{}
	backtrace(0, board, &res, n, &cols, &diag1, &diag2)
	return res
}

func backtrace(r int, board [][]string, res *[][]string, n int, cols, diag1, diag2 *map[int]bool) {

	if r == n {
		temp := make([]string, len(board))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		*res = append(*res, temp)
		return
	}

	for c := 0; c < n; c++ {
		if !(*cols)[c] && !(*diag1)[r+c] && !(*diag2)[r-c] {
			board[r][c] = "Q"
			(*cols)[c] = true
			(*diag1)[r+c] = true
			(*diag2)[r-c] = true
			backtrace(r+1, board, res, n, cols, diag1, diag2)
			board[r][c] = "."
			(*cols)[c] = false
			(*diag1)[r+c] = false
			(*diag2)[r-c] = false
		}
	}
}
