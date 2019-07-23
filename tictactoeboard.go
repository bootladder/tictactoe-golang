package main

import "errors"

type squareValue int

// Hello
const (
	SquareEmpty squareValue = 0
	SquareX     squareValue = 1
	SquareO     squareValue = 2
)

type boardState int

// Hello
const (
	WinnerX        boardState = 0
	WinnerO        boardState = 1
	GameInProgress boardState = 2
	TieGame        boardState = 3
)

type tictactoeboard struct {
	board [3][3]squareValue
}

func (b *tictactoeboard) init() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.board[i][j] = SquareEmpty
		}
	}
}

func (b *tictactoeboard) makeMove(player squareValue, row, col int) error {
	if b.board[row][col] != SquareEmpty {
		return errors.New("Square Already Taken")
	}
	b.board[row][col] = player
	return nil
}

func (b *tictactoeboard) getSquareValue(row, col int) squareValue {
	return b.board[row][col]
}

func (b *tictactoeboard) determineBoardState() boardState {

	if b.checkWinner(SquareX) {
		return WinnerX
	}
	if b.checkWinner(SquareO) {
		return WinnerO
	}

	if b.allSquaresTaken() {
		return TieGame
	}
	return GameInProgress
}

func (b *tictactoeboard) checkWinner(val squareValue) bool {

	//try all of the 3-in-a-rows
	b00 := b.board[0][0]
	b01 := b.board[0][1]
	b02 := b.board[0][2]
	b10 := b.board[1][0]
	b11 := b.board[1][1]
	b12 := b.board[1][2]
	b20 := b.board[2][0]
	b21 := b.board[2][1]
	b22 := b.board[2][2]

	//rows
	if b00 == b01 && b01 == b02 && b02 == val {
		return true
	}
	if b10 == b11 && b11 == b12 && b12 == val {
		return true
	}
	if b20 == b21 && b21 == b22 && b22 == val {
		return true
	}
	//cols
	if b00 == b10 && b10 == b20 && b20 == val {
		return true
	}
	if b01 == b11 && b11 == b21 && b21 == val {
		return true
	}
	if b02 == b12 && b12 == b22 && b22 == val {
		return true
	}
	//diagonals
	if b00 == b11 && b11 == b22 && b22 == val {
		return true
	}
	if b02 == b11 && b11 == b20 && b20 == val {
		return true
	}
	return false
}

func (b *tictactoeboard) allSquaresTaken() bool {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.board[i][j] == SquareEmpty {
				return false
			}
		}
	}
	return true
}
