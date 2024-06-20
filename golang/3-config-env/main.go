package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/3-config-env/config"
	cartRepo "github.com/isd-sgcu/onboarding-backend/golang/3-config-env/repository/cart"
	cartSvc "github.com/isd-sgcu/onboarding-backend/golang/3-config-env/service/cart"
)

func main() {
	// we can load every envs here, but it would be a lot of code. It's better to put it in a function
	// In main.go, we only need to know that we load the config here, and whether it's successful or not
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	println("Config loaded successfully: ", conf.App.Port, conf.App.Env, conf.Database.Url, conf.Cors.AllowOrigins)
	println("Is development: ", conf.App.IsDevelopment())

	cartRepo := cartRepo.NewRepository()
	cartService := cartSvc.NewService(cartRepo)

	cartService.AddOrder(1, 2)
	cartService.AddOrder(2, 3)
	cartService.RemoveOrder(1)
	total := cartService.Checkout()

	println("Total items: ", total)
}
