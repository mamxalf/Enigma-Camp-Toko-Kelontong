package main

type Warehouse struct {
	Id string
	Name string
	Address string
	Large	float64
	Info string
	Price float64
}

func NewWarehouse(name string, address string, large float64, info string, price float64) Warehouse {
	return Warehouse{
		Name: name,
		Address: address,
		Large: large,
		Info: info,
		Price: price,
	}
}
