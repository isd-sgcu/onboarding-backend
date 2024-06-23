package router

import "github.com/gin-gonic/gin"

func (r *Router) V1Get(path string, handler func(c Context)) {
	r.v1.GET(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) V1Post(path string, handler func(c Context)) {
	r.v1.POST(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) V1Delete(path string, handler func(c Context)) {
	r.v1.DELETE(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}
