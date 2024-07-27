package product

import (
	"context"
	"yet-again-templates/resty/internal/config"
	"yet-again-templates/resty/internal/domain/repo"
)

type /* other module */ productService struct {
	repo repo.ProductRepository
}

func NewProductService(ctx context.Context, config config.Config, repo repo.ProductRepository) (IproductService, error) {
	return &productService{}, nil

}

type IproductService interface {
}
type /* other module */ productHandler struct {
	service IproductService
}
