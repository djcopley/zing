package repository

type SessionRepositoryInterface interface {
	Create(username string, token string) error
	Read(username string) (string, error)
}

type InMemorySessionRepository struct {
	// username to session token
	sessions map[string]string
}

var _ SessionRepositoryInterface = &InMemorySessionRepository{}

func NewInMemorySessionRepository() *InMemorySessionRepository {
	return &InMemorySessionRepository{
		sessions: make(map[string]string),
	}
}

func (r *InMemorySessionRepository) Create(username, token string) error {
	r.sessions[username] = token
	return nil
}

func (r *InMemorySessionRepository) Read(username string) (string, error) {
	return r.sessions[username], nil
}
