package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/config"
)

type Router struct {
	*gin.Engine
	V1 *gin.RouterGroup // capital V1 to make it public
}

func New(conf *config.Config) *Router {
	if !conf.App.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1") // prefix every endpoint with /api/v1 is a good practice

	return &Router{r, v1}
}
