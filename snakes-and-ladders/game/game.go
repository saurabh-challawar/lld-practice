package game

import (
	"fmt"
	"math/rand"
	"snakes-and-ladders/board"
	"snakes-and-ladders/piece"
)

type Game struct {
	pieceList       []*piece.Piece
	numberOfPlayers int
	board           *board.Board
	gameComplete    bool
}

func NewGame(pieceList []*piece.Piece) *Game {

	snakes := map[int]int{
		16: 6,  // Snake from 16 to 6
		47: 26, // Snake from 47 to 26
		49: 11, // Snake from 49 to 11
		56: 53, // Snake from 56 to 53
		62: 19, // Snake from 62 to 19
		64: 60, // Snake from 64 to 60
		87: 24, // Snake from 87 to 24
		93: 73, // Snake from 93 to 73
		95: 75, // Snake from 95 to 75
		98: 78, // Snake from 98 to 78
	}

	ladders := map[int]int{
		4:  14,  // Ladder from 4 to 14
		9:  31,  // Ladder from 9 to 31
		21: 42,  // Ladder from 21 to 42
		28: 84,  // Ladder from 28 to 84
		36: 44,  // Ladder from 36 to 44
		51: 67,  // Ladder from 51 to 67
		71: 91,  // Ladder from 71 to 91
		80: 100, // Ladder from 80 to 100
	}

	board := board.NewBoard(snakes, ladders)

	return &Game{
		pieceList:       pieceList,
		board:           board,
		numberOfPlayers: len(pieceList),
		gameComplete:    false,
	}
}

func (game *Game) StartGame() {
	idx := 0

	snakes := game.board.GetSnakes()
	ladders := game.board.GetLadders()

	for !game.gameComplete {
		currentPiece := game.pieceList[idx]
		diceOutput := game.RollDice()

		fmt.Printf("%v rolled %v with colour %v.\n", currentPiece.GetPlayer().GetName(), diceOutput, currentPiece.GetPieceType())

		if currentPiece.GetPosition()+diceOutput > 100 {
			idx = (idx + 1) % game.numberOfPlayers
			continue
		}
		if value, ok := snakes[currentPiece.GetPosition()+diceOutput]; ok {
			fmt.Println("It's a Snake. Going down.")
			currentPiece.SetPosition(value)
		} else if value, ok := ladders[currentPiece.GetPosition()+diceOutput]; ok {
			fmt.Println("It's a Ladder. Going up.")
			currentPiece.SetPosition(value)
		} else {
			currentPiece.SetPosition(currentPiece.GetPosition() + diceOutput)
		}

		fmt.Printf("%v with colour %v is at %v.\n", currentPiece.GetPlayer().GetName(), currentPiece.GetPieceType(), currentPiece.GetPosition())
		fmt.Println()
		fmt.Println()

		if currentPiece.GetPosition() == 100 {
			fmt.Printf("%v won with colour %v.\n", currentPiece.GetPlayer().GetName(), currentPiece.GetPieceType())
			game.gameComplete = true
		}
		fmt.Println()
		idx = (idx + 1) % game.numberOfPlayers

	}
}

func (game *Game) RollDice() int {
	min := 1
	max := 6
	return rand.Intn(max-min+1) + min
}
