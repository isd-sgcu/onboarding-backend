package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/1-struct-interface/service"
)

func main() {
	// function main() in package main is 99.999% the entry point of the program

	// create a new instance of cart that implements the Cart interface
	cartOne := service.NewEmptyCart()
	// add an order to the cart
	cartOne.AddOrder(1, 2)

	// create a new instance of cart that implements the Cart interface
	orders := map[int]int{
		1: 2,
		2: 3,
	}
	cartTwo := service.NewCartWithOrders(orders)
	// remove an order from the cart
	cartTwo.RemoveOrder(1)

	// checkout the cart
	totalOne := cartOne.Checkout()
	println("Total items one:", totalOne)

	totalTwo := cartTwo.Checkout()
	println(fmt.Sprintf("Total items two: %d", totalTwo)) // fmt.Sprintf() is similar to String.format() in Java

}
