package usecase

import (
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/domain"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/usecase/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
}

func (i *UserInteractor) GetAllUsers() ([]domain.User, error) {
	return i.UserRepository.GetAllUsers()
}

func (i *UserInteractor) CreateUser(user domain.User) (uint, error) {
	return i.UserRepository.CreateUser(user)
}
