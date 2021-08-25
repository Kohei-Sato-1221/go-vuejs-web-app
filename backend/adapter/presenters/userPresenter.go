package presenters

import (
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/gateway"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/domain"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/usecase"
	"gorm.io/gorm"
)

type UserPresenter struct {
	Interactor usecase.UserInteractor
}

func NewUserPresenter(db *gorm.DB) *UserPresenter {
	return &UserPresenter{
		Interactor: usecase.UserInteractor{
			UserRepository: &gateway.UserRepository{
				DB: db,
			},
		},
	}
}

func (controller *UserPresenter) GetAllUsers() ([]domain.User, error) {
	users, err := controller.Interactor.GetAllUsers()
	if err == nil {
		return users, err
	} else {
		return nil, err
	}
}

func (controller *UserPresenter) CreateUser(user domain.User) (domain.User, error) {
	_, err := controller.Interactor.CreateUser(user)
	if err == nil {
		return user, nil
	} else {
		return user, err
	}
}
