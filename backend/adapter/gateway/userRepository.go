package gateway

import (
	"fmt"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/domain"
	"gorm.io/gorm"
)

//ref:https://gorm.io/ja_JP/docs/query.html
type (
	UserRepository struct {
		DB *gorm.DB
	}

	User struct {
		ID     uint   `gorm:"primaryKey"`
		Name   string `gorm:"size:100"`
		Email  string `gorm:"size:100"`
		Gender string `gorm:"size:30"`
		gorm.Model
	}
)

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []User
	result := r.DB.Find(&users)
	if result.Error != nil {
		fmt.Println("ERROR!")
		fmt.Println(result.Error)
		return nil, result.Error
	}

	var domainUsers []domain.User
	for i := 0; i < len(users); i++ {
		user := users[i]
		domainUser := domain.User{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Gender: user.Gender,
		}
		domainUsers = append(domainUsers, domainUser)
	}

	return domainUsers, nil
}

func (r *UserRepository) CreateUser(user domain.User) (uint, error) {
	repoUser := &User{
		Name:   user.Name,
		Email:  user.Email,
		Gender: user.Gender,
	}

	result := r.DB.Create(&repoUser)
	if result.Error != nil {
		fmt.Println("ERROR!")
		fmt.Println(result.Error)
		return 0, result.Error
	}

	return repoUser.ID, nil
}
