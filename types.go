package main

import (
	"time"
)

type Action struct {
	ID          int
	ActionName  string
	Occurrences []Occurrence
}

type Occurrence struct {
	ID       int
	ActionID int
	Time     time.Time
}
