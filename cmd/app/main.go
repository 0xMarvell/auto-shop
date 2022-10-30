package main

import (
	"fmt"

	"github.com/0xMarvell/auto-shop/pkg/models"
)

// func containsProduct(s []models.Product, p *models.Product) bool {
// 	for _, v := range s {
// 		sameProduct := (v.MakeandModel.Model == p.MakeandModel.Model) && (v.MakeandModel.Name == p.MakeandModel.Name) && (v.MakeandModel.ManufactureYear == p.MakeandModel.ManufactureYear)
// 		if sameProduct {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	fmt.Println("hello")
	var cars []models.Product

	camry := &models.Product{
		Id: 1,
		MakeandModel: models.MakeandModel{
			Name:            "Toyota",
			Model:           "Camry",
			ManufactureYear: 2019,
		},
		CarDetails: models.CarDetails{
			Transmission: "Auto",
			Color:        "Red",
		},
		AvailableUnits: 3,
		Price:          5000000,
	}

	// camry2 := &models.Product{
	// 	Id: 1,
	// 	MakeandModel: models.MakeandModel{
	// 		Name:            "Toyota",
	// 		Model:           "Camry",
	// 		ManufactureYear: 2019,
	// 	},
	// 	CarDetails: models.CarDetails{
	// 		Transmission: "Auto",
	// 		Color:        "Red",
	// 	},
	// 	AvailableUnits: 3,
	// 	Price:          5000000,
	// }

	benz := &models.Product{
		Id: 2,
		MakeandModel: models.MakeandModel{
			Name:            "Mercedes",
			Model:           "Benz",
			ManufactureYear: 2019,
		},
		CarDetails: models.CarDetails{
			Transmission: "Auto",
			Color:        "Blue",
		},
		AvailableUnits: 1,
		Price:          10000000,
	}

	// if camry2.MakeandModel.ManufactureYear && camry2.MakeandModel.Model && camry2.MakeandModel.Name{

	// }

	cars = append(cars, *camry, *benz)
	fmt.Println(cars)

	// fmt.Println(containsProduct(cars, camry2))

	// camry.DisplayProduct()
	// camry.InStock()
	// fmt.Println(camry.MakeandModel.Model)
}
