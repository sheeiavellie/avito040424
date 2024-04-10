package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sheeiavellie/avito040424/util"
)

func ValidateBannerQueryParameters(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tagIDs := r.URL.Query()["tag_id"]
		featureID := r.URL.Query().Get("feature_id")

		if featureID == "" || len(tagIDs) == 0 {
			err := fmt.Errorf("parameters weren't provided")
			log.Printf("an error occur at ValidateBannerQueryParameters: %s", err)
			util.SerHTTPErrorBadRequest(w)
			return
		}

		next(w, r)
	}
}
