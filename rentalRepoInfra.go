package main

import (
	"encoding/json"
	"io/ioutil"
)

type RentalRepositoryInfrastructure struct {
	datapath string
}

func NewRentalRepoInfra(datapath string) *RentalRepositoryInfrastructure {
	return &RentalRepositoryInfrastructure{datapath}
}

func (rri *RentalRepositoryInfrastructure) saveToFileR(RentalList *[]*Rental) {
	file, _ := json.MarshalIndent(RentalList, "", " ")
	_ = ioutil.WriteFile(rri.datapath, file, 0644)
}

func (rri *RentalRepositoryInfrastructure) readFileW(RentalList *[]*Rental) {
	file, _ := ioutil.ReadFile(rri.datapath)
	_ = json.Unmarshal(file, RentalList)
}