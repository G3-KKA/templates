package middleware

import (
	"net/http"
	"vortex/internal/logger"
)

// Struct to hold state for PanicMiddleware
type PanicMdwState struct {
}

// Constructor for PanicMdwState
func NewPanicMdwState() *PanicMdwState {
	return &PanicMdwState{}

}

// Middleware to Handle panic via recover()
func (state *PanicMdwState) PanicMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Debug("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
