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
	) (*data.BannerContent, error)
}
