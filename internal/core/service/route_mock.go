package service

import (
	"context"
	"errors"
	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/aquaswim/moyoki/internal/core/port"
	"github.com/gorilla/reverse"
	"strings"
)

type routeMock struct {
	routeRepo port.RouteRepository
}

func NewRouteMockService(
	routeRepo port.RouteRepository,
) port.RouteMockService {
	return &routeMock{
		routeRepo: routeRepo,
	}
}

func (r routeMock) Resolve(ctx context.Context, req *domain.MockRequest) (*domain.RouteItem, error) {
	// note: find can be filtered by method
	routes, err := r.routeRepo.Find(ctx)
	if err != nil {
		return nil, err
	}
	for _, route := range routes {
		if !strings.EqualFold(route.Method, req.Method) {
			continue
		}

		path, err := reverse.NewGorillaPath(route.Path, false)
		if err != nil {
			return nil, err
		}
		if path.MatchString(req.Path) {
			return &route, nil
		}
	}
	return nil, errors.New("route not found")
}
