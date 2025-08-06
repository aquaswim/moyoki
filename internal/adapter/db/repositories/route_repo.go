package repositories

import (
	"context"
	"encoding/json"
	"github.com/aquaswim/moyoki/internal/adapter/db/repositories/model"
	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/aquaswim/moyoki/internal/core/port"
	"gorm.io/gorm"
	"time"
)

type routeRepository struct {
	routeDB gorm.Interface[model.Route]
}

func NewRouteRepository(
	db *gorm.DB,
) port.RouteRepository {
	return &routeRepository{
		routeDB: gorm.G[model.Route](db),
	}
}

func (r routeRepository) Find(ctx context.Context) ([]domain.RouteItem, error) {
	rows, err := r.routeDB.Find(ctx)
	if err != nil {
		return nil, err
	}

	routes := make([]domain.RouteItem, 0, len(rows))
	for _, row := range rows {
		routes = append(routes, *routeModelToEntity(&row))
	}
	return routes, nil
}

func (r routeRepository) FindByID(ctx context.Context, id int) (*domain.RouteItem, error) {
	row, err := r.routeDB.Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return routeModelToEntity(&row), nil
}

func (r routeRepository) Save(ctx context.Context, route *domain.RouteItem) error {
	d := routeToModel(route)

	return r.routeDB.Create(ctx, d)
}

func (r routeRepository) Delete(ctx context.Context, id int) error {
	_, err := r.routeDB.Where("id = ?", id).Delete(ctx)
	return err
}

func (r routeRepository) Update(ctx context.Context, route *domain.RouteItem) error {
	headers, _ := json.Marshal(route.ResponseHeader)

	_, err := r.routeDB.Where("id = ?").Updates(ctx, model.Route{
		Path:            route.Path,
		Description:     route.Description,
		ResponseCode:    route.ResponseCode,
		ResponseBody:    route.ResponseBody,
		ResponseHeaders: headers,
	})
	return err
}

func routeToModel(route *domain.RouteItem) *model.Route {
	rawResponseHeaders, _ := json.Marshal(route.ResponseHeader)

	return &model.Route{
		ID:              route.ID,
		Method:          route.Method,
		Path:            route.Path,
		Description:     route.Description,
		ResponseCode:    route.ResponseCode,
		ResponseBody:    route.ResponseBody,
		ResponseHeaders: rawResponseHeaders,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func routeModelToEntity(route *model.Route) *domain.RouteItem {
	var headers []domain.RouteResponseHeader
	_ = json.Unmarshal(route.ResponseHeaders, &headers)
	return &domain.RouteItem{
		ID:             route.ID,
		Method:         route.Method,
		Path:           route.Path,
		Description:    route.Description,
		ResponseCode:   route.ResponseCode,
		ResponseBody:   route.ResponseBody,
		ResponseHeader: headers,
	}
}
