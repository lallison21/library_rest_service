package services

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Ping() {}
