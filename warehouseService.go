package main

type WarehouseService struct {
	w *WarehouseRepository
}

func NewWarehouseService(repo *WarehouseRepository) *WarehouseService {
	return &WarehouseService{w: repo}
}

func (ws *WarehouseService) registerNewWarehouse(w *Warehouse) {
	ws.w.AddNewWarehouse(w)
}

func (ws *WarehouseService) getAllWarehouseList() []*Warehouse {
	return ws.w.findAllWarehouse()
}

func (ws *WarehouseService) searchWarehouseById(id string) []*Warehouse {
	var funcFilter = func(w Warehouse) bool {
		return w.Id == id
	}
	return ws.w.findWarehouseByField(funcFilter)
}

//func (ws *WarehouseService) searchWareHouseByName{
//
//}