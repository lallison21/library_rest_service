package services

type StatusRepo interface {
	Ping() (string, error)
}

type AuthRepo interface {
}
