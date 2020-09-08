package main

import (
	"encoding/json"
	"io/ioutil"
)

type WarehouseRepositoryInfrastructure struct {
	datapath string
}

func NewWarehouseRepoInfra(datapath string) *WarehouseRepositoryInfrastructure {
	return &WarehouseRepositoryInfrastructure{datapath}
}

func (sri *WarehouseRepositoryInfrastructure) saveToFileW(warehouseList *[]*Warehouse) {
	file, _ := json.MarshalIndent(warehouseList, "", " ")
	_ = ioutil.WriteFile(sri.datapath, file, 0644)
}

func (sri *WarehouseRepositoryInfrastructure) readFileW(warehouseList *[]*Warehouse) {
	file, _ := ioutil.ReadFile(sri.datapath)
	_ = json.Unmarshal(file, warehouseList)
}