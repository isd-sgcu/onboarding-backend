package main

import (
	// We must rename the packages because they use the same name
	cartRepo "github.com/isd-sgcu/onboarding-backend/golang/2-dependency-injection/repository/cart"
	cartSvc "github.com/isd-sgcu/onboarding-backend/golang/2-dependency-injection/service/cart"
)

func main() {
	cartRepo := cartRepo.NewRepository()
	cartService := cartSvc.NewService(cartRepo)

	cartService.AddOrder(1, 2)
	cartService.AddOrder(2, 3)
	cartService.RemoveOrder(1)
	total := cartService.Checkout()

	println("Total items: ", total)

}
