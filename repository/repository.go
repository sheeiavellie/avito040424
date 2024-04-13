package repository

import (
	"context"

	"github.com/sheeiavellie/avito040424/data"
)

type BannerRepository interface {
	GetBannerContent(
		ctx context.Context,
		featureID int,
		tagIDs []int,
		useLastRevision bool,
	) (*data.BannerContent, error)

	GetBanners(
		ctx context.Context,
		filter *data.BannerFilter,
	) ([]data.Banner, error)

	CreateBanner(
		ctx context.Context,
		featureID int,
		tagIDs []int,
		content *data.BannerContent,
		isActive bool,
	) (int, error)

	UpdateBanner(
		ctx context.Context,
		bannerID int,
		featureID int,
		tagIDs []int,
		content *data.BannerContent,
		isActive bool,
	) error

	DeleteBanner(
		ctx context.Context,
		bannerID int,
	) error
}
