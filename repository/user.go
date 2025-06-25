package repository

type UserRepositoryInterface interface {
	Authenticate(username, password string) bool
	CreateUser(username, password string) error
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]string),
	}
}

var _ UserRepositoryInterface = &InMemoryUserRepository{}

type InMemoryUserRepository struct {
	users map[string]string
}

func (r *InMemoryUserRepository) Authenticate(username, password string) bool {
	return false
}

func (r *InMemoryUserRepository) CreateUser(username, password string) error {
	return nil
}
