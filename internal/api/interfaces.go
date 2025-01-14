package api

type StatusService interface {
	Ping() (string, error)
}

type AuthService interface {
}
