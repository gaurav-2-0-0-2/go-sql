package db

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connectdb(path string) error {
	var dberr error
	DB, dberr  = sql.Open("sqlite3", path)
	if dberr != nil {
		log.Fatal(dberr)
	}

	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to db successfully\n")
	return nil
}
