package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type RentalDelivery struct {
	rentalService *RentalService
}

func NewRentalDelivery(ac *AppConfig) *RentalDelivery {
	repo := NewRentalRepository(ac.getDBRPath(), ac.getDBR())
	rentalService := NewRentalService(repo)
	return &RentalDelivery{rentalService}
}

func (rd *RentalDelivery) Create() {
	var isExist = false
	var userChoice string

	rd.mainMenuForm()
	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			rd.registrationRentalForm()
		case userChoice == "02":
			rd.viewRentalList()
		case userChoice == "q":
			start()
		default:
			fmt.Println("Unknown Menu Code")

		}
	}
}

func (rd *RentalDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (rd *RentalDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Add Rental",
		"02": "View list Rental",
		"q":  "Back aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Rental Bro !")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range rd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}

func (rd *RentalDelivery) registrationRentalForm() {
	consoleClear()
	//var Id string
	//var Id string
	var PWarehouse string
	var Name string
	var RentalType string
	var Large	float64
	var DateIn string
	var DateOut string
	var Price float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Add your New Rental !")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Choose Warehouse : ")
	sPWarehouse, _ := scanner.ReadString('\n')
	PWarehouse = strings.TrimSpace(sPWarehouse)
	fmt.Print("Name Tenant : ")
	sName, _ := scanner.ReadString('\n')
	Name = strings.TrimSpace(sName)
	fmt.Print("Rental Type : ")
	sRentalType, _ := scanner.ReadString('\n')
	RentalType = strings.TrimSpace(sRentalType)
	fmt.Print("Large /m : ")
	sLarge, _ := scanner.ReadString('\n')
	Large, _ = strconv.ParseFloat(strings.TrimSpace(sLarge), 64)
	fmt.Print("Date In : ")
	sDateIn, _ := scanner.ReadString('\n')
	DateIn = strings.TrimSpace(sDateIn)
	fmt.Print("Date Out : ")
	sDateOut, _ := scanner.ReadString('\n')
	DateOut = strings.TrimSpace(sDateOut)
	//fmt.Print("Price : ")
	//sPrice, _ := scanner.ReadString('\n')
	//Price, _ = strconv.ParseFloat(strings.TrimSpace(sPrice), 64)
	Price = 0
	fmt.Println("Save to collection? :Y/n")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" || confirmation == "y" {
		newRental := NewRental(PWarehouse, Name, RentalType, Large, DateIn, DateOut, Price)
		rd.rentalService.registerNewRental(&newRental)
	}
	consoleClear()
	rd.mainMenuForm()
}

func (rd *RentalDelivery) viewRentalList() {
	consoleClear()
	rentals := rd.rentalService.getAllRentalList()
	fmt.Println("")
	fmt.Println("My Rental")
	fmt.Printf("%s\n", strings.Repeat("=", 180))
	fmt.Printf("%-40s%-25s%-25s%-20s%-15s%-25s%-25s%-15s\n", "ID", "Gudang", "Name", "Rental Type", "Large /m", "Date IN", "Date OUT", "Price (Rp.)")
	fmt.Printf("%s\n", strings.Repeat("-", 180))
	for _, w := range rentals {
		fmt.Printf("%-40s%-25s%-25s%-20s%-15.2f%-25s%-25s%-15.2f\n", w.Id, w.PWarehouse, w.Name, w.RentalType, w.Large, w.DateIn, w.DateOut, w.Price)
	}
	var confirmation string
	fmt.Print("Back To Menu ? : (Y)")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" || confirmation == "y" {
		rd.mainMenuForm()
	} else {
		rd.viewRentalList()
	}
}