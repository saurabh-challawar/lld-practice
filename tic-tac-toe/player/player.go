package player

import "tic-tac-toe/piece"

type Player struct {
	name  string
	id    int
	piece piece.Piece
}

func (pl *Player) NewPlayer(name string, id int, piece piece.Piece) {
	pl.name = name
	pl.id = id
	pl.piece = piece
}

func (pl *Player) SetName(name string) {
	pl.name = name
}

func (pl *Player) GetName() string {
	return pl.name
}

func (pl *Player) GetPlayerPiece() string {
	return pl.piece.GetPiece()
}
