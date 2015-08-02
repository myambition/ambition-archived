package main

import (
	"time"
)

type User struct {
	Id       int
	UserName string
	Email    string
	Sets     []Set
}

type Set struct {
	Id      int
	SetName string
	Actions []Action
}

type Action struct {
	Id          int
	ActionName  string
	Occurrences []Occurrence
	Sets        []Set
}

type Occurrence struct {
	Id       int
	ActionId int
	Time     time.Time
}
