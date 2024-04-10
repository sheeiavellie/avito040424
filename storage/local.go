package storage

import (
	"context"
	"fmt"
	"hash"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/sheeiavellie/avito040424/data"
)

const (
	lruTTL = 5 * time.Minute
)

type LocalLRUStorage struct {
	cache *expirable.LRU[hash.Hash32, data.Banner]
}

func NewLocalLRUStorage(size int) *expirable.LRU[hash.Hash32, data.Banner] {
	return expirable.NewLRU[hash.Hash32, data.Banner](size, nil, lruTTL)
}

func (ls *LocalLRUStorage) GetBanner(
	ctx context.Context,
	bannerKey hash.Hash32,
) (*data.Banner, error) {
	banner, ok := ls.cache.Get(bannerKey)
	if !ok {
		err := fmt.Errorf("value is not cached")
		return nil, err
	}

	return &banner, nil
}

func (ls *LocalLRUStorage) SetBanner(
	ctx context.Context,
	bannerKey hash.Hash32,
	banner *data.Banner,
) error {
	ls.cache.Add(bannerKey, *banner)
	return nil
}
