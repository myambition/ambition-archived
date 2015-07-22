package main

import (
	"time"
)

type User struct {
	ID       int
	UserName string
	Email    string
	Sets     []Set
}

type Set struct {
	ID      int
	SetName string
	Actions []Action
}

type Action struct {
	ID          int
	ActionName  string
	Occurrences []Occurrence
	Sets        []Set
}

type Occurrence struct {
	ID       int
	ActionID int
	Time     time.Time
}
