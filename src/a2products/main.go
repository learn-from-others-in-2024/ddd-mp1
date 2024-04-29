package main

import (
	cat "a2products/catalog"
	"fmt"
)

// main is the start point
func main() {

	sells := cat.NewSells(cat.NewProductRepositoryImplementation())
	products := sells.ListProducts()

	fmt.Println("List of products:\n")

	for _, v := range products {
		fmt.Printf("%s - %s - %f\n \n", v.Id, v.Name, v.Price)
	}
}
