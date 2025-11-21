package port

import (
	"context"

	"github.com/aquaswim/moyoki/internal/core/domain"
)

type RouteRepository interface {
	CrudRepository[domain.RouteItem]
	Count(ctx context.Context) (int64, error)
}

type RouteService interface {
	Find(ctx context.Context) ([]domain.RouteItem, error)
	FindByID(ctx context.Context, id int) (*domain.RouteItem, error)
	Save(ctx context.Context, route *domain.RouteItemRequestData) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, route *domain.RouteItemRequestData) error
	Count(ctx context.Context) (int64, error)
}

type RouteMockService interface {
	Resolve(ctx context.Context, req *domain.MockRequest) (*domain.RouteItem, error)
}
