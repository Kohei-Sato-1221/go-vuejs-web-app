package usecase

type UserInteractor struct {

}

func NewUserInteractor() *UserInteractor {
	return &UserInteractor{}
}

func (i *UserInteractor) GetAllUsers() (string, error) {
	return "[{id:1234, name:sugar}, {id:8321, name:salt}]", nil
}