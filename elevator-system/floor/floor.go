package floor

type Floor struct {
	floorNumber    int
	externalButton ExternalButton
}

func NewFloor(floorNumber int) *Floor {
	return &Floor{
		floorNumber:    floorNumber,
		externalButton: ExternalButton{},
	}

}

func (f *Floor) PressButton(direction Direction, floorNumber int) {
	f.externalButton.PressButton(floorNumber, direction)
}
