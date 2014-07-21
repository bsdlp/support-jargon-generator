package main

import (
	"log"
	"time"

	"github.com/mxk/go-sqlite/sqlite3"
)

type jargon struct {
	ID        int
	name      string
	quote     string
	source    string
	timestamp time.Time
}

var jargonCollection []jargon

func main() {
	db, err := sqlite3.Open("sqlite3", "./jargon.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := `
	CREATE TABLE jargon (id integer NOT NULL PRIMARY KEY,
						 name TEXT,
						 quote TEXT,
						 source TEXT,
						 timestamp INTEGER)
	`
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
}
