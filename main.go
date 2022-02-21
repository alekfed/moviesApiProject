package main

import (
	_ "github.com/lib/pq"
	"log"
	"moviesApiProject/pkg/controllers"
	"moviesApiProject/pkg/sqldb"
	"net/http"
	"os"
)

func main() {
	sqldb.ConnectDB(os.Getenv("DB_HOST"))

	http.HandleFunc(controllers.ActorsBaseURL, controllers.ActorsAll)
	http.HandleFunc(controllers.ActorsBaseURL+"/", controllers.ActorsById)

	http.HandleFunc(controllers.BaseFilmsURL, controllers.FilmsAll)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
