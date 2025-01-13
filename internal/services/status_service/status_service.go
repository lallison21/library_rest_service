package status_service

import "github.com/lallison21/library_rest_service/internal/services"

type Service struct {
	repo services.StatusRepo
}

func New(repo services.StatusRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Ping() error {
	return nil
}
