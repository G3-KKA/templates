package handlers

import (
	"net/http"
	"vortex/internal/logger"
)

// Optional fields fo '/else' handler
type Else struct {
}

// Constructor for '/else' state
func NewElseState() *Else {
	return &Else{}
}

// Handler for '/else'
func (state *Else) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("/else got something")
}
