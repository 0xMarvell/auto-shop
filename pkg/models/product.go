package models

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/0xMarvell/auto-shop/pkg/utils"
)

type MakeandModel struct {
	Name            string
	Model           string
	ManufactureYear uint
	ExpiryYear      uint
}

type Product struct {
	Id uint
	MakeandModel
	CarDetails
	AvailableUnits uint
	SoldUnits      uint
	Price          uint
}

// DisplayProduct displays all details of a single product
func (p *Product) DisplayProduct() {
	utils.NewLine(1)
	fmt.Println("------------------------------------------------------------DISPLAY PRODUCT--------------------------------------------------------------")
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "Id\tName\tModel\tManufacture Year\tExpiry Year\tTransmission\tColor\tAvailable Units\tPrice")

	var formattedPrice, exp, transmission, color string

	formattedPrice = "₦" + strconv.Itoa(int(p.Price))

	if p.MakeandModel.ExpiryYear == 0 {
		exp = "nil"
	} else {
		exp = strconv.Itoa(int(p.MakeandModel.ExpiryYear))
	}

	if p.CarDetails.Transmission == "" && p.CarDetails.Color == "" {
		transmission = "nil"
		color = "nil"
	} else {
		transmission = p.CarDetails.Transmission
		color = p.CarDetails.Color
	}

	fmt.Fprintln(w, p.Id, "\t", p.MakeandModel.Name, "\t", p.MakeandModel.Model, "\t", p.ManufactureYear, "\t", exp, "\t", transmission, "\t", color, "\t", p.AvailableUnits, "\t", formattedPrice)

	w.Flush()
}

// InStock checks if a product is in stock
func (p *Product) InStock() string {
	if p.AvailableUnits < 1 {
		if p.MakeandModel.Model == "nil" || p.MakeandModel.Model == "" {
			return fmt.Sprintf("Product [%v] is currently OUT OF STOCK\n", p.MakeandModel.Name)
		} else {
			return fmt.Sprintf("Product [%v %v %v] is currently OUT OF STOCK\n", p.MakeandModel.Name, p.MakeandModel.Model, p.ManufactureYear)
		}
	} else {
		if p.MakeandModel.Model == "nil" || p.MakeandModel.Model == "" {
			return fmt.Sprintf("Product [%v] is currently IN STOCK\n", p.MakeandModel.Name)
		} else {
			return fmt.Sprintf("Product [%v %v %v] is currently IN STOCK\n", p.MakeandModel.Name, p.MakeandModel.Model, p.ManufactureYear)
		}
	}
}

// NewProduct instantiates a new product
func NewProduct(
	name string,
	model string,
	prodYear uint,
	expiryYear uint,
	transmission string,
	color string,
	numberOfUnits uint,
	price uint,
) *Product {
	return &Product{
		Id: uint(rand.Intn(10000)),
		MakeandModel: MakeandModel{
			Name:            name,
			Model:           model,
			ManufactureYear: prodYear,
			ExpiryYear:      expiryYear,
		},
		CarDetails: CarDetails{
			Transmission: transmission,
			Color:        color,
		},
		AvailableUnits: numberOfUnits,
		Price:          price,
	}
}
