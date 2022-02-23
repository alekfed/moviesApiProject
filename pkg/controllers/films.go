package controllers

import (
	"moviesApiProject/pkg/models"
	"net/http"
)

func Films(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		models.GetFilms(w, r)

	default:
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
