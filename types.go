package ambition

import (
	"time"
)

type User struct {
	Id             int    `json:"id"`
	UserName       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword []byte
	PasswordSalt   []byte
}

type Set struct {
	Id      int    `json:"id"`
	SetName string `json:"setName"`
}

// TODO: Add metadata / extradata / data field. It can have any structure.
//		 Action specifies the structure.
type Action struct {
	Id         int    `json:"id"`
	ActionName string `json:"actionName"`
	SetId      int    `json:"setId"`
}

// TODO: Add metadata / extradata / data field. It can have any structure.
//		 Occurrence should follow the sturcture specified in Action.
type Occurrence struct {
	Id       int       `json:"id"`
	ActionId int       `json:"actionId"`
	Time     time.Time `json:"time"`
}
