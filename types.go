package main

import (
	"time"
)

type Event struct {
	ID         int
	EventName  string
	EventTimes []EventTime
}

type EventTime struct {
	ID      int
	EventID int
	Time    time.Time
}
