package display

import (
	"elevator-system/floor"
)

type Display struct {
	floor     int
	direction floor.Direction
}

func (d *Display) UpdateDirection(direction floor.Direction) {
	d.direction = direction
}

func (d *Display) UpdateFloor(floor int) {
	d.floor = floor
}
