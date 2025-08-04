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

	server := internal.Resolve[port.Server](c)

	go func() {
		err := server.Start()
		if err != nil {
			log.Errorf("server start error: %s", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-stopChan // This blocks the main thread until an interrupt is received
	log.Info("Gracefully shutting down...")
	_ = server.Stop()

	log.Info("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	log.Info("Fiber was successful shutdown.")
}
