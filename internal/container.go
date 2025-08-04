package internal

import (
	httpServer "github.com/aquaswim/ndobol/internal/adapter/http_server"
	"github.com/aquaswim/ndobol/internal/core/port"
	"github.com/golobby/container/v3"
)

func NewContainer() container.Container {
	c := container.New()
	// initialize module here

	container.MustSingletonLazy(c, func() port.Server {
		return httpServer.NewServer()
	})
	return c
}

func Resolve[T any](c container.Container) T {
	var t T
	container.MustResolve(c, &t)
	return t
}
