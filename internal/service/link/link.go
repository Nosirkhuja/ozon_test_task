package link

import (
	"context"
	"database/sql"
	"ozon_test_task/internal/model"
	"ozon_test_task/internal/pkg/token_generator"
)

type Service struct {
	repos Repository
}

func NewService(repos Repository) *Service {
	return &Service{repos: repos}
}

func (s Service) CreateShortURL(ctx context.Context, link *model.Link) (string, error) {
	link.Token = token_generator.GenerateToken()
	token, err := s.repos.CreateShortURL(ctx, link)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) GetBaseURL(ctx context.Context, link *model.Link) (string, error) {
	baseURL, err := s.repos.GetBaseURL(ctx, link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return baseURL, nil
}
