package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/shopping-list/api/app/db"
	"github.com/skeletonkey/shopping-list/api/app/eventing"
	"github.com/skeletonkey/shopping-list/api/app/server"
)

const (
	shutdownDelay = 5 * time.Second
	version	   = "0.0.1"
)

func main() {
	wg := new(sync.WaitGroup)

	log := logger.Get()
	log.Info().Str("Version", version).Msg("Service starting up")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	context.AfterFunc(ctx, func() {
		log.Info().Msg("Shutdown signal received")
		time.Sleep(shutdownDelay)
		log.Info().Dur("shutdown delay", shutdownDelay).Msg("Forcing shutdown after delay")
		os.Exit(1)
	})

	startService(ctx, wg)

	wg.Wait()
	log.Info().Msg("Service shutting down")
}

func startService(ctx context.Context, wg *sync.WaitGroup) {
	log := logger.Get()
	log.Info().Msg("initializing dependencies")

	if exists, err := eventing.TouchFileExists(); exists || err != nil {
		if err != nil {
			log.Panic().Err(err).Msg("attempting to see if eventing touch file exists")
		}
		log.Panic().Msg("eventing touch file exists")
	}

	if err := db.New(ctx, wg); err != nil {
		log.Panic().Err(err).Msg("failed to initialize database")
	}

	server.New(ctx, shutdownDelay, wg)
}
