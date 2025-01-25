package main

import "time"

type Rent struct {
	id                 int
	user               *User
	vehicle            *Vehicle
	startTime          time.Time
	expectedReturnTime time.Time
	actualReturnTime   time.Time
}

func NewRent(user *User, vehicle *Vehicle, startTime time.Time, expectedReturnTime time.Time) *Rent {
	var rent Rent
	rent.user = user
	rent.vehicle = vehicle
	rent.startTime = startTime
	rent.expectedReturnTime = expectedReturnTime
	return &rent
}
