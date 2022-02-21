package sqldb

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func ConnectDB(host string) {
	dsn := fmt.Sprintf("user=alekfed database=alekfed sslmode=disable host=%s", host)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
