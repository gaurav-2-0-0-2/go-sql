package db

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Connectdb(path string) (*sql.DB, error) {
	// Connecting to sqlite db
	db, dberr := sql.Open("sqlite3", path)
	if dberr != nil {
		log.Fatal(dberr)
	}

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to db successfully\n")
	return db, nil
}