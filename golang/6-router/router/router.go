package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/onboarding-backend/golang/6-router/config"
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
	v1 := r.Group("/api/v1") // prefix every endpoint with /api/v1 is a good practice

	return &Router{r, v1}
}

func (r *Router) V1() *gin.RouterGroup { // getter for v1 (without this, v1 is not accessible from outside the package)
	return r.v1
}
