package service

import (
	"context"
	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/aquaswim/moyoki/internal/core/port"
)

type routeService struct {
	repo port.RouteRepository
}

func NewRouteService(repo port.RouteRepository) port.RouteService {
	return &routeService{repo: repo}
}

func (r routeService) Find(ctx context.Context) ([]domain.RouteItem, error) {
	return r.repo.Find(ctx)
}

func (r routeService) FindByID(ctx context.Context, id int) (*domain.RouteItem, error) {
	return r.repo.FindByID(ctx, id)
}

func (r routeService) Save(ctx context.Context, route *domain.RouteItemRequestData) error {
	routeItem := &domain.RouteItem{
		Method:         route.Method,
		Path:           route.Path,
		Description:    route.Description,
		ResponseCode:   route.ResponseCode,
		ResponseHeader: route.ResponseHeaders,
		ResponseBody:   route.ResponseBody,
	}
	return r.repo.Save(ctx, routeItem)
}

func (r routeService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func (r routeService) Update(ctx context.Context, id int, route *domain.RouteItemRequestData) error {
	routeItem := &domain.RouteItem{
		ID:             id,
		Method:         route.Method,
		Path:           route.Path,
		Description:    route.Description,
		ResponseCode:   route.ResponseCode,
		ResponseHeader: route.ResponseHeaders,
		ResponseBody:   route.ResponseBody,
	}
	return r.repo.Update(ctx, routeItem)
}
