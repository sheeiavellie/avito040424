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
) (*data.Banner, bool) {
	banner, ok := ls.cache.Get(bannerKey)
	return &banner, ok
}

// TODO: Refactor method signature to know if value was added
func (ls *LRUCacheStorage) SetBanner(
	ctx context.Context,
	bannerKey string,
	banner *data.Banner,
) bool {
	return ls.cache.Add(bannerKey, *banner)
}
