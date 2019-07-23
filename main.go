package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/buger/goterm"
)

type commandLineAppState int

const (
	AppStateReady         commandLineAppState = 0
	AppStatePlaying       commandLineAppState = 1
	AppStateGameOver      commandLineAppState = 2
	AppStateDoingFuzzTest commandLineAppState = 3
)

type commandLineAppCommand struct {
	ctype commandLineAppCommandType
	cdata int // int for now
}

type commandLineAppCommandType int

const (
	AppCommandStartGame commandLineAppCommandType = 0
	AppCommandMakeMove  commandLineAppCommandType = 1
	AppCommandFuzzTest  commandLineAppCommandType = 2
)

type commandLineApp struct {
	state commandLineAppState
	board tictactoeboard
}

func main() {

	var app commandLineApp
	app.init()

	for ok := true; ok; ok = true {

		clearScreen()
		app.printAppState()
		app.printAppPrompt()

		//Read command from user
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		command, err := parseUserInput(text)
		if err != nil {
			continue
		}
		app.handleCommand(command)
	}
}

func clearScreen() {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	goterm.Flush()
}

func (app *commandLineApp) init() {
	app.state = AppStateReady
}

func (app *commandLineApp) printAppState() {
	switch state := app.state; state {
	case AppStateReady:
		fmt.Print("Ready to Play!\n")
	case AppStatePlaying:
		fmt.Print("Playing!\n")
		app.printBoard()
	case AppStateGameOver:
		fmt.Print("GAME OVER!\n")
		switch app.board.determineBoardState() {
		case WinnerX:
			fmt.Print("X Wins!\n")
		case WinnerO:
			fmt.Print("O Wins!\n")
		case TieGame:
			fmt.Print("Tie Game!\n")
		}
		app.printBoard()
	case AppStateDoingFuzzTest:
		fmt.Print("Fuzzing... will stop when CPU loses\n")
	}
}

func (app *commandLineApp) printAppPrompt() {
	switch state := app.state; state {
	case AppStateReady:
		fmt.Print("[n] : Start new game\n")
		fmt.Print("[f] : Do CPU Fuzz Test\n")
		fmt.Print("[x] : Awesome unimplemented features\n")
		fmt.Print("Enter Choice:  ")
	case AppStatePlaying:
		fmt.Print("To move, enter the number of the square.\n")
		fmt.Print("What's your move?: ")
	case AppStateGameOver:
		fmt.Print("[n] : Start new game\n")
		fmt.Print("[x] : Awesome unimplemented features\n")
		fmt.Print("Enter Choice:  ")
	case AppStateDoingFuzzTest:
		fmt.Print("CTRL-C to exit\n")
	}
}

func parseUserInput(text string) (commandLineAppCommand, error) {

	if text == "f" {
		return commandLineAppCommand{AppCommandFuzzTest, 1}, nil
	}

	if text == "n" {
		return commandLineAppCommand{AppCommandStartGame, 1}, nil
	}

	if text == "1" {
		return commandLineAppCommand{AppCommandMakeMove, 1}, nil
	}

	if text == "2" {
		return commandLineAppCommand{AppCommandMakeMove, 2}, nil
	}

	if text == "3" {
		return commandLineAppCommand{AppCommandMakeMove, 3}, nil
	}

	if text == "4" {
		return commandLineAppCommand{AppCommandMakeMove, 4}, nil
	}

	if text == "5" {
		return commandLineAppCommand{AppCommandMakeMove, 5}, nil
	}

	if text == "6" {
		return commandLineAppCommand{AppCommandMakeMove, 6}, nil
	}

	if text == "7" {
		return commandLineAppCommand{AppCommandMakeMove, 7}, nil
	}

	if text == "8" {
		return commandLineAppCommand{AppCommandMakeMove, 8}, nil
	}

	if text == "9" {
		return commandLineAppCommand{AppCommandMakeMove, 9}, nil
	}

	return commandLineAppCommand{AppCommandStartGame, 1}, errors.New("Invalid")
}

