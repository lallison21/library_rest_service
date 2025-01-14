package status_repo

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Ping() (string, error) {
	return "pong", nil
}
