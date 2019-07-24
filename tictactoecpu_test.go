package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cpuFindForkMove_(t *testing.T) {
	var board tictactoeboard

	//Corner square is a fork
	board.board = [3][3]squareValue{
		{SquareEmpty, SquareO, SquareEmpty},
		{SquareEmpty, SquareX, SquareO},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	row, col, err := cpuFindForkMove(board)
	assert.NoError(t, err)
	assert.Equal(t, 0, row)
	assert.Equal(t, 2, col)

	//Middle Square is a Fork
	board.board = [3][3]squareValue{
		{SquareEmpty, SquareO, SquareX},
		{SquareEmpty, SquareEmpty, SquareO},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	row, col, err = cpuFindForkMove(board)
	assert.NoError(t, err)
	assert.Equal(t, 1, row)
	assert.Equal(t, 1, col)
}

//Test the helper method
func Test_cpuFindForksForPlayer(t *testing.T) {

	var board tictactoeboard

	//Corner square is a fork for O
	board.board = [3][3]squareValue{
		{SquareEmpty, SquareO, SquareEmpty},
		{SquareEmpty, SquareX, SquareO},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	forkMoves := cpuFindForksForPlayer(board, SquareO)
	assert.Equal(t, 1, len(forkMoves))
	assert.Equal(t, 0, forkMoves[0].row)
	assert.Equal(t, 2, forkMoves[0].col)
}

func Test_cpuFindOppositeCorner(t *testing.T) {
	var board tictactoeboard

	board.board = [3][3]squareValue{
		{SquareX, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	row, col, err := cpuFindOppositeCorner(board)
	assert.NoError(t, err)
	assert.Equal(t, 2, row)
	assert.Equal(t, 2, col)
}

func Test_cpuFindOppositeCorner_noOppositeCorner_returnsError(t *testing.T) {
	var board tictactoeboard

	board.board = [3][3]squareValue{
		{SquareX, SquareEmpty, SquareX},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareO, SquareEmpty, SquareO},
	}

	_, _, err := cpuFindOppositeCorner(board)
	assert.Error(t, err)
}

func Test_cpuFindEmptyCorner(t *testing.T) {
	var board tictactoeboard

	board.board = [3][3]squareValue{
		{SquareX, SquareEmpty, SquareX},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareO, SquareEmpty, SquareEmpty},
	}

	row, col, err := cpuFindEmptyCorner(board)
	assert.NoError(t, err)
	assert.Equal(t, 2, row)
	assert.Equal(t, 2, col)
}

func Test_cpuFindEmptySide(t *testing.T) {
	var board tictactoeboard

	board.board = [3][3]squareValue{
		{SquareX, SquareEmpty, SquareX},
		{SquareO, SquareEmpty, SquareX},
		{SquareO, SquareO, SquareEmpty},
	}

	row, col, err := cpuFindEmptySide(board)
	assert.NoError(t, err)
	assert.Equal(t, 0, row)
	assert.Equal(t, 1, col)
}

func Test_cpuDoesPlayerHaveTwoInARow(t *testing.T) {
	var board tictactoeboard

	board.board = [3][3]squareValue{
		{SquareX, SquareEmpty, SquareX},
		{SquareO, SquareEmpty, SquareX},
		{SquareO, SquareO, SquareEmpty},
	}

	twoInARow := cpuDoesPlayerHaveTwoInARow(board, SquareO)
	assert.Equal(t, true, twoInARow)

	board.board = [3][3]squareValue{
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}

	twoInARow = cpuDoesPlayerHaveTwoInARow(board, SquareO)
	assert.Equal(t, false, twoInARow)
}

////////////////////////////////////////////////////
// "old tests from before implementing wikipedia strategy"
///////////////////
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

func assertNonCornerMove(t *testing.T, row, col int) {
	ok := false
	if row == 0 && col == 1 {
		ok = true
	}
	if row == 1 && col == 0 {
		ok = true
	}
	if row == 1 && col == 2 {
		ok = true
	}
	if row == 2 && col == 1 {
		ok = true
	}
	assert.Equal(t, true, ok)
}

func Test_cpuMakeMove_cpuCouldWinRightNow_picksWinningMove(t *testing.T) {
	var board tictactoeboard
	board.board = [3][3]squareValue{
		{SquareX, SquareO, SquareEmpty},
		{SquareEmpty, SquareO, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareX},
	}

	row, col := cpuMakeMove(board)
	assert.Equal(t, 2, row)
	assert.Equal(t, 1, col)
}
