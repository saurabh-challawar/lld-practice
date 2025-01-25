package floor

type Direction string

const (
	up   Direction = "up"
	down Direction = "down"
)

type ExternalButton struct {
}

func (b *ExternalButton) PressButton(floorNumber int, direction Direction) {
	// logic to be implemented
}
