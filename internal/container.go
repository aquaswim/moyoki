package internal

import (
	"github.com/aquaswim/moyoki/internal/adapter/db"
	"github.com/aquaswim/moyoki/internal/adapter/db/repositories"
	mockServer "github.com/aquaswim/moyoki/internal/adapter/mock_server"
	panelServer "github.com/aquaswim/moyoki/internal/adapter/panel_server"
	"github.com/aquaswim/moyoki/internal/config"
	"github.com/aquaswim/moyoki/internal/core/service"
	"github.com/golobby/container/v3"
)

func NewContainer() container.Container {
	c := container.New()
	// initialize module here
	container.MustSingleton(c, config.MustLoad[config.AppConfig])
	container.MustSingleton(c, config.MustLoad[config.DBConfig])

	container.MustSingletonLazy(c, db.Connect)
	container.MustNamedSingletonLazy(c, "db", db.Closer)

	container.MustSingletonLazy(c, repositories.NewRouteRepository)

	container.MustSingletonLazy(c, service.NewRouteService)

	container.MustSingletonLazy(c, panelServer.NewPanelHandler)
	container.MustNamedSingletonLazy(c, "panel", panelServer.NewServer)
	container.MustNamedSingletonLazy(c, "mock", mockServer.New)
	return c
}

func ResolveNamed[T any](c container.Container, name string) T {
	var t T
	container.MustNamedResolve(c, &t, name)
	return t
}
