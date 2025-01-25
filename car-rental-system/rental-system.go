package main

type RentalSystem struct {
	userList  []*User
	storeList []*Store
}

func (rs *RentalSystem) GetNearbyStore(loc *Location) *Store {
	//some logic to return store
	return rs.storeList[0]
}
