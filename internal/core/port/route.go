package port

import (
	"context"
	"github.com/aquaswim/moyoki/internal/core/domain"
)

type RouteRepository interface {
	Find(ctx context.Context) ([]domain.RouteItem, error)
	FindByID(ctx context.Context, id int) (*domain.RouteItem, error)
	Save(ctx context.Context, route *domain.RouteItem) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, route *domain.RouteItem) error
}

type RouteService interface {
	Find(ctx context.Context) ([]domain.RouteItem, error)
	FindByID(ctx context.Context, id int) (*domain.RouteItem, error)
	Save(ctx context.Context, route *domain.RouteItemRequestData) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, route *domain.RouteItemRequestData) error
}
