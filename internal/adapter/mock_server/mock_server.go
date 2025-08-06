package mockServer

import (
	"fmt"
	"github.com/aquaswim/moyoki/internal/config"
	"github.com/aquaswim/moyoki/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
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
		str := fmt.Sprintf("Will handle this request info:\n%s", c.Request().String())
		// todo: handle mocking logic here
		return c.SendString(str)
	})

	return &mockServer{
		app: app,
		cfg: cfg,
	}
}
