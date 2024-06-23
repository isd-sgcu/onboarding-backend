package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/onboarding-backend/golang/9-mocks/config"
)

type Router struct {
	*gin.Engine
	v1 *gin.RouterGroup
}

func New(conf *config.Config) *Router {
	if !conf.App.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")

	return &Router{r, v1}
}
