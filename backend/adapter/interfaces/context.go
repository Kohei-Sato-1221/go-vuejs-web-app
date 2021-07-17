package interfaces

type Context interface {
	JSON(int, interface{}) error
}
