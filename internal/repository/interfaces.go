package repository

type Repository interface {
	Ping() error
}
