package util

import (
	"encoding/json"
	"net/http"
)

func SetHTTPStatusNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func SetHTTPErrorBadRequest(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusBadRequest),
		http.StatusBadRequest,
	)
}

func SetHTTPErrorUnauthorized(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusUnauthorized),
		http.StatusUnauthorized,
	)
}

func SetHTTPErrorForbidden(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusForbidden),
		http.StatusForbidden,
	)
}

func SetHTTPErrorNotFound(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusNotFound),
		http.StatusNotFound,
	)
}

func SetHTTPErrorConflict(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusConflict),
		http.StatusConflict,
	)
}

func SetHTTPErrorInternalServerError(w http.ResponseWriter) {
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
