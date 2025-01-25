package doors

import "fmt"

type Doors struct {
}

func (d *Doors) OpenDoors(floor int) {
	fmt.Println("Doors are open at", floor)
}

func (d *Doors) CloseDoors(floor int) {
	fmt.Println("Doors are Closed at", floor)
}
