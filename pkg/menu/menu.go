package menu

import (
	"fmt"

	"github.com/0xMarvell/auto-shop/pkg/models"
	"github.com/0xMarvell/auto-shop/pkg/utils"
)

// Welcome displays welcome message.
func Welcome() {
	utils.Header("---------------AUTO SHOP------------------")
	fmt.Println("\t      WELCOME!")
}

// Menu displays all available operations to user.
func ShowMenu() {
	utils.Header("----------SELECT OPERATION----------")
	fmt.Println("1. View all auto shop products \t 2. Add a new product")
	fmt.Println("3. View all products for sale \t 4. View all sold products")
	fmt.Println("5. Sell a product")

	var autoShop models.Store

	var selectedOption int
	_, err := fmt.Scan(&selectedOption)
	utils.CheckErr(err)

	switch selectedOption {
	case 1:
		fmt.Println("1 - finish")
	case 2:
		autoShop.AddProduct()
		fmt.Println(autoShop.Products)
	case 3:
		fmt.Println("3 - finish")
	case 4:
		fmt.Println("4  -finish")
	case 5:
		fmt.Println("5 - finish")
	default:
		fmt.Println("ERROR: invalid input, select a valid operation")
		utils.NewLine(1)
		ShowMenu()
	}

}
