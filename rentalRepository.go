package main

import (
	"crypto/md5"
	"fmt"
)

type RentalRepository struct {
	rentalList *[]*Rental
	rentalListInfra *RentalRepositoryInfrastructure
	warehouseList []*Warehouse
	warehouseRepo *WarehouseRepository
	//warehouseList *WarehouseRepositoryInfrastructure

}

func NewRentalRepository(datapath string, rentalList []*Rental) *RentalRepository {
	rentalRepoInfra := NewRentalRepoInfra(datapath)
	return &RentalRepository{rentalList: &rentalList, rentalListInfra: rentalRepoInfra}
}

func (rr *RentalRepository) AddNewRental(rental *Rental) {
	data := []byte(rental.Name)
	rental.Id = fmt.Sprintf("%x", md5.Sum(data))

	//var price float64
	//for _, j := range rr.warehouseRepo.findAllWarehouse(){
	//	if rental.PWarehouse == j.Name {
	//		price = j.Price
	//	}
	//}

	//rental.Price = rental.Large * price
	rental.Price = rental.Large * 10000
	*rr.rentalList = append(*rr.rentalList, rental)
	rr.rentalListInfra.saveToFileR(rr.rentalList)
}

func (rr *RentalRepository) findAllRental() []*Rental {
	rr.rentalListInfra.readFileW(rr.rentalList)
	return *rr.rentalList
}

func (rr *RentalRepository) findRentalByField(fnFilter func(Rental) bool) []*Rental {
	rr.rentalListInfra.readFileW(rr.rentalList)
	var searchResult = make([]*Rental, 0)
	for _, w := range *rr.rentalList {
		if fnFilter(*w) {
			searchResult = append(searchResult, w)
		}
	}
	return searchResult
}
