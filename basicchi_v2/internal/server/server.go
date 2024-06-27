package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"vortex/internal/config"
	"vortex/internal/logger"
	"vortex/internal/server/handlers"
	"vortex/internal/server/middleware"
	"vortex/internal/server/servererror"

	"github.com/go-chi/chi/v5"
)

// Do not use without context
type ServerState struct {
	ctx context.Context
}

// Returns ServerState with filled context field
func NewServerState(ctx context.Context) *ServerState {
	return &ServerState{
		ctx: ctx,
	}
}

// May (?Should) be started as goroutine
// Can be freely started in errgroup.Run
// Starter holds context
func (starter *ServerState) StartServer() error {
	addr := config.C.HTTPServer.Address

	// Check address
	if err := checkAddress(addr); err != nil {
		return err
	}

	// Create Router
	chiRouter := chi.NewRouter()

	// Route states
	rootState := handlers.NewRootState()
	elseState := handlers.NewElseState()
	generalState := handlers.NewGeneralState()

	// Routes
	chiRouter.Get("/", rootState.ServeHTTP)
	chiRouter.Get("/else", elseState.ServeHTTP)
	chiRouter.Get("/general", generalState.ServeHTTP)

	// Middleware
	logMdw := middleware.NewLoggingMdwState()
	mdw := logMdw.LoggingMiddleware(chiRouter)
	panicMdw := middleware.NewPanicMdwState()
	mdw = panicMdw.PanicMiddleware(mdw)

	// Server Options
	server := http.Server{
		Addr:              addr,
		Handler:           mdw,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		//BaseContext:       func(l net.Listener) context.Context { return starter.ctx },
	}

	// Start Server
	logger.Info("Start Serving at:", addr)
	go handleShutdown(starter.ctx, &server)
	server.ListenAndServe()

	logger.Info("Done serving, no errors occured")
	return nil

}

// Utilitary functions
func handleShutdown(ctx context.Context, server *http.Server) {
	<-ctx.Done()
	logger.Info("Shutting Down")
	server.Shutdown(ctx)
	logger.Info("Shutdown by context.Done call")
}
func checkAddress(addr string) error {
	if addr == "" {
		return servererror.New(fmt.Errorf("empty address"), ".StartServer().checkAddress()")
	}
	return nil

}
