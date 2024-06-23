package main

import (
	"fmt"

	"github.com/isd-sgcu/onboarding-backend/golang/6-router/config"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/database"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/dto"
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
	userService := user.NewService(userRepo)

	r := router.New(conf)
	r.V1Post("/user", func(c router.Context) { // use V1Post() instead of V1.POST() and router.Context instead of *gin.Context
		var createUserDto dto.CreaterUserRequest
		if err := c.Bind(&createUserDto); err != nil {
			c.JSON(400, err)
			return
		}

		createdUser, err := userService.Create(&createUserDto)
		if err != nil {
			c.JSON(500, err)
			return
		}

		c.JSON(200, createdUser)
	})

	r.V1Get("/user/:id", func(c router.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id is required in url param")
			return
		}

		total, apperr := userService.FindOne(&dto.FindOneUserRequest{Id: id})
		if apperr != nil {
			c.JSON(500, apperr)
			return
		}

		c.JSON(200, total)
	})

	r.V1Delete("/user/:id", func(c router.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id is required in url param")
			return
		}

		total, apperr := userService.Delete(&dto.DeleteUserRequest{Id: id})
		if apperr != nil {
			c.JSON(500, apperr)
			return
		}

		c.JSON(200, total)
	})

	if err := r.Run(fmt.Sprintf(":%v", conf.App.Port)); err != nil {
		panic(fmt.Sprintf("Failed to run router: %v", err))
	}
}
