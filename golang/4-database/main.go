package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/4-database/config"
	cartRepo "github.com/isd-sgcu/onboarding-backend/golang/4-database/repository/cart"
	cartSvc "github.com/isd-sgcu/onboarding-backend/golang/4-database/service/cart"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	println("Config loaded successfully: ", conf.App.Port, conf.App.Env, conf.Database.Url, conf.Cors.AllowOrigins)

	cartRepo := cartRepo.NewRepository()
	cartService := cartSvc.NewService(cartRepo)

	cartService.AddOrder(1, 2)
	cartService.AddOrder(2, 3)
	cartService.RemoveOrder(1)
	total := cartService.Checkout()

	println("Total items: ", total)
}
