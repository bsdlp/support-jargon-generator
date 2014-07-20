package main

import (
	"database/sql"
	"log"
	"time"
)

type jargon struct {
	ID     int
	name   string
	quote  string
	source string
	date   time.Time
}

var jargonCollection []jargon

func main() {
	db, err := sql.Open("sqlite3", "./jargon.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := `
	CREATE TABLE jargon (id integer NOT NULL PRIMARY KEY,
						 name VARCHAR(11),
						 quote TEXT,
						 source VARCHAR(11),
						 date TIMESTAMP)
	`
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
}
