package main

import (
	"fmt"
	"time"
)

type RentManager struct {
	rentList map[string]*Rent
}

func NewRentManager() *RentManager {
	return &RentManager{
		rentList: make(map[string]*Rent),
	}
}

func (rm *RentManager) BookVehicle(vehicle *Vehicle, store *Store, user *User, days int) {

	rentObj := NewRent(user, vehicle, time.Now(), time.Now().Add(time.Duration(days)*24*time.Hour))
	fmt.Println("Rent obj", rentObj)
	rm.rentList[vehicle.registrationNumber] = rentObj
	store.inventoryManagement.RemoveVehicle(vehicle)
}

func (rm *RentManager) ReturnVehicle(vehicle *Vehicle, store *Store, user *User) {
	rentObj := rm.rentList[vehicle.registrationNumber]
	bill := &Bill{
		receiptStr: vehicle.registrationNumber + "-" + user.name + "-" + time.Now().GoString(),
	}
	if !bill.PayBill(rentObj) {
		return
	}
	delete(rm.rentList, rentObj.vehicle.registrationNumber)
	store.inventoryManagement.AddVehicle(vehicle)
}
