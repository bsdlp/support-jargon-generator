package main

import (
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
}
