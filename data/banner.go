package data

import _ "github.com/gorilla/schema"

type UserBannerRequest struct {
	FeatureID       int   `schema:"feature_id,required"`
	TagIDs          []int `schema:"tag_id,required"`
	UseLastRevision bool  `schema:"use_last_revision,default:false"`
}

type AdminBannerRequest struct {
	FeatureID int   `schema:"feature_id,required"`
	TagIDs    []int `schema:"tag_id,required"`
	Limit     int   `schema:"limit,default:10"`
	Offset    int   `schema:"offset,default:10"`
}
