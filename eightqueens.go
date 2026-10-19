package piscine

import "github.com/01-edu/z01"

func isSafe(board []int, col int, row int) bool {
	for i := 0; i < col; i++ {
		horizontal := i - col
		vertical := board[i] - row

		if horizontal < 0 {
			horizontal = -horizontal
		}

		if vertical < 0 {
			vertical = -vertical
		}

		if board[i] == row || horizontal == vertical {
			return false
		}
	}
	return true
}

func solve(board []int, col int) {
	if col == 8 {
		for i := 0; i < len(board); i++ {
			z01.PrintRune(rune(board[i]) + '0')
		}
		z01.PrintRune('\n')
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func EightQueens() {
	board := [8]int{}
	solve(board[:], 0)
}
