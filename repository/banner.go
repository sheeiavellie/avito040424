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
// It mb good to start cooking key when request was parsed
// It can be done with gorutine and channel:
// add string channel
// as soon as request was parsed, start gorutine
// write result into channel and get it here

func (br *bannerRepository) GetBannerContent(
	ctx context.Context,
	featureID int,
	tagIDs []int,
	useLastRevision bool,
) (*data.BannerContent, error) {

	var bannerKeyBuffer strings.Builder

	bannerKeyBuffer.WriteString(strconv.Itoa(featureID))
	slices.Sort(tagIDs)

	for _, e := range tagIDs {
		bannerKeyBuffer.WriteString(strconv.Itoa(e))
	}
	bannerKey := bannerKeyBuffer.String()

	if !useLastRevision {
		if banner, ok := br.cache.GetBanner(ctx, bannerKey); ok {
			return banner, nil
		}
	}

	banner, err := br.storage.GetBannerContent(
		ctx,
		featureID,
		tagIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("can't get banner from storage: %w", err)
	}

	go func() {
		_ = br.cache.SetBanner(ctx, bannerKey, banner)
	}()

	return banner, nil
}

func (br *bannerRepository) GetBanners(
	ctx context.Context,
	filter *data.BannerFilter,
) ([]data.Banner, error) {
	banners, err := br.storage.GetBanners(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("can't get banners from storage: %w", err)
	}
	return banners, nil
}

func (br *bannerRepository) CreateBanner(
	ctx context.Context,
	featureID int,
	tagIDs []int,
	content *data.BannerContent,
	isActive bool,
) (int, error) {
	bannerID, err := br.storage.CreateBanner(
		ctx,
		featureID,
		tagIDs,
		content,
		isActive,
	)
	if err != nil {
		return 0, fmt.Errorf("can't create banner: %w", err)
	}

	return bannerID, nil
}

func (br *bannerRepository) UpdateBanner(
	ctx context.Context,
	bannerID int,
	featureID int,
	tagIDs []int,
	content *data.BannerContent,
	isActive bool,
) error {
	return nil
}

func (br *bannerRepository) DeleteBanner(
	ctx context.Context,
	bannerID int,
) error {
	err := br.storage.DeleteBanner(ctx, bannerID)
	if err != nil {
		return fmt.Errorf("can't delete banner: %w", err)
	}

	return nil
}
