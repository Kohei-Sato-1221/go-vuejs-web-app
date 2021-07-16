package controllers

import "github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/usecase"

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController() *UserController {
	return &UserController{
		Interactor: *usecase.NewUserInteractor(),
	}
}

func (c *UserController) GetAllUsers() (string, error) {
	return c.Interactor.GetAllUsers()
}