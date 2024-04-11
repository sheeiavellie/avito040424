package storage

import (
	"context"

	"github.com/sheeiavellie/avito040424/data"
)

type Storage interface {
	GetBanners(
		ctx context.Context,
		filter *data.AdminBannerFilter,
	) ([]data.Banner, error)

	GetBanner(
		ctx context.Context,
		filter *data.UserBannerFilter,
	) (*data.Banner, error)

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
		bannerKey string,
	) (*data.Banner, error)

	SetBanner(
		ctx context.Context,
		bannerKey string,
		banner *data.Banner,
	) error
}
