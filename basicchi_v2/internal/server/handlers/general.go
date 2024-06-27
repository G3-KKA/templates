package handlers

import (
	"net/http"
	"vortex/internal/logger"
)

// Optional fields fo '/general' handler
type General struct {
}

// Constructor for '/general' state
func NewGeneralState() *General {
	return &General{}
}

// Handler for '/general'
func (state *General) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("/general got something")
}
