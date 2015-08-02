package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//	"reflect"
)

/*
func (db DB) createTables() {
	fmt.Print(reflect.TypeOf(db))
	db.CreateTable(&Set{})
	db.CreateTable(&Action{})
	db.CreateTable(&Occurrence{})

}

func (db DB) dropTables() {
	db.DropTable(&User{})
	db.DropTable(&Set{})
	db.DropTable(&Action{})
	db.DropTable(&Occurrence{})
}
*/
func (db DB) seedTables() {
	seedJSON, err := ioutil.ReadFile("./config/seed.json")
	check(err)
	fmt.Print(string(seedJSON))

	var seed []Action

	json.Unmarshal(seedJSON, &seed)

	fmt.Print(seed)

	for _, action := range seed {
		db.InsertAction(&action)
	}

}

/*
func (db DB) refreshTables() {
	database.dropTables()
	database.createTables()
	database.seedTables()
}
*/
