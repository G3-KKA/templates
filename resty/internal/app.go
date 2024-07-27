package poc

import (
	"context"
	"net/http"
	"yet-again-templates/resty/internal/cache"
	"yet-again-templates/resty/internal/config"
	"yet-again-templates/resty/internal/database"
	"yet-again-templates/resty/internal/domain/repo"
	"yet-again-templates/resty/internal/services/product"
	"yet-again-templates/resty/internal/services/user"

	"github.com/go-chi/chi/v5"
)

type BusinesApp struct {
	PublicMux http.Handler
	// logger

}

func NewBusinesApp(ctx context.Context, config config.Config) (app *BusinesApp, err error) {
	db, _ := database.NewDatabase(ctx, config)

	urepo, _ := repo.NewUserRepository(ctx, config, db)
	cache := cache.NewCacher()
	// MAY CONSIST OF MULTIPLE SUB-SERVICES
	uservice, _ := user.NewUserService(ctx, config, urepo, cache)
	uhandlers, _ := user.NewUserHandler(ctx, config, uservice)

	chiMux := chi.NewMux()
	// .Mount() ??? should be even better
	chiMux.Route("/user/{uuid}", func(r chi.Router) {
		r.Get("/", uhandlers.GetUser)
		r.Post("/", uhandlers.CreateUser)
	})
	prepo, _ := repo.NewProductRepository(ctx, config, db)
	pService, _ := product.NewProductService(ctx, config, prepo)
	phandler, _ := product.NewProductHandler(ctx, config, pService)

	chiMux.Route("/product/{prodid}", func(r chi.Router) {
		r.Put("/", phandler.UpdateProduct)
	})

	return &BusinesApp{
			PublicMux: chiMux,
		},
		nil
}
