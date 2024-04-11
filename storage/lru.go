package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/sheeiavellie/avito040424/data"
)

const (
	lruTTL = 5 * time.Minute
)

type LRUCacheStorage struct {
	cache *expirable.LRU[string, data.Banner]
}

func NewLRUCacheStorage(size int) *LRUCacheStorage {
	cache := expirable.NewLRU[string, data.Banner](size, nil, lruTTL)
	return &LRUCacheStorage{
		cache: cache,
	}
}

func (ls *LRUCacheStorage) GetBanner(
	ctx context.Context,
	bannerKey string,
) (*data.Banner, error) {
	banner, ok := ls.cache.Get(bannerKey)
	if !ok {
		err := fmt.Errorf("value is not cached")
		return nil, err
	}

	return &banner, nil
}

func (ls *LRUCacheStorage) SetBanner(
	ctx context.Context,
	bannerKey string,
	banner *data.Banner,
) error {
	ls.cache.Add(bannerKey, *banner)
	return nil
}
