package storage

import (
	"context"
	"hash"

	"github.com/sheeiavellie/avito040424/data"
)

type Storage interface {
	GetBanners(
		ctx context.Context,
		filter *data.AdminBannerFilter,
	) ([]data.Banner, error)

	CreateBanner(
		ctx context.Context,
		banner *data.Banner,
	) error

	UpdateBanner(
		ctx context.Context,
		bannerID int,
	) error

	DeleteBanner(
		ctx context.Context,
		bannerID int,
	) error
}

type CacheStorage interface {
	GetBanner(
		ctx context.Context,
		bannerKey hash.Hash32,
	) (*data.Banner, error)

	SetBanner(
		ctx context.Context,
		bannerKey hash.Hash32,
		banner *data.Banner,
	) error
}
