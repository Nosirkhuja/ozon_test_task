package link

import (
	"context"
	"ozon_test_task/internal/model"
)

//go:generate mockgen -package ${GOPACKAGE}_test -source $GOFILE -destination mocks_test.go
type Repository interface {
	CreateShortURL(ctx context.Context, link *model.Link) (string, error)
	GetBaseURL(ctx context.Context, link *model.Link) (string, error)
}
