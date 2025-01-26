package main

import (
	"snakes-and-ladders/game"
	"snakes-and-ladders/piece"
	"snakes-and-ladders/player"
)

func main() {
	player1 := player.NewPlayer(1, "Player1")
	player2 := player.NewPlayer(2, "Player2")
	player3 := player.NewPlayer(3, "Player3")

	piece1 := piece.NewPiece(piece.RED, player1)
	piece2 := piece.NewPiece(piece.GREEN, player2)
	piece3 := piece.NewPiece(piece.BLUE, player3)

	pieceList := []*piece.Piece{
		piece1, piece2, piece3,
	}

	game := game.NewGame(pieceList)
	game.StartGame()

}
