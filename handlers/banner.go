package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sheeiavellie/avito040424/data"
	"github.com/sheeiavellie/avito040424/repository"
	"github.com/sheeiavellie/avito040424/util"
)

func HandleGetBanners(
	ctx context.Context,
	bannerRepo repository.BannerRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

		var bannerRequest data.BannerFilterRequest
		if err := schema.NewDecoder().Decode(&bannerRequest, r.Form); err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SerHTTPErrorBadRequest(w)
			return
		}

		filter := &data.BannerFilter{
			FeatureIDs: bannerRequest.FeatureIDs,
			TagIDs:     bannerRequest.TagIDs,
			Limit:      bannerRequest.Limit,
			Offset:     bannerRequest.Offset,
		}
		banners, err := bannerRepo.GetBanners(ctx, filter)
		if err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

		err = util.WriteJSON(w, http.StatusOK, banners)
		if err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}
	}
}

func HandlePostBanner(
	ctx context.Context,
	bannerRepo repository.BannerRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
