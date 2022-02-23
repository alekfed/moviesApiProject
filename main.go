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

	http.HandleFunc("/actors", controllers.Actors)
	http.HandleFunc("/films", controllers.Films)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
