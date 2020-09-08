package main

import "path"

type AppConfig struct {
	dataPath string
	dbw      []*Warehouse
	dbr		 []*Rental
}

func warehouseConfig() *AppConfig {
	const DATA_PATH = "data"
	const DATA_FILE_NAME = "warehouse.json"

	var warehouseList = make([]*Warehouse, 0)
	return &AppConfig{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		dbw:       warehouseList,
	}
}

func (c *AppConfig) getDBW() []*Warehouse {
	return c.dbw
}

func (c *AppConfig) getDBWPath() string {
	return c.dataPath
}

func rentalConfig() *AppConfig {
	const DATA_PATH = "data"
	const DATA_FILE_NAME = "rental.json"

	var rentalList = make([]*Rental, 0)
	return &AppConfig{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		dbr:       rentalList,
	}
}

func (c *AppConfig) getDBR() []*Rental {
	return c.dbr
}

func (c *AppConfig) getDBRPath() string {
	return c.dataPath
}
