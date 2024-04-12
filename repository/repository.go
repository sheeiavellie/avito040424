package repository

import (
	"context"

	"github.com/sheeiavellie/avito040424/data"
)

type BannerRepository interface {
	GetBanner(
		ctx context.Context,
		bannerRequest *data.UserBannerRequest,
	) (*data.Banner, error)
}
