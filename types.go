package main

import (
	"time"
)

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type Set struct {
	Id      int    `json:"id"`
	SetName string `json:"setName"`
}

type Action struct {
	Id         int    `json:"id"`
	ActionName string `json:"actionName"`
	SetId      int    `json:"setId"`
}

type Occurrence struct {
	Id       int       `json:"id"`
	ActionId int       `json:"actionId"`
	Time     time.Time `json:"time"`
}
