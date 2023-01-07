// Author 		: Praveen
// Date   		: 07/01/2023
// Question 	: https://leetcode.com/problems/available-captures-for-rook/
// Submission 	: https://leetcode.com/problems/available-captures-for-rook/submissions/873246059/

/*
	Question:
	--------------------------------
	On an 8 x 8 chessboard, there is exactly one white rook 'R' and some number of white bishops 'B', black pawns 'p', and empty squares '.'.

	When the rook moves, it chooses one of four cardinal directions (north, east, south, or west), then moves in that direction until it chooses to stop, reaches the edge of the board, captures a black pawn, or is blocked by a white bishop. A rook is considered attacking a pawn if the rook can capture the pawn on the rook's turn. The number of available captures for the white rook is the number of pawns that the rook is attacking.

	Return the number of available captures for the white rook.


	Example 1:
	Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
	Output: 3


*/

package main

import "fmt"

func main() {
	board := [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'B', 'R', 'B', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(numRookCaptures(board))
}

func numRookCaptures(board [][]byte) int {
	var count int

	rowInd := 0
	colInd := 0

outer:
	for i, row := range board {
		for j, col := range row {
			if col == 'R' {
				rowInd = i
				colInd = j
				break outer
			}
		}
	}

	l := 8

	// Left
	for i := colInd; i >= 0; i-- {
		if board[rowInd][i] == 'B' {
			break
		}
		if board[rowInd][i] == 'p' {
			count++
			break
		}
	}

	// Right
	for i := colInd; i < l; i++ {
		if board[rowInd][i] == 'B' {
			break
		}
		if board[rowInd][i] == 'p' {
			count++
			break
		}
	}

	// Up
	for i := rowInd; i >= 0; i-- {
		if board[i][colInd] == 'B' {
			break
		}
		if board[i][colInd] == 'p' {
			count++
			break
		}
	}

	// Down
	for i := rowInd; i < l; i++ {
		if board[i][colInd] == 'B' {
			break
		}
		if board[i][colInd] == 'p' {
			count++
			break
		}
	}

	return count
}
