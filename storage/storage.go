package storage

import (
	"context"

	"github.com/sheeiavellie/avito040424/data"
)

type Storage interface {
	GetBannerContent(
		ctx context.Context,
		featureID int,
		tagIDs []int,
	) (*data.BannerContent, error)

	GetBanners(
		ctx context.Context,
		filter *data.BannerFilter,
	) ([]data.Banner, error)

	CreateBanner(
		ctx context.Context,
		banner *data.BannerContent,
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
	) (*data.BannerContent, bool)

	SetBanner(
		ctx context.Context,
		bannerKey string,
		banner *data.BannerContent,
	) bool
}
