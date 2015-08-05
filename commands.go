package main

import (
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

//Fix this method to use jsonHandler.go
func (db DB) seedTables() {
	actionJson, err := ioutil.ReadFile("./config/actions-seed.json")
	check(err)
	occurrenceJson, err := ioutil.ReadFile("./config/occurrences-seed.json")
	check(err)

	PostArrayOfActionsJson(actionJson)
	PostArrayOfOccurrencesJson(occurrenceJson)

}

/*
func (db DB) refreshTables() {
	database.dropTables()
	database.createTables()
	database.seedTables()
}
*/
