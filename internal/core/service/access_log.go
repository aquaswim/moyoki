package service

import (
	"context"

	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/aquaswim/moyoki/internal/core/port"
)

type accessLogService struct {
	repo port.AccessLogRepo
}

func NewAccessLogService(
	repo port.AccessLogRepo,
) port.AccessLogService {
	return &accessLogService{
		repo: repo,
	}
}

func (a accessLogService) Insert(ctx context.Context, log *domain.AccessLog) error {
	return a.repo.Save(ctx, log)
}

func (a accessLogService) Find(ctx context.Context, param *domain.FindAccessLogParam) ([]domain.AccessLog, error) {
	return a.repo.Find(ctx, param)
}
