package main

import "fmt"

type app struct {
	config      *AppConfig
	warehouseService *WarehouseService
	rentalService *RentalService
}

func (a app) run() {
	NewWarehouseDelivery(a.config).Create()
}

func warehouseApp() app {
	config := warehouseConfig()
	return app{config: config}
}

func (a app) run2() {
	NewRentalDelivery(a.config).Create()
}

func rentalApp() app {
	config := rentalConfig()
	return app{config: config}
}

func start() {
	var kodeNumber string
	fmt.Println("A. Warehouse App")
	fmt.Println("B. Rental App")
	fmt.Printf("Masukan Kode : ")
	fmt.Scanln(&kodeNumber)
	switch {
	case kodeNumber == "A" || kodeNumber == "a":
		warehouseApp().run()
	case kodeNumber == "B" || kodeNumber == "b":
		rentalApp().run2()
	default:
		fmt.Println("Oopss!")
		start()
	}
}
