package middleware

import "net/http"

// __FILLME_HIGH_REGISTER__
// __DESCRIPTION__

// Struct to hold state for __FILLME_HIGH_REGISTER__Middleware
type __FILLME_HIGH_REGISTER__MdwState struct {
}

// Constructor for __FILLME_HIGH_REGISTER__MdwState
func New__FILLME_HIGH_REGISTER__MdwState() *__FILLME_HIGH_REGISTER__MdwState {
	return &__FILLME_HIGH_REGISTER__MdwState{}

}

// Middleware to __DESCRIPTION__
func (state *__FILLME_HIGH_REGISTER__MdwState) __FILLME_HIGH_REGISTER__Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
