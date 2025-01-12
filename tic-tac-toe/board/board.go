package board

import (
	"fmt"
)

type Board struct {
	rows    int
	columns int
	board   [][]string
}

func (b *Board) NewBoard(rows int, columns int) {
	b.rows = rows
	b.columns = columns
	b.board = make([][]string, b.rows)
	for i := 0; i < b.rows; i++ {
		b.board[i] = make([]string, b.columns)
	}
}

func (b *Board) AddPieceToBoard(piece string, row int, column int) error {
	if row >= b.rows || column >= b.columns || row < 0 || column < 0 {
		return fmt.Errorf("provide row and column do not exist. given row: %v, given column: %v", row, column)
	}
	b.board[row][column] = piece
	return nil
}
