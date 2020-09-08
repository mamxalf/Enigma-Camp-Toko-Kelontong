package main

import (
	"crypto/md5"
	"fmt"
)

type WarehouseRepository struct {
	w *Warehouse
	warehouseList *[]*Warehouse
	warehouseListInfra *WarehouseRepositoryInfrastructure
}

func NewWarehouseRepository(datapath string, warehouseList []*Warehouse) *WarehouseRepository {
	warehouseRepoInfra := NewWarehouseRepoInfra(datapath)
	return &WarehouseRepository{warehouseList: &warehouseList, warehouseListInfra: warehouseRepoInfra}
}

func (wr *WarehouseRepository) AddNewWarehouse(warehouse *Warehouse) {
	data := []byte(warehouse.Name)
	warehouse.Id = fmt.Sprintf("%x", md5.Sum(data))
	*wr.warehouseList = append(*wr.warehouseList, warehouse)
	wr.warehouseListInfra.saveToFileW(wr.warehouseList)
}

func (wr *WarehouseRepository) findAllWarehouse() []*Warehouse {
	wr.warehouseListInfra.readFileW(wr.warehouseList)
	return *wr.warehouseList
}

func (wr *WarehouseRepository) findByName(Name string) float64 {
	//var filter = func(w Warehouse) {
	var price float64
		if Name == wr.w.Name {
			price = wr.w.Price
		}
		return price
	//}
}

func (wr *WarehouseRepository) findWarehouseByField(fnFilter func(Warehouse) bool) []*Warehouse {
	wr.warehouseListInfra.readFileW(wr.warehouseList)
	var searchResult = make([]*Warehouse, 0)
	for _, w := range *wr.warehouseList {
		if fnFilter(*w) {
			searchResult = append(searchResult, w)
		}
	}
	return searchResult
}
