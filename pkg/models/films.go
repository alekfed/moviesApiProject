package models

import (
	"encoding/json"
	"log"
	"moviesApiProject/pkg/sqldb"
	"net/http"
)

type Film struct {
	FilmId      int    `json:"film_id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	Rating      string `json:"rating"`
}

func GetFilms(w http.ResponseWriter, r *http.Request) {
	q, err := NewCommonQueryParamsFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	films := make([]*Film, 0)

	if q.Id != nil {
		q.Limit = 1
		q.Offset = *q.Id - 1
	}

	query := `SELECT film_id, title, release_year, rating FROM film ORDER BY film_id LIMIT $1 OFFSET $2`

	rows, err := sqldb.DB.Query(query, q.Limit, q.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	for rows.Next() {
		var film Film

		err = rows.Scan(&film.FilmId, &film.Title, &film.ReleaseYear, &film.Rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		films = append(films, &film)
	}

	err = rows.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string][]*Film{"films": films})
	if err != nil {
		log.Println(err)
	}
}
