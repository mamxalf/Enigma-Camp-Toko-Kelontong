package main

type Rental struct {
	Id string
	PWarehouse string
	Name string
	RentalType string
	Large	float64
	DateIn string
	DateOut string
	Price float64
}

func NewRental(pWarehouse string, name string, rentalType string, large float64, dateIn string, dateOut string, price float64) Rental {
	return Rental{
		PWarehouse: pWarehouse,
		Name: name,
		RentalType: rentalType,
		Large: large,
		DateIn: dateIn,
		DateOut: dateOut,
		Price: price,
	}
}
