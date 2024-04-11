package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sheeiavellie/avito040424/data"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (ps *PostgresStorage) Close() error {
	err := ps.db.Close()
	return fmt.Errorf("error closing db: %w", err)
}

// # Performs transaction
func (ps *PostgresStorage) execTx(
	ctx context.Context,
	opts *sql.TxOptions,
	f func(tx *sql.Tx) error,
) error {
	tx, err := ps.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("error executing tx: %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	execErr := f(tx)
	if execErr != nil {
		txErr := tx.Rollback()
		return fmt.Errorf("error executing tx: %w, %w", execErr, txErr)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error executing tx: %w", err)
	}

	return nil
}

func (ps *PostgresStorage) GetBanners(
	ctx context.Context,
	filter *data.AdminBannerFilter,
) ([]data.Banner, error) {
	return nil, nil
}

func (ps *PostgresStorage) GetBanner(
	ctx context.Context,
	filter *data.UserBannerFilter,
) (*data.Banner, error) {
	query := `
    SELECT title, text, url FROM banners WHERE feature_id = $1 AND
    (tag_ids <@ $2 and tag_ids @> $2);`

	var banner data.Banner
	err := ps.db.QueryRowContext(
		ctx,
		query,
		filter.FeatureID,
		filter.TagIDs,
	).Scan(&banner.Title, &banner.Text, &banner.URL)
	if err != nil {
		return nil, err
	}

	return &banner, nil
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
