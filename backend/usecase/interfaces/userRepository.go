package interfaces

type UserRepository interface {
	GetAllUsers() (string, error)
}