package main

type Location struct {
	address string
	pincode int
	city    string
	state   string
	country string
}

func (location *Location) GetAddress() string {
	return location.address
}

func (location *Location) GetPincode() int {
	return location.pincode
}

func (location *Location) GetCity() string {
	return location.city
}

func (location *Location) GetState() string {
	return location.state
}

func (location *Location) GetCountry() string {
	return location.country
}

func (location *Location) SetAddress(address string) *Location {
	location.address = address
	return location
}

func (location *Location) SetPincode(pincode int) *Location {
	location.pincode = pincode
	return location
}

func (location *Location) SetCity(city string) *Location {
	location.city = city
	return location
}

func (location *Location) SetState(state string) *Location {
	location.state = state
	return location
}

func (location *Location) SetCountry(country string) *Location {
	location.country = country
	return location
}
