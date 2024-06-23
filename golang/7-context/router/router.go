package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/config"
)

type Router struct {
	*gin.Engine
	v1 *gin.RouterGroup // now v1 is private because we created public methods for it in v1.router.go
}

func New(conf *config.Config) *Router {
	if !conf.App.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")

	return &Router{r, v1}
}
