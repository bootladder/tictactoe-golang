package main

import (
	"errors"
	"math/rand"
	"time"
)

// Following the strategy from Wikipedia:
// https://en.wikipedia.org/wiki/Tic-tac-toe#Strategy
// There are 8 rules:

func cpuMakeMove(board tictactoeboard) (int, int) {

	// Rule #1: Win
	if cpuCanWinRightNow(board) {
		move := cpuPickWinningMove(board)
		return move.row, move.col
	}

	// Rule #2: Block
	row, col, err := cpuGetMoveThatStopsThreeInARow(board)
	if err == nil {
		return row, col
	}

	// Rule #3: Fork (Move that makes two 2-in-a-row's)
	row, col, err = cpuFindForkMove(board)
	if err == nil {
		return row, col
	}

	// Rule #4: Block Fork
	// If Opponent has 1 possible fork, block it
	// If Opponent has >1 possible forks AND blocking one of them makes a 2-in-a-row, pick it
	// If CPU can make a 2-in-a-row, and defending it does not create a fork for opponent, pick it
	forkMoves := cpuFindForksForPlayer(board, SquareX)
	//fmt.Printf("LEN OF FORKMOVES IS    %d  ", len(forkMoves))
	//fmt.Printf("%v", forkMoves)
	if len(forkMoves) == 1 {
		return forkMoves[0].row, forkMoves[0].col
	}

	if len(forkMoves) > 1 {
		// Find moves that block one of the forks
		// AND create a 2 in a row

		// If we make a move and it lowers the number of forks, then it's a block
		for _, move := range cpuGetAvailableMoves(board) {
			tempBoard := board
			tempBoard.makeMove(SquareO, move.row, move.col)

			forkMovesAfter := cpuFindForksForPlayer(tempBoard, SquareX)
			if len(forkMovesAfter) < len(forkMoves) {
				//got a fork blocker.  does it also make a 2 in a row for us?
				if cpuDoesPlayerHaveTwoInARow(tempBoard, SquareO) {

					//got a two in a row.  Does defending the two in a row
					//give the opponent a fork?
					blockingRow, blockingCol, err := cpuGetMoveThatStopsThreeInARowForPlayer(tempBoard, SquareO)
					if err == nil {

						opponentStoppingWillMakeAFork := false

						forkMovesFinally := cpuFindForksForPlayer(tempBoard, SquareX)
						for _, forkMoveFinally := range forkMovesFinally {

							if forkMoveFinally.row == blockingRow &&
								forkMoveFinally.col == blockingCol {
								opponentStoppingWillMakeAFork = true
							}
						}

						if opponentStoppingWillMakeAFork == false {
							return move.row, move.col
						}
					}

				}
			}
		}
	}

	//if cpuDoesPlayerHaveTwoInARow(board, SquareO) {
	//	fmt.Print("ZZZZZZZZZZZZZZZZZZZZZZZZZZZZ")
	//	// AND defending it does not create a fork, pick it
	//}

	// Rule #5: Take Center
	if board.getSquareValue(1, 1) == SquareEmpty {
		return 1, 1
	}

	// Rule #6: Opposite Corner (If opponent has a corner, pick opposite corner)
	row, col, err = cpuFindOppositeCorner(board)
	if err == nil {
		return row, col
	}
	// Rule #7: Empty Corner
	row, col, err = cpuFindEmptyCorner(board)
	if err == nil {
		return row, col
	}
	// Rule #8: Empty Side
	row, col, err = cpuFindEmptySide(board)
	if err == nil {
		return row, col
	}

	// Catch-all: Random Move
	move := cpuPickRandomMove(board)
	return move.row, move.col
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

	return cpuGetMoveThatStopsThreeInARowForPlayer(board, SquareX)
}

