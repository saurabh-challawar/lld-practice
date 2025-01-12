package piece

type PieceX struct {
	piece string
}

func (p *PieceX) SetPiece() {
	p.piece = "X"
}

func (p *PieceX) GetPiece() string {
	return p.piece
}
