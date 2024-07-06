package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"vortex/internal/config"
	"vortex/internal/logger"
	"vortex/internal/server"

	"golang.org/x/sync/errgroup"
)

// ✓
// ============part1===========
// chi router
// zaplog
// viper
// .env
// config.yaml
// ============part2===========
// graceful shutdown
// ручки без логики
//

func main() {
	// Startup, may panic, which is ok
	// Service shouldnt start if config cant be read or logger init fails
	config.InitConfig()
	logger.InitGlobalLogger()
	logger.Info("Service started")

	// Create context for shutdown via idle timeout
	ctx, cancel := context.WithTimeout(context.Background(), config.C.HTTPServer.IdleTimeout)
	if config.C.HTTPServer.IdleTimeout == 0 {
		ctx, cancel = context.WithCancel(context.Background())
	}

	// Minimal Graceful Shutdown
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		logger.Info("Graceful shutdown initiated")
		cancel()
	}()

	// State of this specific server
	// Should contant context
	// May contain other fields

	serverState := server.NewServerState(ctx)

	// Start Server in goroutine inside errgroup to be able to return an error
	ewg := &errgroup.Group{}
	ewg.Go(serverState.StartServer)

	// Wait for server to finish, return non-zero code if error occured
	if err := ewg.Wait(); err != nil {
		logger.Info(err)
		os.Exit(1)
	}
}
