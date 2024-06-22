package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/config"
	"github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/database"
	userRepo "github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/repository/user"
	userSvc "github.com/isd-sgcu/onboarding-backend/golang/5-file-structure/internal/service/user"
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

	userRepo := userRepo.NewRepository(db)
	userService := userSvc.NewService(userRepo)

	userService.AddUser(1, 2)
	userService.AddUser(2, 3)

	total, err := userService.Checkout()
	if err != nil {
		panic(fmt.Sprintf("Failed to checkout: %v", err))
	}

	println("Total items: ")
	for _, user := range *total {
		println(fmt.Sprintf("Item %d: %d", user.ItemId, user.Quantity))
	}
}
