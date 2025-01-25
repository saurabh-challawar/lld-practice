package main

import "time"

type Rent struct {
	id                 int
	user               User
	vehicle            Vehicle
	startTime          time.Time
	expectedReturnTime time.Time
	actualReturnTime   time.Time
}

func (rent *Rent) NewRent(user User, vehicle Vehicle, startTime time.Time, expectedReturnTime time.Time) {
	rent.user = user
	rent.vehicle = vehicle
	rent.startTime = startTime
	rent.expectedReturnTime = expectedReturnTime
}
