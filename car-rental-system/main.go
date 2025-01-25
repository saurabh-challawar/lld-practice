package main

import "fmt"

func addStores() []*Store {
	var storeList []*Store
	baseName := "store"
	for i := 0; i < 10; i++ {
		store := &Store{
			name:                fmt.Sprintf("%v%v", baseName, i),
			id:                  i,
			inventoryManagement: &InventoryManagement{},
		}
		for j := 0; j < 15; j++ {
			car := &Vehicle{
				id:                 j,
				model:              fmt.Sprintf("car%v", j),
				registrationNumber: fmt.Sprintf("%vcar%v", i, j),
				seater:             4,
				vehicleType:        CAR,
			}
			store.inventoryManagement.SetInventory()
			store.inventoryManagement.AddVehicle(car)
		}
		storeList = append(storeList, store)
	}
	return storeList
}

func addUsers() []*User {
	var userList []*User
	baseName := "user"
	for i := 0; i < 10; i++ {
		user := &User{
			name: baseName + string(i),
			id:   i,
		}
		userList = append(userList, user)
	}
	return userList
}

func main() {
	userList := addUsers()
	storeList := addStores()

	rentalSystem := RentalSystem{
		userList:  userList,
		storeList: storeList,
	}

	location := Location{
		pincode: 500032,
		city:    "Hyderabad",
		state:   "Telangana",
	}

	// users gets results from nearest store
	currentStore := rentalSystem.GetNearbyStore(&location)

	vehiclesInStore := currentStore.inventoryManagement.GetVehicles()

	rentMgr := NewRentManager()
	// selected vehicle
	fmt.Println(vehiclesInStore)
	vehicleSelected := vehiclesInStore["0car3"]
	currentUser := userList[0]

	fmt.Println("Vehicle Selcted by user", vehicleSelected)

	// booked vehicle
	rentMgr.BookVehicle(vehicleSelected, currentStore, currentUser, 3)

	// returning vehicle to another store
	rentMgr.ReturnVehicle(vehicleSelected, storeList[3], currentUser)

	// check to see the returned car is in new store
	fmt.Println(storeList[3].inventoryManagement.GetVehicles())
}
