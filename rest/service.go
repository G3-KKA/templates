package rest

import (
	"net/http"
)

type Service struct {
	// should be embeeded in Mux
	//HH  *handlers.HandlerHolder
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
