package util

import (
	"encoding/json"
	"net/http"
)

func SerHTTPErrorBadRequest(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusBadRequest),
		http.StatusBadRequest,
	)
}

func SerHTTPErrorUnauthorized(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusUnauthorized),
		http.StatusUnauthorized,
	)
}

func SerHTTPErrorForbidden(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusForbidden),
		http.StatusForbidden,
	)
}

func SerHTTPErrorNotFound(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusNotFound),
		http.StatusNotFound,
	)
}

func SerHTTPErrorInternalServerError(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

func ReadJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
