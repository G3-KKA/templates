package handlers

import (
	"net/http"
	"vortex/internal/logger"
)

//__FILLME_LOW_REGISTER__
//
//__FILLME_HIGH_REGISTER__

// Optional fields fo '/__FILLME_LOW_REGISTER__' handler
type __FILLME_HIGH_REGISTER__ struct {
}

// Constructor for '/__FILLME_LOW_REGISTER__' state
func New__FILLME_HIGH_REGISTER__State() *__FILLME_HIGH_REGISTER__ {
	return &__FILLME_HIGH_REGISTER__{}
}

// Handler for '/__FILLME_LOW_REGISTER__'
func (state *__FILLME_HIGH_REGISTER__) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("/__FILLME_LOW_REGISTER__ got something")
}
