package services

import "fmt"

type StatusService struct {
	repo StatusRepo
}

func NewStatus(repo StatusRepo) *StatusService {
	return &StatusService{
		repo: repo,
	}
}

func (s *StatusService) Ping() (string, error) {
	ping, err := s.repo.Ping()
	if err != nil {
		return "", fmt.Errorf("[Ping] ping service: %w", err)
	}

	return ping, nil
}
