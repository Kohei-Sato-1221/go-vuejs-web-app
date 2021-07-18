package interfaces

import "github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/domain"

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
}
