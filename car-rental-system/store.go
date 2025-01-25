package main

type Store struct {
	id                  int
	name                string
	inventoryManagement InventoryManagement
	location            Location
}

func (store *Store) BookVehicle(vehicle Vehicle) {
	store.inventoryManagement.RemoveVehicle(vehicle)
}

func (store *Store) ReturnVehicle(vehicle Vehicle) {
	store.inventoryManagement.AddVehicle(vehicle)
}
