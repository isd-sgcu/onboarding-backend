package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/5-architecture/config"
	"github.com/isd-sgcu/onboarding-backend/golang/5-architecture/database"
	"github.com/isd-sgcu/onboarding-backend/golang/5-architecture/internal/dto"
	"github.com/isd-sgcu/onboarding-backend/golang/5-architecture/internal/user"
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

	createNewUserReq := &dto.CreaterUserRequest{
		Email:    "john@gmail.com",
		Password: "1234",
	}

	res, apperr := userService.Create(createNewUserReq)
	if apperr != nil {
		panic(fmt.Sprintf("Failed to create new user: %v", apperr))
	}
	println(fmt.Sprintf("Created user: %v", res.User))
}
