package rest

import (
	"net/http"
)

type Service struct {
	Mux http.Handler
}

// Should be created in main()
// service, err := rest.NewService(db)
//
//	errgroup.Go(func() error {
//			service.something()
//	})
func NewService(config struct{}) (*Service, error) {
	return &Service{}, nil
}

//=================================================
