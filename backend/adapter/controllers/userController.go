package controllers

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetAllUsers() string{
	return "[id:1234, name:sugar]"
}