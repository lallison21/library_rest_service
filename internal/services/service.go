package services

import "github.com/lallison21/library_rest_service/internal/repository"

type Service struct {
	Repository repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{Repository: repo}
}

func (s *Service) Ping() {}
