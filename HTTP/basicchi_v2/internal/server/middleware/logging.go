package middleware

import (
	"net/http"
	"vortex/internal/logger"
)

// Struct to hold state for LoggingMiddleware
type LoggingMdwState struct {
}

// Constructor for LoggingMdwState
func NewLoggingMdwState() *LoggingMdwState {
	return &LoggingMdwState{}

}

// Middleware to log every request that goes through mdw
func (state *LoggingMdwState) LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Got Request. Method: ", r.Method, " Path: ", r.URL.Path)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
