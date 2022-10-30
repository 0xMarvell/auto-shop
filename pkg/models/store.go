package models

import (
	"bufio"
	"fmt"
	"os"

	"github.com/0xMarvell/auto-shop/pkg/utils"
)

type Store struct {
	Products []Product
}

func (s *Store) ListAllProducts() {

}

func (s *Store) ListProductsForSale() {

}

func (s *Store) ListSoldProducts() {

}

func (s *Store) AddProduct() {
	utils.NewLine(1)
	fmt.Println("--------------ADD NEW PRODUCT----------------")
	utils.NewLine(1)

	var (
		product       *Product
		carOrNotCar   string
		name          string
		model         string
		prodYear      uint
		expiryYear    uint
		transmission  string
		color         string
		numberOfUnits uint
		price         uint
	)

	fmt.Print("Before you proceed, is the product to be added a car? (y/n): ")
	fmt.Scan(&carOrNotCar)
	fmt.Println("---------------------------------------------")

	switch carOrNotCar {
	case "y":
		fmt.Print("Enter car name: ")
		fmt.Scan(&name)

		fmt.Print("Enter car model: ")
		fmt.Scan(&model)

		fmt.Print("Enter car manufacture year: ")
		fmt.Scan(&prodYear)

		fmt.Print("Enter car transmission type: ")
		fmt.Scan(&transmission)

		fmt.Print("Enter car color: ")
		fmt.Scan(&color)

		fmt.Print("Enter number of car units to add to store: ")
		fmt.Scan(&numberOfUnits)

		fmt.Print("Enter car price per unit: ")
		fmt.Scan(&price)

		expiryYear = 0

		product = NewProduct(
			name,
			model,
			prodYear,
			expiryYear,
			transmission,
			color,
			numberOfUnits,
			price,
		)

		s.Products = append(s.Products, *product)
		fmt.Println("New product added successfully")
	case "n":
		fmt.Print("Enter product name: ")
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			name = sc.Text()
		}

		fmt.Print("Enter product model (If product does not have model or you are unsure, enter 'nil'): ")
		fmt.Scan(&model)

		fmt.Print("Enter product manufacture year: ")
		fmt.Scan(&prodYear)

		fmt.Print("Enter product expiry year: ")
		fmt.Scan(&expiryYear)

		fmt.Print("Enter number of product units to add to store: ")
		fmt.Scan(&numberOfUnits)

		fmt.Print("Enter product price per unit: ")
		fmt.Scan(&price)

		transmission = "nil"
		color = "nil"

		product = NewProduct(
			name,
			model,
			prodYear,
			expiryYear,
			transmission,
			color,
			numberOfUnits,
			price,
		)

		s.Products = append(s.Products, *product)
		fmt.Println("New product added successfully")
	default:
		fmt.Println("error: invalid input")
	}

	// s.Products = append(s.Products, *product)
}

func (s *Store) SellProduct() {

}
