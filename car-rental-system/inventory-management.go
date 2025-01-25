package main

type InventoryManagement struct {
	VehicleList map[string]Vehicle
}

func (im *InventoryManagement) AddVehicle(vehicle Vehicle) {
	im.VehicleList[vehicle.registrationNumber] = vehicle
}

func (im *InventoryManagement) RemoveVehicle(Vehicle Vehicle) {
	delete(im.VehicleList, Vehicle.registrationNumber)
}
