package repo

import (
	"context"
	"yet-again-templates/resty/internal/config"
	domain "yet-again-templates/resty/internal/domain/objects"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	UserById(ctx context.Context, id uint64) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
}
type ProductRepository interface {
	UpdateProduct(ctx context.Context, product *domain.Product) error
}

type repository struct {
	db *sqlx.DB
}

// CreateUser implements Repository.
func (r *repository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (username) VALUES ($1) RETURNING id", user.Name)
	return err
}

// UpdateProduct implements Repository.
func (r *repository) UpdateProduct(ctx context.Context, product *domain.Product) error {
	panic("unimplemented")
}

// UserById implements Repository.
func (r *repository) UserById(ctx context.Context, id uint64) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserRepository(ctx context.Context, config config.Config, db *sqlx.DB) (UserRepository, error) {
	// from config get DSN
	return &repository{db: db}, nil
}
func NewProductRepository(ctx context.Context, config config.Config, db *sqlx.DB) (ProductRepository, error) {
	// from config get DSN
	return &repository{db: db}, nil
}
