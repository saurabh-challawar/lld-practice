package main

type InventoryManagement struct {
	vehicleList map[string]*Vehicle
}

func (im *InventoryManagement) SetInventory() {
	if im.vehicleList == nil {
		im.vehicleList = make(map[string]*Vehicle)
	}
}

func (im *InventoryManagement) GetVehicles() map[string]*Vehicle {
	return im.vehicleList
}

func (im *InventoryManagement) AddVehicle(vehicle *Vehicle) {
	im.vehicleList[vehicle.registrationNumber] = vehicle
}

func (im *InventoryManagement) RemoveVehicle(Vehicle *Vehicle) {
	delete(im.vehicleList, Vehicle.registrationNumber)
}
