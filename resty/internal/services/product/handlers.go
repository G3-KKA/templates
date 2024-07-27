package product

import (
	"context"
	"net/http"
	"yet-again-templates/resty/internal/config"
)

func NewProductHandler(ctx context.Context, config config.Config, service IproductService) (IproductHandler, error) {
	return &productHandler{
		service: service,
	}, nil
}

func (p *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {}

type IproductHandler interface {
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}
