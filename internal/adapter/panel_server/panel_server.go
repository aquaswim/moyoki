package httpServer

import (
	"time"

	"github.com/aquaswim/moyoki/internal/config"
	"github.com/aquaswim/moyoki/internal/core/port"
	"github.com/aquaswim/moyoki/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type server struct {
	cfg *config.AppConfig
	app *fiber.App
}

func (s server) Start() error {
	log.Infof("Panel Server started on port %s", ":3000")
	return s.app.Listen(s.cfg.PanelListenAddr)
}

func (s server) Stop() error {
	return s.app.ShutdownWithTimeout(5 * time.Second)
}

func NewServer(
	cfg *config.AppConfig,
	panelHandler *PanelHandler,
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
	app.Use(logger.New(
		logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		},
	))

	// register all panel rout
	panelHandler.RegisterHandler(app)

	web.RegisterRoutes(app)

	return &server{
		app: app,
		cfg: cfg,
	}
}
