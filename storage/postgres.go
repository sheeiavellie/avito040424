package storage

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sheeiavellie/avito040424/data"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (ps *PostgresStorage) GetBanners(
	ctx context.Context,
	filter *data.AdminBannerFilter,
) ([]data.Banner, error) {
	return nil, nil
}

func (ps *PostgresStorage) CreateBanner(
	ctx context.Context,
	banner *data.Banner,
) error {
	return nil
}

func (ps *PostgresStorage) UpdateBanner(
	ctx context.Context,
	bannerID int,
) error {
	return nil
}

func (ps *PostgresStorage) DeleteBanner(
	ctx context.Context,
	bannerID int,
) error {
	return nil
}
