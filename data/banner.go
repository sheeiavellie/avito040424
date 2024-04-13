package data

import (
	"time"

	_ "github.com/gorilla/schema"
)

type UserBannerRequest struct {
	FeatureID       int   `schema:"feature_id,required"`
	TagIDs          []int `schema:"tag_id,required"`
	UseLastRevision bool  `schema:"use_last_revision,default:false"`
}

type BannerFilterRequest struct {
	FeatureIDs []int `schema:"feature_id,default:|"`
	TagIDs     []int `schema:"tag_id,default:|"`
	Limit      int   `schema:"limit,default:10"`
	Offset     int   `schema:"offset,default:0"`
}

type BannerFilter struct {
	FeatureIDs []int
	TagIDs     []int
	Limit      int
	Offset     int
}

type Banner struct {
	ID        int           `json:"id"`
	FeatureID int           `json:"feature_id"`
	TagIDs    []string      `json:"tag_ids"`
	Content   BannerContent `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	IsActive  bool          `json:"is_active"`
}

type BannerContent struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}
