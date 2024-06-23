package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/6-router/config"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/database"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/user"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	db, err := database.InitDatabase(&conf.Database, conf.App.IsDevelopment())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)

	userService.AddUser(1, 2)
	userService.AddUser(2, 3)

	total, apperr := userService.Checkout()
	if apperr != nil {
		panic(fmt.Sprintf("Failed to checkout: %v", err))
	}

	println("Total items: ")
	for _, user := range *total {
		println(fmt.Sprintf("Item %d: %d", user.ItemId, user.Quantity))
	}
}
