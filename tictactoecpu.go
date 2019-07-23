package main

import (
	"errors"
	"math/rand"
)

func cpuMakeMove(board tictactoeboard) (int, int) {

	row, col, err := cpuGetMoveThatStopsThreeInARow(board)
	if err == nil {
		return row, col
	}

	//Takes Middle Square if possible (may remove this later)
	if cpuCheckBoardEmpty(board) {
		return 1, 1
	}

	return 0, 0
}

func cpuCheckBoardEmpty(board tictactoeboard) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.board[i][j] != SquareEmpty {
				return false
			}
		}
	}
	return true
}

type rowcolTuple struct {
	row int
	col int
}

func cpuGetMoveThatStopsThreeInARow(board tictactoeboard) (int, int, error) {

	availableMoves := cpuGetAvailableMoves(board)

	// If the opponent were to pick this move, would they win?
	for _, move := range availableMoves {
		tempBoard := board
		tempBoard.makeMove(SquareX, move.row, move.col)
		if tempBoard.determineBoardState() == WinnerX {
			return move.row, move.col, nil
		}
	}
	return 0, 0, errors.New("No Move that stops Three in a row")
}

func cpuGetAvailableMoves(board tictactoeboard) []rowcolTuple {

	moves := []rowcolTuple{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.board[i][j] == SquareEmpty {
				moves = append(moves, rowcolTuple{i, j})
			}
		}
	}
	return moves
}

func cpuPickRandomMove(board tictactoeboard) rowcolTuple {
	moves := cpuGetAvailableMoves(board)
	randomIndex := rand.Intn(len(moves))

	return moves[randomIndex]
}
