package controllers

import (
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/gateway"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/interfaces"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/usecase"
	"gorm.io/gorm"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &gateway.UserRepository{
				DB: db,
			},
		},
	}
}

func (controller *UserController) GetAllUsers(c interfaces.Context) error {
	users, err := controller.Interactor.GetAllUsers()
	if err == nil {
		return c.JSON(200, users)
	} else {
		return c.JSON(400, "Error")
	}
}
