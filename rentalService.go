package main

type RentalService struct {
	r *RentalRepository
}

func NewRentalService(repo *RentalRepository) *RentalService {
	return &RentalService{r: repo}
}

func (rs *RentalService) registerNewRental(r *Rental) {
	rs.r.AddNewRental(r)
}

func (rs *RentalService) getAllRentalList() []*Rental {
	return rs.r.findAllRental()
}

func (rs *RentalService) searchRentalById(id string) []*Rental {
	var funcFilter = func(w Rental) bool {
		return w.Id == id
	}
	return rs.r.findRentalByField(funcFilter)
}

//func (rs *RentalService) searchWareHouseByName{
//
//}