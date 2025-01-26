package board

type Board struct {
	positions []int
	snakes    map[int]int
	ladders   map[int]int
}

func NewBoard(snakes map[int]int, ladders map[int]int) *Board {
	return &Board{
		positions: make([]int, 100),
		snakes:    snakes,
		ladders:   ladders,
	}
}

func (board *Board) GetPositions() []int {
	return board.positions
}

func (board *Board) GetSnakes() map[int]int {
	return board.snakes
}

func (board *Board) GetLadders() map[int]int {
	return board.ladders
}

func (board *Board) SetPositions(positions []int) *Board {
	board.positions = positions
	return board
}

func (board *Board) SetSnakes(snakes map[int]int) *Board {
	board.snakes = snakes
	return board
}

func (board *Board) SetLadders(ladders map[int]int) *Board {
	board.ladders = ladders
	return board
}
