package controllers

import (
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/gateway"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/interfaces"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/domain"
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

func (controller *UserController) CreateUser(c interfaces.Context) error {
	type (
		Response struct {
			ID uint `json:"id"`
		}
	)
	newUserName := c.QueryParam("name")
	newUserEmail := c.QueryParam("email")
	newUserGender := c.QueryParam("gender")

	if len(newUserName) == 0 ||
		len(newUserEmail) == 0 ||
		len(newUserEmail) == 0 {
		return c.JSON(400, "Error / invalid parameters")
	}

	user := domain.User{
		Name:   newUserName,
		Email:  newUserEmail,
		Gender: newUserGender,
	}

	newUserId, err := controller.Interactor.CreateUser(user)
	if err == nil {
		return c.JSON(200, Response{ID: newUserId})
	} else {
		return c.JSON(400, "Error")
	}
}
