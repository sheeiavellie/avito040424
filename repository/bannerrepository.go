package repository

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/sheeiavellie/avito040424/data"
	"github.com/sheeiavellie/avito040424/storage"
)

type bannerRepository struct {
	storage storage.Storage
	cache   storage.CacheStorage
}

func NewBannerRepository(
	storage storage.Storage,
	cache storage.CacheStorage,
) *bannerRepository {
	return &bannerRepository{
		storage: storage,
		cache:   cache,
	}
}

// TODO: Think more about do we really need to create key every time
func (br *bannerRepository) GetBanner(
	ctx context.Context,
	bannerRequest *data.UserBannerRequest,
) (*data.Banner, error) {

	var bannerKeyBuffer strings.Builder

	bannerKeyBuffer.WriteString(strconv.Itoa(bannerRequest.FeatureID))
	slices.Sort(bannerRequest.TagIDs)

	for _, e := range bannerRequest.TagIDs {
		bannerKeyBuffer.WriteString(strconv.Itoa(e))
	}
	bannerKey := bannerKeyBuffer.String()

	if !bannerRequest.UseLastRevision {
		if banner, ok := br.cache.GetBanner(ctx, bannerKey); ok {
			return banner, nil
		}
	}

	banner, err := br.storage.GetBanner(ctx, &data.UserBannerFilter{
		FeatureID: bannerRequest.FeatureID,
		TagIDs:    bannerRequest.TagIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("can't get banner from storage: %w", err)
	}
	_ = br.cache.SetBanner(ctx, bannerKey, banner)

	return banner, nil
}
