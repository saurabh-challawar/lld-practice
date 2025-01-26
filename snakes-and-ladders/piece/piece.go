package piece

import "snakes-and-ladders/player"

type Piece struct {
	pieceType PieceType
	player    *player.Player
	position  int
}

func NewPiece(pieceType PieceType, player *player.Player) *Piece {
	return &Piece{
		pieceType: pieceType,
		player:    player,
		position:  1,
	}
}

func (piece *Piece) GetPieceType() PieceType {
	return piece.pieceType
}

func (piece *Piece) GetPlayer() *player.Player {
	return piece.player
}

func (piece *Piece) GetPosition() int {
	return piece.position
}

func (piece *Piece) SetPieceType(pieceType PieceType) *Piece {
	piece.pieceType = pieceType
	return piece
}

func (piece *Piece) SetPlayer(player *player.Player) *Piece {
	piece.player = player
	return piece
}

func (piece *Piece) SetPosition(position int) *Piece {
	piece.position = position
	return piece
}
