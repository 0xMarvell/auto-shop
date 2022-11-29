package models

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/0xMarvell/auto-shop/pkg/utils"
)

type Store struct {
	Products     []Product
	SoldProducts []Product
}

// ListAllProducts displays all products in the store
func (s *Store) ListAllProducts() {
	utils.NewLine(1)
	fmt.Println("------------------------------------------------------------ALL PRODUCTS--------------------------------------------------------------")
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Id\tName\tModel\tManufacture Year\tExpiry Year\tTransmission\tColor\tAvailable Units\tPrice\t")

	for _, v := range s.Products {
		var formattedPrice, exp, transmission, color string

		formattedPrice = "₦" + strconv.Itoa(int(v.Price))

		if v.MakeandModel.ExpiryYear == 0 {
			exp = "nil"
		} else {
			exp = strconv.Itoa(int(v.MakeandModel.ExpiryYear))
		}

		if v.CarDetails.Transmission == "" && v.CarDetails.Color == "" {
			transmission = "nil"
			color = "nil"
		} else {
			transmission = v.CarDetails.Transmission
			color = v.CarDetails.Color
		}

		fmt.Fprintln(w, v.Id, "\t", v.MakeandModel.Name, "\t", v.MakeandModel.Model, "\t", v.MakeandModel.ManufactureYear, "\t", exp, "\t", transmission, "\t", color, "\t", v.AvailableUnits, "\t", formattedPrice, "\t")
	}

	w.Flush()
}

// ListSoldProducts displays all sold products
func (s *Store) ListSoldProducts() {
	utils.NewLine(1)
	fmt.Println("------------------------------------------------------------SOLD PRODUCTS----------------------------------------------------------------")
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Id\tName\tModel\tManufacture Year\tExpiry Year\tTransmission\tColor\tSold Units\tPrice\t")

	for _, v := range s.SoldProducts {
		var formattedPrice, exp, transmission, color string

		formattedPrice = "₦" + strconv.Itoa(int(v.Price))

		if v.MakeandModel.ExpiryYear == 0 {
			exp = "nil"
		} else {
			exp = strconv.Itoa(int(v.MakeandModel.ExpiryYear))
		}

		if v.CarDetails.Transmission == "" && v.CarDetails.Color == "" {
			transmission = "nil"
			color = "nil"
		} else {
			transmission = v.CarDetails.Transmission
			color = v.CarDetails.Color
		}

		fmt.Fprintln(w, v.Id, "\t", v.MakeandModel.Name, "\t", v.MakeandModel.Model, "\t", v.MakeandModel.ManufactureYear, "\t", exp, "\t", transmission, "\t", color, "\t", v.SoldUnits, "\t", formattedPrice, "\t")
	}

	w.Flush()
}

// AddProduct adds a new product to the store
func (s *Store) AddProduct(product *Product) {
	utils.NewLine(1)
	fmt.Println("-------------------ADD NEW PRODUCT-------------------")
	s.Products = append(s.Products, *product)
	if product.MakeandModel.Model == "nil" || product.MakeandModel.Model == "" {
		fmt.Printf("Successfully added new product => %v [%v units]\n", product.MakeandModel.Name, product.AvailableUnits)
	} else {
		fmt.Printf("Successfully added new product => %v %v %v [%v units]\n", product.MakeandModel.Name, product.MakeandModel.Model, product.ManufactureYear, product.AvailableUnits)
	}

}

// SellProduct sells a product from the store
func (s *Store) SellProduct(product *Product, numberofUnits uint) {
	utils.NewLine(1)
	fmt.Println("-------------------SELL PRODUCT-------------------")
	if product.AvailableUnits > 0 {
		product.AvailableUnits -= numberofUnits
		product.SoldUnits += numberofUnits
		s.SoldProducts = append(s.SoldProducts, *product)

	} else {
		fmt.Println("ERROR: PRODUCT IS CURRENTLY OUT OF STOCK")
	}

	if product.MakeandModel.Model == "nil" || product.MakeandModel.Model == "" {
		fmt.Printf("Successfully sold %v units of Product [%v]\n", numberofUnits, product.MakeandModel.Name)
	} else {
		fmt.Printf("Successfully sold %v units of Product [%v %v %v]\n", numberofUnits, product.MakeandModel.Name, product.MakeandModel.Model, product.MakeandModel.ManufactureYear)
	}

}

// TotalPriceOfSoldProducts displays the total price of all sold products
func (s *Store) TotalPriceOfSoldProducts() {
	utils.NewLine(1)
	var totalPriceSoldProducts uint
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)

	fmt.Println("-------------------TOTAL PRICE OF SOLD PRODUCTS-------------------")
	fmt.Fprintln(w, "Id\tProduct\tPrice (single unit)\tSold Units\t")

	for _, v := range s.SoldProducts {
		var product, formattedPrice string
		totalPriceSoldProducts += (v.Price * v.SoldUnits)
		formattedPrice = "₦" + strconv.Itoa(int(v.Price))
		if v.MakeandModel.Model == "nil" || v.MakeandModel.Model == "" {
			product = v.MakeandModel.Name
		} else {
			product = v.MakeandModel.Name + v.MakeandModel.Model + strconv.Itoa(int(v.MakeandModel.ManufactureYear))
		}
		fmt.Fprintln(w, v.Id, "\t", product, "\t", formattedPrice, "\t", v.SoldUnits, "\t")
	}

	w.Flush()
	utils.NewLine(1)
	fmt.Println("TOTAL PRICE OF SOLD PRODUCTS =>", "₦"+strconv.Itoa(int(totalPriceSoldProducts)))
}
