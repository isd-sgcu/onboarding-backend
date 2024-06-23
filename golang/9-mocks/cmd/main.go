package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/6-router/config"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/database"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/user"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/router"
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
	userSvc := user.NewService(userRepo)
	userHdr := user.NewHandler(userSvc)

	r := router.New(conf)
	r.V1Post("/user", userHdr.Create)

	r.V1Get("/user/:id", userHdr.FindOne)

	r.V1Delete("/user/:id", userHdr.Delete)

	if err := r.Run(fmt.Sprintf(":%v", conf.App.Port)); err != nil {
		panic(fmt.Sprintf("Failed to run router: %v", err))
	}
}