func cpuGetMoveThatStopsThreeInARowForPlayer(board tictactoeboard, player squareValue) (int, int, error) {

	availableMoves := cpuGetAvailableMoves(board)

	// If the player were to pick this move, would they win?
	for _, move := range availableMoves {
		tempBoard := board
		tempBoard.makeMove(player, move.row, move.col)
		if tempBoard.determineBoardState() == WinnerX &&
			player == SquareX {
			return move.row, move.col, nil
		}
		if tempBoard.determineBoardState() == WinnerO &&
			player == SquareO {
			return move.row, move.col, nil
		}
	}

	return 0, 0, errors.New("No Moves that Stops 3 in a row")
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

func cpuFindForkMove(board tictactoeboard) (int, int, error) {

	forkMoves := cpuFindForksForPlayer(board, SquareO)
	if len(forkMoves) > 0 {
		return forkMoves[0].row, forkMoves[0].col, nil
	}
	return 0, 0, errors.New("No Forks Found")
}

func cpuFindForksForPlayer(board tictactoeboard, player squareValue) []rowcolTuple {

	forkMoves := []rowcolTuple{}

	availableMoves := cpuGetAvailableMoves(board)

	// Make the move on a temp board, then count the number of winning squares for CPU.
	// If the number of winning squares is 2 or more, then the square is a fork
	for _, move := range availableMoves {
		tempBoard := board
		tempBoard.makeMove(player, move.row, move.col)

		availableMovesAfterPlacingMove := cpuGetAvailableMoves(tempBoard)

		countWinningSquares := 0
		for _, afterMove := range availableMovesAfterPlacingMove {

			tempBoardAfter := tempBoard
			tempBoardAfter.makeMove(player, afterMove.row, afterMove.col)

			//ugly
			if tempBoardAfter.determineBoardState() == WinnerO &&
				player == SquareO {
				countWinningSquares++
			}
			if tempBoardAfter.determineBoardState() == WinnerX &&
				player == SquareX {
				countWinningSquares++
			}
		}

		if countWinningSquares >= 2 {
			forkMoves = append(forkMoves, move)
		}
	}

	return forkMoves
}

func cpuPickRandomMove(board tictactoeboard) rowcolTuple {
	rand.Seed(time.Now().UnixNano())
	moves := cpuGetAvailableMoves(board)
	randomIndex := rand.Intn(len(moves))

	return moves[randomIndex]
}

func cpuPickCorner(board tictactoeboard) (rowcolTuple, error) {
	if board.board[0][0] == SquareEmpty {
		return rowcolTuple{0, 0}, nil
	}
	if board.board[0][2] == SquareEmpty {
		return rowcolTuple{0, 2}, nil
	}
	if board.board[2][0] == SquareEmpty {
		return rowcolTuple{2, 0}, nil
	}
	if board.board[2][2] == SquareEmpty {
		return rowcolTuple{2, 2}, nil
	}

	return rowcolTuple{0, 0}, errors.New("No Corner Available")
}

func cpuOpponentHasOppositeCorners(board tictactoeboard) bool {
	if board.board[0][0] == SquareX && board.board[2][2] == SquareX {
		return true
	}
	if board.board[0][2] == SquareX && board.board[2][0] == SquareX {
		return true
	}
	return false
}
func cpuPickNonCorner(board tictactoeboard) (rowcolTuple, error) {

	if board.board[0][1] == SquareEmpty {
		return rowcolTuple{0, 1}, nil
	}
	if board.board[1][0] == SquareEmpty {
		return rowcolTuple{1, 0}, nil
	}
	if board.board[1][2] == SquareEmpty {
		return rowcolTuple{1, 2}, nil
	}
	if board.board[2][1] == SquareEmpty {
		return rowcolTuple{2, 1}, nil
	}

	return rowcolTuple{0, 0}, errors.New("No Non Corner Available")
}

func cpuCanWinRightNow(board tictactoeboard) bool {

	availableMoves := cpuGetAvailableMoves(board)

	// If the CPU were to pick this move, would CPU win?
	for _, move := range availableMoves {
		tempBoard := board
		tempBoard.makeMove(SquareO, move.row, move.col)
		if tempBoard.determineBoardState() == WinnerO {
			return true
		}
	}
	return false
}

func cpuPickWinningMove(board tictactoeboard) rowcolTuple {
	availableMoves := cpuGetAvailableMoves(board)

	// If the CPU were to pick this move, would CPU win?
	for _, move := range availableMoves {
		tempBoard := board
		tempBoard.makeMove(SquareO, move.row, move.col)
		if tempBoard.determineBoardState() == WinnerO {
			return move
		}
	}
	return rowcolTuple{0, 0} // THIS SHLD NOT HAPPEN!
}

func cpuFindOppositeCorner(board tictactoeboard) (int, int, error) {
	if board.getSquareValue(0, 0) == SquareX &&
		board.getSquareValue(2, 2) == SquareEmpty {
		return 2, 2, nil
	}
	if board.getSquareValue(0, 2) == SquareX &&
		board.getSquareValue(2, 0) == SquareEmpty {
		return 2, 0, nil
	}
	if board.getSquareValue(2, 0) == SquareX &&
		board.getSquareValue(0, 2) == SquareEmpty {
		return 0, 2, nil
	}
	if board.getSquareValue(2, 2) == SquareX &&
		board.getSquareValue(0, 0) == SquareEmpty {
		return 0, 0, nil
	}
	return 0, 0, errors.New("No Opposite Corner Found")
}

func cpuFindEmptyCorner(board tictactoeboard) (int, int, error) {
	if board.getSquareValue(0, 0) == SquareEmpty {
		return 0, 0, nil
	}
	if board.getSquareValue(0, 2) == SquareEmpty {
		return 0, 2, nil
	}
	if board.getSquareValue(2, 0) == SquareEmpty {
		return 2, 0, nil
	}
	if board.getSquareValue(2, 2) == SquareEmpty {
		return 2, 2, nil
	}

	return 0, 0, errors.New("No Empty Corner Found")
}

func cpuFindEmptySide(board tictactoeboard) (int, int, error) {
	if board.getSquareValue(0, 1) == SquareEmpty {
		return 0, 1, nil
	}
	if board.getSquareValue(1, 0) == SquareEmpty {
		return 1, 0, nil
	}
	if board.getSquareValue(1, 2) == SquareEmpty {
		return 1, 2, nil
	}
	if board.getSquareValue(2, 1) == SquareEmpty {
		return 2, 1, nil
	}

	return 0, 0, errors.New("No Empty Side Found")
}

func cpuDoesPlayerHaveTwoInARow(board tictactoeboard, player squareValue) bool {

	availableMoves := cpuGetAvailableMoves(board)
	for _, move := range availableMoves {
		tempBoard := board
		tempBoard.makeMove(player, move.row, move.col)
		//ugly
		if tempBoard.determineBoardState() == WinnerX &&
			player == SquareX {
			return true
		}
		if tempBoard.determineBoardState() == WinnerO &&
			player == SquareO {
			return true
		}
	}
	return false
}
