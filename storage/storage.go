package storage

import (
	"context"
	"hash"

	"github.com/sheeiavellie/avito040424/data"
)

type Storage interface {
}

type CacheStorage interface {
	GetBanner(ctx context.Context, bannerKey hash.Hash32) (*data.Banner, error)
	SetBanner(ctx context.Context, banner *data.Banner) error
}
