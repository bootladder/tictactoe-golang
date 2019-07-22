package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/buger/goterm"
)

type commandLineAppState int

// Hello
const (
	AppStateReady    commandLineAppState = 0
	AppStatePlaying  commandLineAppState = 1
	AppStateGameOver commandLineAppState = 2
)

type commandLineAppCommand int

//Hello
const (
	AppCommandStartGame commandLineAppCommand = 0
	AppCommandMakeMove  commandLineAppCommand = 0
)

type commandLineApp struct {
	state commandLineAppState
	board tictactoeboard
}

func main() {

	var app commandLineApp
	app.init()

	clearScreen()
	fmt.Print("\n\n\nWelcome to Tic Tac Toe!\n\n\n")

	for ok := true; ok; ok = true {

		app.printAppState()
		app.printAppPrompt()
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		command := parseUserInput(text)
		//Pass to App
		app.handleCommand(command)

		fmt.Println(text)
		clearScreen()
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
	if app.state == AppStateReady {
		fmt.Print("Ready to Play!\n")
	} else {

		fmt.Print("Playing!\n")
		app.printBoard()
	}
}

func (app *commandLineApp) printAppPrompt() {
	if app.state == AppStateReady {
		fmt.Print("[n] : Start new game\n")
		fmt.Print("[x] : Awesome unimplemented features\n")
		fmt.Print("Enter Choice:  ")
	} else {
		fmt.Print("To move, enter the number of the square.\n")
		fmt.Print("What's your move?: ")
	}
}

func parseUserInput(text string) commandLineAppCommand {

	return AppCommandStartGame

}

func (app *commandLineApp) handleCommand(command commandLineAppCommand) {

	if app.state == AppStateReady {
		app.startGame()
		app.state = AppStatePlaying
	}
	if app.state == AppStatePlaying {
	}
}

func (app *commandLineApp) startGame() {
	app.board.init()
}

func (app *commandLineApp) printBoard() {

	b00 := app.board.getSquareValue(0, 0)
	b01 := app.board.getSquareValue(0, 0)
	b02 := app.board.getSquareValue(0, 0)
	b10 := app.board.getSquareValue(0, 0)
	b11 := app.board.getSquareValue(0, 0)
	b12 := app.board.getSquareValue(0, 0)
	b20 := app.board.getSquareValue(0, 0)
	b21 := app.board.getSquareValue(0, 0)
	b22 := app.board.getSquareValue(0, 0)

	fmt.Printf("| %v | %v | %v |\n", b00, b01, b02)
	fmt.Printf("| %v | %v | %v |\n", b10, b11, b12)
	fmt.Printf("| %v | %v | %v |\n", b20, b21, b22)
}
