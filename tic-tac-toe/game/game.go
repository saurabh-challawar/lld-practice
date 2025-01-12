package game

import (
	"fmt"
	"tic-tac-toe/board"
	"tic-tac-toe/piece"
	"tic-tac-toe/player"
)

type Game struct {
	playerList []player.Player
	board      board.Board
	playerTurn int
}

func (g *Game) NewGame() {
	var p1 player.Player
	var p2 player.Player

	var p1_piece piece.PieceX
	var p2_piece piece.PieceO

	p1_piece.SetPiece()
	p2_piece.SetPiece()

	p1.NewPlayer("P1", 1, &p1_piece)
	p2.NewPlayer("P2", 2, &p2_piece)

	g.playerList = append(g.playerList, p1)
	g.playerList = append(g.playerList, p2)

	g.board.NewBoard(3, 3)

	g.playerTurn = 0
}

func (g *Game) StartGame() {
	g.NewGame()

	for !g.anyWinner() && g.board.GetFreePlaces() != 0 {
		g.board.PrintBoard()

		fmt.Printf("Player %v's turn\n", g.playerList[g.playerTurn])
		fmt.Printf("Free pieces %v\n", g.board.GetFreePlaces())

		var row int
		var col int
		fmt.Scanf("%v %v", &row, &col)

		err := g.board.AddPieceToBoard(g.playerList[g.playerTurn].GetPlayerPiece(), row, col)
		if err != nil {
			fmt.Printf("error in adding piece: %v\n", err)
			continue
		}

		if g.anyWinner() {
			fmt.Printf("Player %v won.", g.playerList[g.playerTurn])
			break
		}

		g.playerTurn = g.playerTurn ^ 1
	}
}

func (g *Game) anyWinner() bool {
	return false
}
