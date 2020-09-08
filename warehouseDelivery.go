package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WarehouseDelivery struct {
	warehouseService *WarehouseService
}

func NewWarehouseDelivery(ac *AppConfig) *WarehouseDelivery {
	repo := NewWarehouseRepository(ac.getDBWPath(), ac.getDBW())
	warehouseService := NewWarehouseService(repo)
	return &WarehouseDelivery{warehouseService}
}

func (wd *WarehouseDelivery) Create() {
	var isExist = false
	var userChoice string

	wd.mainMenuForm()
	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			wd.registrationWarehouseForm()
		case userChoice == "02":
			wd.viewWarehouseList()
		case userChoice == "q":
			start()
		default:
			fmt.Println("Unknown Menu Code")

		}
	}
}

func (wd *WarehouseDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (wd *WarehouseDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Add Warehouse",
		"02": "View list Warehouse",
		"q":  "Back aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Warehouse Bro !")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range wd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}

func (wd *WarehouseDelivery) registrationWarehouseForm() {
	consoleClear()
	//var Id string
	var name string
	var address string
	var large float64
	var info string
	var price float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Add your New Warehouse !")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Name : ")
	sName, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(sName)
	fmt.Print("Address : ")
	sAddress, _ := scanner.ReadString('\n')
	address = strings.TrimSpace(sAddress)
	fmt.Print("Large /m : ")
	sLarge, _ := scanner.ReadString('\n')
	large, _ = strconv.ParseFloat(strings.TrimSpace(sLarge), 64)
	fmt.Print("Info : ")
	sInfo, _ := scanner.ReadString('\n')
	info = strings.TrimSpace(sInfo)
	fmt.Print("Price : ")
	sPrice, _ := scanner.ReadString('\n')
	price, _ = strconv.ParseFloat(strings.TrimSpace(sPrice), 64)

	fmt.Println("Save to collection? :Y/n")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" || confirmation == "y" {
		newWarehouse := NewWarehouse(name, address, large, info, price)
		wd.warehouseService.registerNewWarehouse(&newWarehouse)
	}
	consoleClear()
	wd.mainMenuForm()
}

func (wd *WarehouseDelivery) viewWarehouseList() {
	consoleClear()
	warehouses := wd.warehouseService.getAllWarehouseList()
	fmt.Println("")
	fmt.Println("My Warehouse")
	fmt.Printf("%s\n", strings.Repeat("=", 140))
	fmt.Printf("%-40s%-25s%-25s%-15s%-15s%-15s\n", "ID", "Name", "Address", "Large /m", "Info", "Price (Rp.)")
	fmt.Printf("%s\n", strings.Repeat("-", 140))
	for _, w := range warehouses {
		fmt.Printf("%-40s%-25s%-25s%-15.2f%-15s%-15.2f\n", w.Id, w.Name, w.Address, w.Large, w.Info, w.Price)
	}
	var confirmation string
	fmt.Print("Back To Menu ? : (Y)")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" || confirmation == "y" {
		wd.mainMenuForm()
	} else {
		wd.viewWarehouseList()
	}
}