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
}
