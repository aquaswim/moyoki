package main

import (
	"github.com/aquaswim/moyoki/internal"
	"github.com/aquaswim/moyoki/internal/core/port"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := internal.NewContainer()

	panelServer := internal.ResolveNamed[port.Server](c, "panel")
	mockServer := internal.ResolveNamed[port.Server](c, "mock")

	go func() {
		err := panelServer.Start()
		if err != nil {
			log.Errorf("server start error: %s", err)
		}
	}()
	go func() {
		err := mockServer.Start()
		if err != nil {
			log.Errorf("server start error: %s", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-stopChan // This blocks the main thread until an interrupt is received
	log.Info("Gracefully shutting down...")
	_ = panelServer.Stop()
	_ = mockServer.Stop()

	log.Info("Running cleanup tasks...")

	// Your cleanup tasks go here
	err := internal.ResolveNamed[port.Closer](c, "db")()
	if err != nil {
		log.Errorf("db close error: %s", err)
	}
	// redisConn.Close()
	log.Info("Fiber was successful shutdown.")
}
