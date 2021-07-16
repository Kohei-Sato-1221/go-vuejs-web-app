package usecase

import "github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/usecase/interfaces"

type UserInteractor struct {
	UserRepository interfaces.UserRepository
}

func (i *UserInteractor) GetAllUsers() (string, error) {
	return i.UserRepository.GetAllUsers()
}