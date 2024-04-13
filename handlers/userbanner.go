package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sheeiavellie/avito040424/data"
	"github.com/sheeiavellie/avito040424/repository"
	"github.com/sheeiavellie/avito040424/storage"
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

		var bannerReq data.UserBannerRequest
		if err := schema.NewDecoder().Decode(&bannerReq, r.Form); err != nil {
			log.Printf("an error occur at HandleGetUserBanner: %s", err)
			util.SerHTTPErrorBadRequest(w)
			return
		}

		banner, err := bannerRepo.GetBannerContent(
			ctx,
			bannerReq.FeatureID,
			bannerReq.TagIDs,
			bannerReq.UseLastRevision,
		)
		if err != nil {
			log.Printf("an error occur at HandleGetUserBanner: %s", err)
			switch {
			case errors.Is(err, storage.ErrorBannerIsNotActive):
				util.SerHTTPErrorConflict(w)
			default:
				util.SerHTTPErrorInternalServerError(w)
			}
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
