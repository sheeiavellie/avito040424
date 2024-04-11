package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sheeiavellie/avito040424/data"
	"github.com/sheeiavellie/avito040424/storage"
	"github.com/sheeiavellie/avito040424/util"
)

func HandleGetUserBanner(
	storage storage.Storage,
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
	}
}