func (app *commandLineApp) handleCommand(command commandLineAppCommand) {

	switch cmd := command.ctype; cmd {

	case AppCommandStartGame:
		app.startGame()
		app.state = AppStatePlaying

	case AppCommandMakeMove:

		//Make Human Move (Human always plays X)
		row, col := squareNumberToRowCol(command.cdata)
		app.board.makeMove(SquareX, row, col)

		//Check Game State
		boardState := app.board.determineBoardState()

		if boardState != GameInProgress {
			app.state = AppStateGameOver
			break
		}

		//Make CPU move
		cpuRow, cpuCol := cpuMakeMove(app.board)
		app.board.makeMove(SquareO, cpuRow, cpuCol)

		//Check Game State
		boardState = app.board.determineBoardState()

		if boardState != GameInProgress {
			app.state = AppStateGameOver
			break
		}

	case AppCommandFuzzTest:
		//for now just doing another infinite loop in here

		fuzzBoard := tictactoeboard{}

		for ok := true; ok; ok = true {
			fmt.Print("\nPlaying fuzz game...")
			if fuzzBoard.determineBoardState() == WinnerX {
				fmt.Print("\nX won.... not good!!")
				fmt.Printf("%v", fuzzBoard.board)
				os.Exit(2)
			}
			if fuzzBoard.determineBoardState() == WinnerO {
				fmt.Print("O won")
			}
			if fuzzBoard.determineBoardState() == TieGame {
				fmt.Print("Tie Game")
			}
			fuzzBoard.init()

			for zzz := true; zzz; zzz = true {
				//Make Random Human Move
				move := cpuPickRandomMove(fuzzBoard)
				//fmt.Printf("X Plays: %d %d\n", move.row, move.col)
				fuzzBoard.makeMove(SquareX, move.row, move.col)

				//Check Game State
				boardState := fuzzBoard.determineBoardState()

				if boardState != GameInProgress {
					break
				}

				//Make CPU move
				cpuRow, cpuCol := cpuMakeMove(fuzzBoard)
				//fmt.Printf("O Plays: %d %d\n", cpuRow, cpuCol)
				fuzzBoard.makeMove(SquareO, cpuRow, cpuCol)

				//Check Game State
				boardState = fuzzBoard.determineBoardState()

				if boardState != GameInProgress {
					break
				}
			}
		}

	default:
		fmt.Print("Z")
	}
}

func squareNumberToRowCol(num int) (int, int) {

	switch n := num; n {
	case 1:
		return 0, 0
	case 2:
		return 0, 1
	case 3:
		return 0, 2
	case 4:
		return 1, 0
	case 5:
		return 1, 1
	case 6:
		return 1, 2
	case 7:
		return 2, 0
	case 8:
		return 2, 1
	case 9:
		return 2, 2
	}

	return 0, 0
}

func (app *commandLineApp) startGame() {
	app.board.init()
}

func (app *commandLineApp) printBoard() {

	b00 := app.board.getSquareValue(0, 0)
	b01 := app.board.getSquareValue(0, 1)
	b02 := app.board.getSquareValue(0, 2)
	b10 := app.board.getSquareValue(1, 0)
	b11 := app.board.getSquareValue(1, 1)
	b12 := app.board.getSquareValue(1, 2)
	b20 := app.board.getSquareValue(2, 0)
	b21 := app.board.getSquareValue(2, 1)
	b22 := app.board.getSquareValue(2, 2)

	m := make(map[squareValue]string)
	m[SquareX] = "X"
	m[SquareO] = "O"
	m[SquareEmpty] = "-"

	s00 := m[b00]
	s01 := m[b01]
	s02 := m[b02]
	s10 := m[b10]
	s11 := m[b11]
	s12 := m[b12]
	s20 := m[b20]
	s21 := m[b21]
	s22 := m[b22]

	fmt.Printf("| %v | %v | %v |\n", s00, s01, s02)
	fmt.Printf("| %v | %v | %v |\n", s10, s11, s12)
	fmt.Printf("| %v | %v | %v |\n", s20, s21, s22)
}
