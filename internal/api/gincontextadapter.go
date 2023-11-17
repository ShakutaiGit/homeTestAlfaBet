package api

import "github.com/gin-gonic/gin"

// GinContextAdapter adapts gin.Context to the custom Context interface
type GinContextAdapter struct {
	C *gin.Context
}

func NewGinContextAdapter(c *gin.Context) Context {
	return &GinContextAdapter{C: c}
}

func (g *GinContextAdapter) BindJSON(obj interface{}) error {
	return g.C.BindJSON(obj)
}

func (g *GinContextAdapter) JSON(code int, obj interface{}) {
	g.C.JSON(code, obj)
}

func (g *GinContextAdapter) Param(key string) string {
	return g.C.Param(key)
}

func (g *GinContextAdapter) Status(code int) {
	g.C.Status(code)
}

func (g *GinContextAdapter) Query(key string) string {
	return g.C.Query(key)
}

// Implement other necessary methods from gin.Context...
