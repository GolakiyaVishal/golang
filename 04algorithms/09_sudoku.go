package main

import "fmt"

var N int = 9

func main() {
	board := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	p := solveSudoku(board, 0, 0)
	if p {
		for i := 0; i < N; i++ {
			fmt.Println(board[i])
		}
	} else {
		fmt.Println("not done")
	}
}

func solveSudoku(board [][]int, row int, col int) bool {
	// if all place visited then return
	if row == N-1 && col == N {
		return true
	}

	// reach to end of row then go to next
	if col >= N {
		row++
		col = 0
	}

	//fmt.Printf("%d-%d = %d\n", row, col, board[row][col])
	// move to next column if spot is non zero
	if board[row][col] != 0 {
		return solveSudoku(board, row, col+1)
	}

	// put 1 to 9 at each place
	for i := 1; i <= N; i++ {
		if canPlace(board, row, col, i) {
			board[row][col] = i
			if solveSudoku(board, row, col+1) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

func canPlace(board [][]int, row int, col int, n int) bool {
	//check for row and column
	for c := 0; c < N; c++ {
		if board[row][c] == n || board[c][col] == n {
			return false
		}
	}

	// check for specific 3x3
	tr := row - row%3
	tc := col - col%3

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r+tr][c+tc] == n {
				return false
			}
		}
	}

	return true
}
