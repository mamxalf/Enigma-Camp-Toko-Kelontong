package main

type RentalRepository struct {
	rentalList *[]*Rental
	rentalListInfra *RentalRepositoryInfrastructure
	warehouseList Warehouse
	warehouseRepo *WarehouseRepository

}

func NewRentalRepository(datapath string, rentalList []*Rental) *RentalRepository {
	rentalRepoInfra := NewRentalRepoInfra(datapath)
	return &RentalRepository{rentalList: &rentalList, rentalListInfra: rentalRepoInfra}
}

func (rr *RentalRepository) AddNewRental(rental *Rental) {

	rental.Price = 9000
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
