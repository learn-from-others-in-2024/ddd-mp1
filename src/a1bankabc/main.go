package main

import (
	"bankapp/transfers"
	"fmt"
	"time"
)

// main is the start point in golang
func main() {

	time := time.Now()
	transfer := transfers.NewTransaction("123", "Educative bank", 23453.98, "USD", time, 453646)

	fmt.Println("**** Before executing a transfer ****")

	// Transfer money, it validates business rules
	if transfer.TransferMoney() {
		fmt.Println("Transfer is done")
	} else {
		fmt.Println("Transfer is not done")
	}

}
