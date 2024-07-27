package user

import (
	"context"
	"yet-again-templates/resty/internal/cache"
	"yet-again-templates/resty/internal/config"
	domain "yet-again-templates/resty/internal/domain/objects"
	"yet-again-templates/resty/internal/domain/repo"
)

type IuserService interface {
	UserById(ctx context.Context, id uint64) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
}
type /* other module */ userService struct {
	cache cache.Cacher
	repo  repo.UserRepository
}

// CreateUser implements IuserService.
func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	return s.repo.CreateUser(ctx, user)
}

// UserById implements IuserService.
func (s *userService) UserById(ctx context.Context, id uint64) (*domain.User, error) {
	/* authorize */
	/* try cache */
	return s.repo.UserById(ctx, id)
}

func NewUserService(ctx context.Context, config config.Config, repo repo.UserRepository, cache cache.Cacher) (IuserService, error) {
	return &userService{
		repo: repo,
	}, nil
}
