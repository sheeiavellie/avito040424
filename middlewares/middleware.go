package middlewares

import "net/http"

func MiddlewareTODO(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
