package main

import (
	cat "a2products/catalog"
	"fmt"
)

// main is the start point
func main() {

	sells := cat.NewSells("123", "987", cat.NewProductRepositoryImplementation(), cat.NewCustomerRepositoryImplementation())

	fmt.Println("Information about a bill: ")
	sells.GenerateBill()

}
