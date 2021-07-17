package gateway

import (
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/domain"
)

type UserRepository struct {
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User

	user1 := domain.User{
		ID:     1234,
		Name:   "sugar",
		Email:  "hogehoge@sugar.co.jp",
		Gender: "Male",
	}

	user2 := domain.User{
		ID:     1235,
		Name:   "salt",
		Email:  "hogehoge@salt.co.jp",
		Gender: "Female",
	}

	users = append(users, user1)
	users = append(users, user2)

	return users, nil
}
