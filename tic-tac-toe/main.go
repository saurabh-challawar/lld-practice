package main

import (
	"fmt"
	"tic-tac-toe/game"
)

func main() {
	fmt.Println("Welcome to TicTacToe")
	var game game.Game
	game.StartGame()
}
