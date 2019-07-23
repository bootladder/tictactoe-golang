package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cpuMakeMove_emptyBoard_PicksMiddleSquare(t *testing.T) {
	var board tictactoeboard
	row, col := cpuMakeMove(board)
	assert.Equal(t, 1, row)
	assert.Equal(t, 1, col)
}
func Test_opponentDoesNotHave2InARow_middleSquareAvailable_picksMiddleSquare(t *testing.T) {
	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareX, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}
	row, col := cpuMakeMove(board)
	assert.Equal(t, 1, row)
	assert.Equal(t, 1, col)
}

//does not lose
func Test_cpuMakeMove_opponentHasTwoInARow_PicksCorrectSquare(t *testing.T) {

	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareX, SquareX, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	row, col := cpuMakeMove(board)
	assert.Equal(t, 0, row)
	assert.Equal(t, 2, col)

	//test another case here
	board.board = [3][3]squareValue{
		{SquareX, SquareEmpty, SquareEmpty},
		{SquareX, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	row, col = cpuMakeMove(board)
	assert.Equal(t, 2, row)
	assert.Equal(t, 0, col)
}

func Test_cpuGetAvailableMoves_movesAreAvailable_returnsMoves(t *testing.T) {
	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareX, SquareX, SquareEmpty},
		{SquareX, SquareX, SquareEmpty},
		{SquareX, SquareX, SquareEmpty},
	}
	moves := cpuGetAvailableMoves(board)

	assertMoveInMoves(t, moves, 0, 2)
	assertMoveInMoves(t, moves, 1, 2)
	assertMoveInMoves(t, moves, 2, 2)
}

func assertMoveInMoves(t *testing.T, moves []rowcolTuple, row, col int) {

	good := false
	for _, move := range moves {
		if move.row == row && move.col == col {
			good = true
		}
	}
	assert.Equal(t, true, good)
}

//Picks random move if nothing better
func Test_pickRandomMove_moveIsValid(t *testing.T) {

	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareX, SquareX, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	move := cpuPickRandomMove(board)
	assertMoveInMoves(t, cpuGetAvailableMoves(board), move.row, move.col)
}

func Test_firstMove_humanPicked_square9_cpuTakesMiddleSquare(t *testing.T) {
	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareX},
	}

	row, col := cpuMakeMove(board)
	assert.Equal(t, 1, row)
	assert.Equal(t, 1, col)
}

// This test depends on randomness so will be flaky
func Test_cpuMakeMove_cornerAvailable_opponentDoesNotHaveTwoInARow_middleUnavailable_cpuPicksRandomCorner(t *testing.T) {
	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareX, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	row, col := cpuMakeMove(board)
	assertCornerMove(t, row, col)
}

func assertCornerMove(t *testing.T, row, col int) {
	ok := false
	if row == 0 && col == 0 {
		ok = true
	}
	if row == 0 && col == 2 {
		ok = true
	}
	if row == 2 && col == 0 {
		ok = true
	}
	if row == 2 && col == 2 {
		ok = true
	}
	assert.Equal(t, true, ok)
}
