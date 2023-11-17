package api

// Context is an interface abstracting the context of an HTTP request
type Context interface {
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	Param(key string) string
	Status(code int)
}
