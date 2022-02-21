package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"moviesApiProject/pkg/sqldb"
	"net/http"
	"strconv"
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
		return nil, errors.New("First name is required")
	}

	if input.LastName == nil {
		return nil, errors.New("Last name is required")
	}

	return &input, nil
}

func NewCommonQueryParamsFromRequest(r *http.Request) (*CommonQueryParams, error) {
	actorsQP := CommonQueryParams{
		Limit:  LimitDefault,
		Offset: OffsetDefault,
	}

	q := r.URL.Query()

	if q.Has("limit") {
		err := actorsQP.ValidateLimit(q)
		if err != nil {
			return nil, err
		}
	}

	if q.Has("offset") {
		err := actorsQP.ValidateOffset(q)
		if err != nil {
			return nil, err
		}
	}

	return &actorsQP, nil
}

func GetActorById(id string, w http.ResponseWriter) {
	err := ValidateActorId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id = $1`

	row := sqldb.DB.QueryRow(query, id)

	var actor Actor

	err = row.Scan(&actor.ActorId, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		log.Println(err)
	}
}

func ValidateActorId(id string) error {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if parsedId < 1 {
		return errors.New("id must be greater than 0")
	}

	return nil
}

func GetActors(w http.ResponseWriter, r *http.Request) {
	q, err := NewCommonQueryParamsFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	actors := make([]*Actor, 0)

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
