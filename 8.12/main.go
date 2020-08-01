package main

import (
	"fmt"
	"strconv"
)

func printNQueens(n int) {
	// Can improve space by holding a []int instead of a [][]bool
	// The single slice represents the rows and the int represents the col of the queen in that row
	board := make([][]bool, n)
	for row := range board {
		board[row] = make([]bool, n)
	}
	printQueens(board, 0)
}

func printQueens(board [][]bool, col int) {
	if col == len(board[0]) {
		printPositions(board)
	}
	for row := range board {
		if !hasConflict(board, row, col) {
			board[row][col] = true
			printQueens(board, col+1)
			board[row][col] = false
		}
	}
}

// Can improve performance if I track only the placed queens positions rather than always checking the entire board
func hasConflict(board [][]bool, row int, col int) bool {
	// Check row conflict
	for pos := range board[row] {
		if board[row][pos] {
			return true
		}
	}
	// Check col conflict
	for pos := range board {
		if board[pos][col] {
			return true
		}
	}
	// Check all diagonals
	step := 1
	for {
		moreSteps := false
		// left up
		if row-step >= 0 && col-step >= 0 {
			moreSteps = true
			if board[row-step][col-step] {
				return true
			}
		}
		// right up
		if row-step >= 0 && col+step < len(board[row]) {
			moreSteps = true
			if board[row-step][col+step] {
				return true
			}
		}
		// down left
		if row+step < len(board) && col-step >= 0 {
			moreSteps = true
			if board[row+step][col-step] {
				return true
			}
		}
		// down right
		if row+step < len(board) && col+step < len(board[row]) {
			moreSteps = true
			if board[row+step][col+step] {
				return true
			}
		}
		// Missed both without IDE
		step++
		if !moreSteps {
			break
		}
	}
	return false
}

func printPositions(board [][]bool) {
	// print the n queens positions
	pos := make([]string, 0, len(board))
	for row := range board {
		for col := range board[row] {
			if board[row][col] {
				pos = append(pos, "("+strconv.Itoa(row)+", "+strconv.Itoa(col)+")")
				break
			}
		}
	}
	fmt.Println(pos)
}
