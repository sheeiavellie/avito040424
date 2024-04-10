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

func (ps *PostgresStorage) execTx(
	ctx context.Context,
	opts *sql.TxOptions,
	f func(tx *sql.Tx) (*sql.Rows, error),
) (*sql.Rows, error) {
	tx, err := ps.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error executing tx: %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	res, execErr := f(tx)
	if execErr != nil {
		txErr := tx.Rollback()
		return nil, fmt.Errorf("error executing tx: %w, %w", execErr, txErr)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error executing tx: %w", err)
	}

	return res, nil
}

func (ps *PostgresStorage) GetBanners(
	ctx context.Context,
	filter *data.AdminBannerFilter,
) ([]data.Banner, error) {
	query := `SELECT * FROM banners;`

	res, err := ps.execTx(ctx, nil, func(tx *sql.Tx) (*sql.Rows, error) {
		return tx.QueryContext(ctx, query)
	})
	if err != nil {
		return nil, fmt.Errorf("error at GetBanners: %w", err)
	}
	defer res.Close()

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
