package models

import (
	"fmt"
	"math/rand"
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
	Price          uint
}

func (p *Product) DisplayProduct() {
	fmt.Println(p)
}

func (p *Product) InStock() bool {
	if p.AvailableUnits < 1 {
		fmt.Println("OUT OF STOCK")
		return false
	} else {
		fmt.Println("IN STOCK")
		return true
	}
}

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
