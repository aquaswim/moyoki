package httpServer

import (
	"github.com/aquaswim/ndobol/internal/core/port"
	"github.com/aquaswim/ndobol/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
)

type server struct {
	app *fiber.App
}

func (s server) Start() error {
	log.Infof("Server started on port %s", ":3000")
	return s.app.Listen(":3000")
}

func (s server) Stop() error {
	return s.app.ShutdownWithTimeout(5 * time.Second)
}

func NewServer() port.Server {
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
	app.Use(logger.New(
		logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		},
	))

	web.RegisterRoutes(app)

	return &server{
		app: app,
	}
}
