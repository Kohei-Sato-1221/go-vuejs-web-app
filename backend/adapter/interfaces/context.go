package interfaces

//ref:https://github.com/labstack/echo/blob/master/context.go
type Context interface {
	JSON(int, interface{}) error
	QueryParam(string) string
}
