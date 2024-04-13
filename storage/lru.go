package storage

import (
	"context"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/sheeiavellie/avito040424/data"
)

type LRUCacheStorage struct {
	cache *expirable.LRU[string, data.BannerContent]
}

func NewLRUCacheStorage(size int) *LRUCacheStorage {
	cache := expirable.NewLRU[string, data.BannerContent](size, nil, 0)
	return &LRUCacheStorage{
		cache: cache,
	}
}

func (ls *LRUCacheStorage) GetBanner(
	ctx context.Context,
	bannerKey string,
) (*data.BannerContent, bool) {
	banner, ok := ls.cache.Get(bannerKey)
	return &banner, ok
}

// Refactor method signature to know if value was added
// Someday later, don't really need that kind of thing now
func (ls *LRUCacheStorage) SetBanner(
	ctx context.Context,
	bannerKey string,
	banner *data.BannerContent,
) bool {
	return ls.cache.Add(bannerKey, *banner)
}
