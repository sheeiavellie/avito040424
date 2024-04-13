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

func HandleGetUserBanner(
	ctx context.Context,
	bannerRepo repository.BannerRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Printf("an error occur at HandleGetUserBanner: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

		var bannerRequest data.UserBannerRequest
		if err := schema.NewDecoder().Decode(&bannerRequest, r.Form); err != nil {
			log.Printf("an error occur at HandleGetUserBanner: %s", err)
			util.SerHTTPErrorBadRequest(w)
			return
		}

		banner, err := bannerRepo.GetBannerContent(
			ctx,
			bannerRequest.FeatureID,
			bannerRequest.TagIDs,
			bannerRequest.UseLastRevision,
		)
		if err != nil {
			log.Printf("an error occur at HandleGetUserBanner: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

		err = util.WriteJSON(w, http.StatusOK, banner)
		if err != nil {
			log.Printf("an error occur at HandleGetUserBanner: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}
	}
}
