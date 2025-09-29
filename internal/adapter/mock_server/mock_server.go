package mockServer

import (
	"fmt"
	"time"

	"github.com/aquaswim/moyoki/internal/config"
	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/aquaswim/moyoki/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type mockServer struct {
	cfg *config.AppConfig
	app *fiber.App
}

func (m mockServer) Start() error {
	log.Infof("Mock Server started on port %s", ":3001")
	return m.app.Listen(m.cfg.MockListenAddr)
}

func (m mockServer) Stop() error {
	return m.app.ShutdownWithTimeout(5 * time.Second)
}

func New(
	cfg *config.AppConfig,
	mockService port.RouteMockService,
) port.Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           10 * time.Second,
		WriteTimeout:          10 * time.Second,
		IdleTimeout:           120 * time.Second,
		ReadBufferSize:        8192,
		WriteBufferSize:       8192,
		ReduceMemoryUsage:     true,
		EnablePrintRoutes:     false,
	})
	app.Use(fiberRecover.New())
	app.Use(requestid.New())

	app.Use(func(c *fiber.Ctx) error {
		req := &domain.MockRequest{
			Method: c.Method(),
			Path:   c.Path(),
		}

		routeItem, err := mockService.Resolve(c.Context(), req)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("failed to resolve route: %s", err))
		}
		for _, header := range routeItem.ResponseHeaders {
			c.Set(header.Key, header.Value)
		}
		return c.Status(routeItem.ResponseCode).SendString(routeItem.ResponseBody)
	})

	return &mockServer{
		app: app,
		cfg: cfg,
	}
}
