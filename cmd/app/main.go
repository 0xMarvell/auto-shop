package main

import (
	"fmt"

	"github.com/0xMarvell/auto-shop/pkg/models"
)

func main() {
	// instantiate store object
	var autoShop models.Store

	// add few products to store (both car and non-car)
	engineOil := models.NewProduct("Engine Oil", "nil", 2021, 2026, "", "", 70, 7500)
	camry := models.NewProduct("Toyota", "Camry", 2021, 0, "Auto", "Black", 5, 4000000)
	benz := models.NewProduct("Mercedes Benz", "C300", 2022, 0, "Auto", "Grey", 8, 18000000)
	shockAbsorber := models.NewProduct("Shock Absorber", "nil", 2022, 2030, "", "", 40, 60000)
	tyreRims := models.NewProduct("Tyre Rims", "nil", 2019, 2030, "nil", "Black, Silver", 50, 25000)

	autoShop.AddProduct(engineOil)
	autoShop.AddProduct(camry)
	autoShop.AddProduct(benz)
	autoShop.AddProduct(shockAbsorber)
	autoShop.AddProduct(tyreRims)

	// display all products in store
	autoShop.ListAllProducts()

	// display single product from store (both car and non car)
	engineOil.DisplayProduct()
	benz.DisplayProduct()

	// check if product is in stock
	fmt.Println()
	fmt.Println(tyreRims.InStock())
	fmt.Println(shockAbsorber.InStock())

	// sell product
	autoShop.SellProduct(shockAbsorber, 10)
	autoShop.SellProduct(camry, 1)

	// display all sold items and total price of sold items
	autoShop.ListSoldProducts()
	autoShop.TotalPriceOfSoldProducts()

}
