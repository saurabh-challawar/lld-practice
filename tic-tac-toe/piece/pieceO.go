package piece

type PieceO struct {
	piece string
}

func (p *PieceO) SetPiece() {
	p.piece = "X"
}

func (p *PieceO) GetPiece() string {
	return p.piece
}
