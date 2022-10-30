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
	var autoShop models.Store
	autoShop.AddProduct()

	// camry := &models.Product{
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

	// benz := &models.Product{
	// 	Id: 2,
	// 	MakeandModel: models.MakeandModel{
	// 		Name:            "Mercedes",
	// 		Model:           "Benz",
	// 		ManufactureYear: 2019,
	// 	},
	// 	CarDetails: models.CarDetails{
	// 		Transmission: "Auto",
	// 		Color:        "Blue",
	// 	},
	// 	AvailableUnits: 1,
	// 	Price:          10000000,
	// }
	// camry.DisplayProduct()
	// camry.InStock()

	// autoShop.AddProduct(camry)
	// autoShop.AddProduct(benz)

	fmt.Println(autoShop.Products)

}
