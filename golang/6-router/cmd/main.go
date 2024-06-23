package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/config"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/database"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/internal/model"
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
	r.V1().POST("/user", func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, err)
			return
		}

		userService.AddUser(user.ItemId, user.Quantity)
		c.JSON(200, gin.H{"message": "success"})
	})

	r.V1().GET("/user", func(c *gin.Context) {
		total, apperr := userService.Checkout()
		if apperr != nil {
			c.JSON(500, apperr)
			return
		}

		c.JSON(200, total)
	})

}
