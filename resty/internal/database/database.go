package database

import (
	"context"
	"yet-again-templates/resty/internal/config"

	"github.com/jmoiron/sqlx"
)

func NewDatabase(ctx context.Context, config config.Config) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", "config.DSN")
}
