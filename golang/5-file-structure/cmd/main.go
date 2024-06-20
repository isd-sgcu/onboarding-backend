package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/config"
	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/database"
	cartRepo "github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/repository/cart"
	cartSvc "github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/service/cart"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	db, err := database.InitPostgresDatabase(&conf.Database, conf.App.IsDevelopment())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	cartRepo := cartRepo.NewRepository(db)
	cartService := cartSvc.NewService(cartRepo)

	cartService.AddOrder(1, 2)
	cartService.AddOrder(2, 3)

	total, err := cartService.Checkout()
	if err != nil {
		panic(fmt.Sprintf("Failed to checkout: %v", err))
	}

	println("Total items: ")
	for _, order := range *total {
		println(fmt.Sprintf("Item %d: %d", order.ItemId, order.Quantity))
	}
}
