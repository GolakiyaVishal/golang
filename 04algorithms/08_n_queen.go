package main

import "fmt"

func main() {
	board := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	result := FindSolution(board, 0)
	if result {
		for i := 0; i < count; i++ {
			fmt.Println(board[i])
		}
	} else {
		fmt.Print("No result found.")
	}

}

const count int = 5

func FindSolution(board [][]int, col int) bool {

	if col >= count {
		return true
	}

	for i := 0; i < count; i++ {
		if isSafe(board, i, col) {
			board[i][col] = 1

			if FindSolution(board, col+1) {
				return true
			}

			board[i][col] = 0
		}
	}

	return false
}

func isSafe(board [][]int, row int, col int) bool {
	// check for left side
	for i := 0; i < col; i++ {
		if board[row][i] == 1 {
			return false
		}
	}

	// check for upper left side
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// check for down left side
	for i, j := row, col; i >= 0 && j < count; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}
