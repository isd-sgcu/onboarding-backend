package router

import "github.com/gin-gonic/gin"

// we wrap the gin's original GET with our own GET method becuase we want to use a custom Context interface
func (r *Router) V1Get(path string, handler func(c Context)) {
	r.v1.GET(path, func(c *gin.Context) { // this is the gin's original GET method. It takes func(c *gin.Context){}
		handler(NewContext(c)) // NewContext(c) transforms gin's c *gin.Context to our custom Context interface
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
