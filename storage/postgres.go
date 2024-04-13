package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/sheeiavellie/avito040424/data"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(
	ctx context.Context,
	connStr string,
) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err = db.PingContext(ctxPing)
	if err != nil {
		return nil, fmt.Errorf("error pinging postgres: %w", err)
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (ps *PostgresStorage) Close() error {
	err := ps.db.Close()
	return fmt.Errorf("error closing db: %w", err)
}

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

func (ps *PostgresStorage) GetBannerContent(
	ctx context.Context,
	featureID int,
	tagIDs []int,
) (*data.BannerContent, error) {
	query := `
    SELECT title, text, url, is_active FROM banners WHERE feature_id = $1 AND
    (tag_ids <@ $2 and tag_ids @> $2);`

	var banner data.BannerContent
	var isActive bool
	err := ps.db.QueryRowContext(
		ctx,
		query,
		featureID,
		pq.Array(tagIDs),
	).Scan(&banner.Title, &banner.Text, &banner.URL, &isActive)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	if !isActive {
		return nil, ErrorBannerIsNotActive
	}

	return &banner, nil
}

func (ps *PostgresStorage) GetBanners(
	ctx context.Context,
	filter *data.BannerFilter,
) ([]data.Banner, error) {
	query := `
    SELECT * FROM banners 
    WHERE feature_id = ANY($1) AND $2 <@ tag_ids
    ORDER BY id
    LIMIT $3 OFFSET $4;`

	res, err := ps.db.QueryContext(
		ctx,
		query,
		pq.Array(filter.FeatureIDs),
		pq.Array(filter.TagIDs),
		filter.Limit,
		filter.Offset,
	)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer res.Close()

	banners := make([]data.Banner, 0, filter.Limit)

	for res.Next() {
		var banner data.Banner
		if err := res.Scan(
			&banner.ID,
			&banner.FeatureID,
			pq.Array(&banner.TagIDs),
			&banner.Content.Title,
			&banner.Content.Text,
			&banner.Content.URL,
			&banner.CreatedAt,
			&banner.UpdatedAt,
			&banner.IsActive,
		); err != nil {
			log.Printf("%s", fmt.Errorf("error scanning result: %w", err))
		}
		banners = append(banners, banner)
	}

	if rerr := res.Close(); rerr != nil {
		return nil, fmt.Errorf("error closing result: %w", err)
	}
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("error scanning result: %w", err)
	}

	return banners, nil
}

func (ps *PostgresStorage) CreateBanner(
	ctx context.Context,
	banner *data.BannerContent,
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
