package board

import (
	"fmt"
)

type Board struct {
	rows       int
	columns    int
	board      [][]string
	freePlaces int
}

func (b *Board) NewBoard(rows int, columns int) {
	b.rows = rows
	b.columns = columns
	b.freePlaces = b.rows * b.columns
	b.board = make([][]string, b.rows)
	for i := 0; i < b.rows; i++ {
		b.board[i] = make([]string, b.columns)
	}
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.columns; j++ {
			b.board[i][j] = " "
		}
	}

}

func (b *Board) AddPieceToBoard(piece string, row int, column int) error {
	if row >= b.rows || column >= b.columns || row < 0 || column < 0 {
		return fmt.Errorf("provide row and column do not exist. given row: %v, given column: %v\n", row, column)
	}
	if b.board[row][column] != " " {
		return fmt.Errorf("provided row and column is already present. given row: %v, given column: %v\n", row, column)
	}
	b.board[row][column] = piece
	b.freePlaces -= 1
	return nil
}

func (b *Board) PrintBoard() {
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.columns; j++ {
			fmt.Printf(" %v ", b.board[i][j])
			fmt.Print("|")
		}
		fmt.Println()
	}
}

func (b *Board) GetFreePlaces() int {
	return b.freePlaces
}
