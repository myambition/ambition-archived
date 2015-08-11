package main

import (
	"io/ioutil"
)

func (db DB) createTables() {
	database.CreateSetTable()
	database.CreateActionTable()
	database.CreateOccurrenceTable()
}

func (db DB) dropTables() {
	database.DropSetTable()
	database.DropActionTable()
	database.DropActionTable()
}

func (db DB) seedTables() {
	setJson, err := ioutil.ReadFile("./config/sets-seed.json")
	check(err)
	actionJson, err := ioutil.ReadFile("./config/actions-seed.json")
	check(err)
	occurrenceJson, err := ioutil.ReadFile("./config/occurrences-seed.json")
	check(err)

	PostArrayOfSetsJson(setJson)
	PostArrayOfActionsJson(actionJson)
	PostArrayOfOccurrencesJson(occurrenceJson)
}
