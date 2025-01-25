package elevator

type InternalButton struct {
	internalButtonList []int
}

func (b *InternalButton) InitialiseFloors(floorNumbers []int) {
	// logic to be implemented
	b.internalButtonList = floorNumbers
}

func (b *InternalButton) PressButton(floorNumber int, elevator Elevator) {
	// logic to be implemented
}
