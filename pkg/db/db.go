package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

//TODO InitConnection

func InitConnection() {

}
