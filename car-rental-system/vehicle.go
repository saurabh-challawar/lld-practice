package main

type Vehicle struct {
	id                 int
	registrationNumber string
	model              string
	make               string
	seater             int
	bootSpace          int
	description        string
	color              string
	vehicleType        VehicleType
}

func (vehicle *Vehicle) NewVehicle(id int, model string, make string, seater int, bootSpace int, description string, color string, vehicleType VehicleType) {
	vehicle.id = id
	vehicle.model = model
	vehicle.make = make
	vehicle.seater = seater
	vehicle.bootSpace = bootSpace
	vehicle.description = description
	vehicle.color = color
	vehicle.vehicleType = vehicleType
}

func (vehicle *Vehicle) GetId() int {
	return vehicle.id
}

func (vehicle *Vehicle) GetRegistrationNumber() string {
	return vehicle.registrationNumber
}

func (vehicle *Vehicle) GetModel() string {
	return vehicle.model
}

func (vehicle *Vehicle) GetMake() string {
	return vehicle.make
}

func (vehicle *Vehicle) GetSeater() int {
	return vehicle.seater
}

func (vehicle *Vehicle) GetBootSpace() int {
	return vehicle.bootSpace
}

func (vehicle *Vehicle) GetDescription() string {
	return vehicle.description
}

func (vehicle *Vehicle) GetColor() string {
	return vehicle.color
}

func (vehicle *Vehicle) GetVehicleType() VehicleType {
	return vehicle.vehicleType
}

func (vehicle *Vehicle) SetId(id int) *Vehicle {
	vehicle.id = id
	return vehicle
}

func (vehicle *Vehicle) SetRegistrationNumber(registrationNumber string) *Vehicle {
	vehicle.registrationNumber = registrationNumber
	return vehicle
}

func (vehicle *Vehicle) SetModel(model string) *Vehicle {
	vehicle.model = model
	return vehicle
}

func (vehicle *Vehicle) SetMake(make string) *Vehicle {
	vehicle.make = make
	return vehicle
}

func (vehicle *Vehicle) SetSeater(seater int) *Vehicle {
	vehicle.seater = seater
	return vehicle
}

func (vehicle *Vehicle) SetBootSpace(bootSpace int) *Vehicle {
	vehicle.bootSpace = bootSpace
	return vehicle
}

func (vehicle *Vehicle) SetDescription(description string) *Vehicle {
	vehicle.description = description
	return vehicle
}

func (vehicle *Vehicle) SetColor(color string) *Vehicle {
	vehicle.color = color
	return vehicle
}

func (vehicle *Vehicle) SetVehicleType(vehicleType VehicleType) *Vehicle {
	vehicle.vehicleType = vehicleType
	return vehicle
}
