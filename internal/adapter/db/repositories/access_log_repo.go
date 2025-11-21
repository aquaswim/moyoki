package repositories

import (
	"context"

	"github.com/aquaswim/moyoki/internal/adapter/db/repositories/model"
	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/aquaswim/moyoki/internal/core/port"
	"gorm.io/gorm"
)

type accessLogRepository struct {
	accessLogDB gorm.Interface[model.AccessLog]
}

func NewAccessLogRepository(
	db *gorm.DB,
) port.AccessLogRepo {
	return &accessLogRepository{
		accessLogDB: gorm.G[model.AccessLog](db),
	}
}

func (a accessLogRepository) Save(ctx context.Context, log *domain.AccessLog) error {
	m := accessLogToModel(log)

	return a.accessLogDB.Create(ctx, m)
}

func accessLogToModel(log *domain.AccessLog) *model.AccessLog {
	return &model.AccessLog{
		ID:         log.ID,
		Path:       log.Path,
		Method:     log.Method,
		RemoteAddr: log.RemoteAddr,
		ReqQuery:   log.ReqQuery,
		ReqBody:    log.ReqBody,
		ReqHeaders: log.ReqHeaders,
		CreatedAt:  log.CreatedAt,
	}
}

func accLogToDomain(acc *model.AccessLog) *domain.AccessLog {
	return &domain.AccessLog{
		ID:         acc.ID,
		Path:       acc.Path,
		Method:     acc.Method,
		RemoteAddr: acc.RemoteAddr,
		ReqQuery:   acc.ReqQuery,
		ReqBody:    acc.ReqBody,
		ReqHeaders: acc.ReqHeaders,
		CreatedAt:  acc.CreatedAt,
	}
}

func (a accessLogRepository) Find(ctx context.Context, param *domain.FindAccessLogParam) ([]domain.AccessLog, error) {
	rows, err := a.accessLogDB.
		Where("created_at > ? AND created_at < ?", param.StartTime, param.EndTime).
		Order("created_at DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	logs := make([]domain.AccessLog, 0, len(rows))
	for i := range rows {
		logs = append(logs, *accLogToDomain(&rows[i]))
	}
	return logs, nil
}
