package sqldb

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("postgres", "user=alekfed database=alekfed sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
