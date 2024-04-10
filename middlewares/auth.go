package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sheeiavellie/avito040424/api"
	"github.com/sheeiavellie/avito040424/data"
	"github.com/sheeiavellie/avito040424/util"
)

func AuthorizeToken(
	next http.HandlerFunc,
	requiredRole api.AccessRole,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("token")

		if tokenHeader == "" {
			err := fmt.Errorf("token wasn't provided")
			log.Printf("an error occur at ValidateToken: %s", err)
			util.SerHTTPErrorUnauthorized(w)
			return
		}

		tokenStr, err := base64.StdEncoding.DecodeString(tokenHeader)
		if err != nil {
			log.Printf("an error occur at ValidateToken: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

		var token data.Token
		err = json.Unmarshal(tokenStr, &token)
		if err != nil {
			log.Printf("an error occur at ValidateToken: %s", err)
			util.SerHTTPErrorInternalServerError(w)
			return
		}

		if !hasRoleOrAdmin(token.Role, requiredRole.GetName()) {
			err := fmt.Errorf("wrong role, access forbidden")
			log.Printf("an error occur at ValidateToken: %s", err)
			util.SerHTTPErrorForbidden(w)
			return
		}

		next(w, r)
	}
}

func hasRoleOrAdmin(tokenRole, targetRole string) bool {
	switch {
	case strings.EqualFold(tokenRole, targetRole):
		return true
	case strings.EqualFold(tokenRole, api.AdminRole.GetName()):
		return true
	}
	return false
}
