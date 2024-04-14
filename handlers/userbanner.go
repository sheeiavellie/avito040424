package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sheeiavellie/avito040424/data"
	"github.com/sheeiavellie/avito040424/repository"
	"github.com/sheeiavellie/avito040424/storage"
	"github.com/sheeiavellie/avito040424/util"
)

func HandleGetUserBanner(
	ctx context.Context,
	timeout time.Duration,
	bannerRepo repository.BannerRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		done := make(chan struct{})

		go func() {
			defer func() {
				done <- struct{}{}
			}()
			if err := r.ParseForm(); err != nil {
				log.Printf("an error occur at HandleGetUserBanner: %s", err)
				util.SetHTTPErrorInternalServerError(w)
				return
			}

			var bannerReq data.UserBannerRequest
			if err := schema.NewDecoder().Decode(&bannerReq, r.Form); err != nil {
				log.Printf("an error occur at HandleGetUserBanner: %s", err)
				util.SetHTTPErrorBadRequest(w)
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
				case errors.Is(err, storage.ErrorBannerDontExist):
					util.SetHTTPErrorNotFound(w)
				case errors.Is(err, storage.ErrorBannerIsNotActive):
					util.SetHTTPErrorConflict(w)
				default:
					util.SetHTTPErrorInternalServerError(w)
				}
				return
			}

			err = util.WriteJSON(w, http.StatusOK, banner)
			if err != nil {
				log.Printf("an error occur at HandleGetUserBanner: %s", err)
				util.SetHTTPErrorInternalServerError(w)
				return
			}
		}()

		select {
		case <-ctx.Done():
			util.SetHTTPErrorRequestTimeout(w)
		case <-done:

		}
	}
}
