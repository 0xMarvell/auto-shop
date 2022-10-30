package models

import (
	"fmt"
)

type MakeandModel struct {
	Name            string
	Model           string
	ManufactureYear uint
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
