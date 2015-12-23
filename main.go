package main

import (
	"fmt"
	"github.com/arukim/discount-store/store"
)

// More usage examples could be found in store/checkout_test.go
// Basically there are all test cases from code exercise
func main() {
	fmt.Println("Welcome to Cabify store")

	pricingRules := []store.Discounter{
		store.NewDiscounterA("VOUCHER", 2),
		store.NewDiscounterB("TSHIRT", 3, store.Euro(19.00)),
	}

	co := store.NewCheckout(pricingRules)
	co.Scan("VOUCHER")
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	price := co.GetTotal()

	fmt.Printf("Total price is %v\n", price)
}
