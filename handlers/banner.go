package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"slices"

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

		var bannerReq data.BannerFilterRequest
		if err := schema.NewDecoder().Decode(&bannerReq, r.Form); err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SerHTTPErrorBadRequest(w)
			return
		}

		filter := &data.BannerFilter{
			FeatureIDs: bannerReq.FeatureIDs,
			TagIDs:     bannerReq.TagIDs,
			Limit:      bannerReq.Limit,
			Offset:     bannerReq.Offset,
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
		bannerReq := data.BannerRequest{
			Content: data.BannerContent{
				Title: "default",
				Text:  "default",
				URL:   "default",
			},
			IsActive: false,
		}

		if err := util.ReadJSON(r, &bannerReq); err != nil {
			log.Printf("an error occur at HandlePostBanner: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

		if bannerReq.FeatureID == 0 || len(bannerReq.TagIDs) == 0 {
			err := fmt.Errorf("featureID or tagIDs weren't provided")
			log.Printf("an error occur at HandlePostBanner: %s", err)
			util.SerHTTPErrorBadRequest(w)
			return
		}

		bannerID, err := bannerRepo.CreateBanner(
			ctx,
			bannerReq.FeatureID,
			bannerReq.TagIDs,
			&bannerReq.Content,
			bannerReq.IsActive,
		)
		if err != nil {

		}

		err = util.WriteJSON(w, http.StatusOK, bannerID)
		if err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

	}
}
