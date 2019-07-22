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
	fmt.Print("\n\n\nWelcome to Tic Tac Toe!\n\n\n")
	for ok := true; ok; ok = true {

		goterm.Clear()
		app.printAppState()
		app.printAppPrompt()
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		command := parseUserInput(text)
		//Pass to App
		app.handleCommand(command)

		fmt.Println(text)
	}
}

func (app *commandLineApp) init() {
	app.state = AppStateReady
}

func (app *commandLineApp) printAppState() {
	if app.state == AppStateReady {
		fmt.Print("Ready to Play!\n")
	} else {

		fmt.Print("Playing!\n")
	}
}

func (app *commandLineApp) printAppPrompt() {
	if app.state == AppStateReady {
		fmt.Print("[n] : Start new game\n")
		fmt.Print("[x] : Awesome unimplemented features\n")
		fmt.Print("Enter Choice:  ")
	} else {
		fmt.Print("What's your move?: ")
	}
}

func parseUserInput(text string) commandLineAppCommand {

	return AppCommandStartGame

}

func (app *commandLineApp) handleCommand(command commandLineAppCommand) {

	if app.state == AppStateReady {
		//start the game
		app.startGame()
		app.state = AppStatePlaying
	}
	if app.state == AppStatePlaying {
	}
}

func (app *commandLineApp) startGame() {
	app.board.init()
}
