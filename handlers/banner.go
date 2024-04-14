package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/schema"
	"github.com/sheeiavellie/avito040424/data"
	"github.com/sheeiavellie/avito040424/repository"
	"github.com/sheeiavellie/avito040424/storage"
	"github.com/sheeiavellie/avito040424/util"
)

func HandleGetBanners(
	ctx context.Context,
	bannerRepo repository.BannerRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SetHTTPErrorInternalServerError(w)
			return
		}

		var bannerReq data.BannerFilterRequest
		if err := schema.NewDecoder().Decode(&bannerReq, r.Form); err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SetHTTPErrorBadRequest(w)
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
			util.SetHTTPErrorInternalServerError(w)
			return
		}

		err = util.WriteJSON(w, http.StatusOK, banners)
		if err != nil {
			log.Printf("an error occur at HandleGetBanners: %s", err)
			util.SetHTTPErrorInternalServerError(w)
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
			util.SetHTTPErrorInternalServerError(w)
			return
		}

		if bannerReq.FeatureID == 0 || len(bannerReq.TagIDs) == 0 {
			err := fmt.Errorf("featureID or tagIDs weren't provided")
			log.Printf("an error occur at HandlePostBanner: %s", err)
			util.SetHTTPErrorBadRequest(w)
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
			log.Printf("an error occur at HandlePostBanner: %s", err)
			switch {
			case errors.Is(err, storage.ErrorFeatureOrTagDontExist):
				util.SetHTTPErrorBadRequest(w)
			case errors.Is(err, storage.ErrorBannerAlreadyExist):
				util.SetHTTPErrorConflict(w)
			default:
				util.SetHTTPErrorInternalServerError(w)
			}
			return
		}

		bannerRes := data.CreateBannerResponse{BannerID: bannerID}
		err = util.WriteJSON(w, http.StatusCreated, bannerRes)
		if err != nil {
			log.Printf("an error occur at HandlePostBanner: %s", err)
			util.SetHTTPErrorInternalServerError(w)
			return
		}
	}
}

func HandleDeleteBanner(
	ctx context.Context,
	bannerRepo repository.BannerRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bannerIDStr := r.PathValue("id")
		bannerID, err := strconv.Atoi(bannerIDStr)
		if err != nil {
			switch {
			case bannerID <= 0:
				err = fmt.Errorf("banner id can't be less than 1")
				log.Printf("an error occur at HandlePatchBanner: %s", err)
				util.SetHTTPErrorBadRequest(w)
			default:
				log.Printf("an error occur at HandlePatchBanner: %s", err)
				util.SetHTTPErrorInternalServerError(w)
			}
			return
		}

		if err := bannerRepo.DeleteBanner(ctx, bannerID); err != nil {
			log.Printf("an error occur at HandleDeleteBanner: %s", err)
			switch {
			case errors.Is(err, storage.ErrorBannerDontExist):
				util.SetHTTPErrorNotFound(w)
			default:
				util.SetHTTPErrorInternalServerError(w)
			}
			return
		}

		util.SetHTTPStatusNoContent(w)
	}
}

func HandlePatchBanner(
	ctx context.Context,
	bannerRepo repository.BannerRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bannerIDStr := r.PathValue("id")
		bannerID, err := strconv.Atoi(bannerIDStr)
		if err != nil {
			switch {
			case bannerID <= 0:
				err = fmt.Errorf("banner id can't be less than 1")
				log.Printf("an error occur at HandlePatchBanner: %s", err)
				util.SetHTTPErrorBadRequest(w)
			default:
				log.Printf("an error occur at HandlePatchBanner: %s", err)
				util.SetHTTPErrorInternalServerError(w)
			}
			return
		}

		bannerChan := make(chan *data.BannerPatchRequest)
		var errPatch error
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()

			err := bannerRepo.UpdateBanner(ctx, bannerID, bannerChan)
			if err != nil {
				close(bannerChan)
				errPatch = err
			}
		}()

		if existingBanner, ok := <-bannerChan; ok {
			if err := util.ReadJSON(r, existingBanner); err != nil {
				log.Printf("an error occur at HandlePatchBanner: %s", err)
				util.SetHTTPErrorInternalServerError(w)
				return
			}
			bannerChan <- existingBanner
		}

		wg.Wait()
		if errPatch != nil {
			log.Printf("an error occur at HandlePatchBanner: %s", errPatch)
			switch {
			case errors.Is(errPatch, storage.ErrorFeatureOrTagDontExist):
				util.SetHTTPErrorBadRequest(w)
			case errors.Is(errPatch, storage.ErrorBannerDontExist):
				util.SetHTTPErrorNotFound(w)
			case errors.Is(errPatch, storage.ErrorBannerAlreadyExist):
				util.SetHTTPErrorConflict(w)
			default:
				util.SetHTTPErrorInternalServerError(w)
			}
			return
		}
	}
}
