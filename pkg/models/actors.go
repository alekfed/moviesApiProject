package models

import (
	"encoding/json"
	"errors"
	"log"
	"moviesApiProject/pkg/sqldb"
	"net/http"
	"time"
)

type Actor struct {
	ActorId    int       `json:"actor_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	LastUpdate time.Time `json:"last_update"`
}

type CreateActorInput struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

func NewCreateActorInputFromRequest(r *http.Request) (*CreateActorInput, error) {
	var input CreateActorInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return nil, err
	}

	if input.FirstName == nil {
		return nil, errors.New(`"first_name" is required`)
	}
	if input.LastName == nil {
		return nil, errors.New(`"last_name" is required`)
	}

	return &input, nil
}

func GetActors(w http.ResponseWriter, r *http.Request) {
	q, err := NewCommonQueryParamsFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	actors := make([]*Actor, 0)

	if q.Id != nil {
		q.Limit = 1
		q.Offset = *q.Id - 1
	}

	query := `SELECT actor_id, first_name, last_name, last_update FROM actor ORDER BY actor_id LIMIT $1 OFFSET $2`

	rows, err := sqldb.DB.Query(query, q.Limit, q.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	for rows.Next() {
		var actor Actor

		err = rows.Scan(&actor.ActorId, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		actors = append(actors, &actor)
	}

	err = rows.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string][]*Actor{"actors": actors})
	if err != nil {
		log.Println(err)
	}
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	actor, err := NewCreateActorInputFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Actor added: %s %s", *actor.FirstName, *actor.LastName)

	w.WriteHeader(http.StatusCreated)
}
