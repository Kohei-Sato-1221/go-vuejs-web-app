package gateway

type UserRepository struct {

}

func (r *UserRepository) GetAllUsers() (string, error) {
	return "[{id:1234, name:sugar}, {id:8321, name:salt}]", nil
}