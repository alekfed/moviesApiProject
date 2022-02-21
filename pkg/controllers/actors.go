package controllers

import (
	"moviesApiProject/pkg/models"
	"net/http"
)

const ActorsBaseURL = "/actors"

func ActorsAll(w http.ResponseWriter, r *http.Request) {
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

func ActorsById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(ActorsBaseURL+"/"):]

	if id == "" {
		http.Redirect(w, r, ActorsBaseURL, http.StatusTemporaryRedirect)
		return
	}

	switch r.Method {
	case http.MethodGet:
		models.GetActorById(id, w)

	default:
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
