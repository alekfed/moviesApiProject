package controllers

import (
	"moviesApiProject/pkg/models"
	"net/http"
)

func Actors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		models.GetActors(w, r)

	case http.MethodPost:
		models.CreateActor(w, r)

	default:
		w.Header().Set("Allow", "GET, POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
