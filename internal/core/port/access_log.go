package port

import (
	"context"

	"github.com/aquaswim/moyoki/internal/core/domain"
)

type AccessLogService interface {
	Find(ctx context.Context, param *domain.FindAccessLogParam) ([]domain.AccessLog, error)
	Insert(ctx context.Context, log *domain.AccessLog) error
}

type AccessLogRepo interface {
	Find(ctx context.Context, param *domain.FindAccessLogParam) ([]domain.AccessLog, error)
	Save(ctx context.Context, log *domain.AccessLog) error
}
