package elevator

import (
	"elevator-system/display"
	"elevator-system/doors"
	"elevator-system/floor"
)

type ElevatorState string

const (
	goingUp   ElevatorState = "going up"
	goingDown ElevatorState = "going down"
	idle      ElevatorState = "idle"
)

type Elevator struct {
	state          ElevatorState
	display        display.Display
	currentFloor   floor.Floor
	doors          doors.Doors
	InternalButton InternalButton
}
