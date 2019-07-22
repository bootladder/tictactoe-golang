package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeMove_emptySquare_marksSquare(t *testing.T) {
	var board tictactoeboard
	board.makeMove(SquareX, 0, 0)
	assert.Equal(t, board.getSquareValue(0, 0), SquareX)

	board.makeMove(SquareO, 1, 1)
	assert.Equal(t, board.getSquareValue(1, 1), SquareO)
}

func Test_makeMove_notEmptySquare_returnsError(t *testing.T) {
	var board tictactoeboard
	err := board.makeMove(SquareX, 0, 0)
	err = board.makeMove(SquareX, 0, 0)

	assert.Error(t, err, "Can't make move twice: ")
}

func Test_determineBoardState_XWins(t *testing.T) {
	var board tictactoeboard
	board.makeMove(SquareX, 0, 0)
	board.makeMove(SquareX, 0, 1)
	board.makeMove(SquareX, 0, 2)

	state := board.determineBoardState()
	assert.Equal(t, WinnerX, state)
}

func Test_determineBoardState_XWins_AcrossTheMiddle(t *testing.T) {
	var board tictactoeboard
	board.makeMove(SquareX, 1, 0)
	board.makeMove(SquareX, 1, 1)
	board.makeMove(SquareX, 1, 2)

	state := board.determineBoardState()
	assert.Equal(t, WinnerX, state)

	fmt.Printf("%v", board.board)
}

func Test_determineBoardState_XWins_DownTheMiddle(t *testing.T) {
	var board tictactoeboard
	board.makeMove(SquareX, 0, 0)
	board.makeMove(SquareX, 1, 0)
	board.makeMove(SquareX, 2, 0)

	state := board.determineBoardState()
	assert.Equal(t, WinnerX, state)

	fmt.Printf("%v", board.board)
}

func Test_determineBoardState_OWins(t *testing.T) {
	var board tictactoeboard
	board.makeMove(SquareO, 0, 0)
	board.makeMove(SquareO, 0, 1)
	board.makeMove(SquareO, 0, 2)

	state := board.determineBoardState()
	assert.Equal(t, WinnerO, state)
}

func NOT____Test_determineBoardState_TieGame(t *testing.T) {
	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareX, SquareO, SquareX},
		{SquareO, SquareO, SquareX},
		{SquareX, SquareX, SquareO},
	}

	state := board.determineBoardState()
	assert.Equal(t, TieGame, state)
}
