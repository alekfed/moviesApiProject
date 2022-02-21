package main

import (
	_ "github.com/lib/pq"
	"log"
	"moviesApiProject/pkg/controllers"
	"moviesApiProject/pkg/sqldb"
	"net/http"
)

func main() {
	sqldb.ConnectDB()

	http.HandleFunc(controllers.ActorsBaseURL, controllers.ActorsAll)
	http.HandleFunc(controllers.ActorsBaseURL+"/", controllers.ActorsById)

	http.HandleFunc(controllers.BaseFilmsURL, controllers.FilmsAll)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
