package main

import (
	"io/ioutil"
	//	"reflect"
)

func (db DB) createTables() {
	database.CreateActionTable()
	database.CreateOccurrenceTable()
}

func (db DB) dropTables() {
	database.DropActionTable()
	database.DropActionTable()
}

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
