package handlers

import (
	"net/http"
	"vortex/internal/logger"
)

// Optional fields fo '/' handler
type Root struct {
}

// Constructor for '/' state
func NewRootState() *Root {
	return &Root{}
}

// Handler for '/'
func (state *Root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("/ got something")
}
