package ambition

import (
	"errors"
	"io/ioutil"
)

func createTables(db DB) {
	db.CreateSetTable()
	db.CreateActionTable()
	db.CreateOccurrenceTable()
}

func dropTables(db DB) {
	db.DropSetTable()
	db.DropActionTable()
	db.DropOccurrenceTable()
}

func seedTables() {
	setJson, err := ioutil.ReadFile("../testdata/seed-data/sets-seed.json")
	check(err)
	actionJson, err := ioutil.ReadFile("../testdata/seed-data/actions-seed.json")
	check(err)
	occurrenceJson, err := ioutil.ReadFile("../testdata/seed-data/occurrences-seed.json")
	check(err)

	PostArrayOfSetsJson(setJson)
	PostArrayOfActionsJson(actionJson)
	PostArrayOfOccurrencesJson(occurrenceJson)
}

func CallCommand(command string) error {
	switch command {
	case "seed":
		seedTables()
	case "create":
		createTables(database)
	case "drop":
		dropTables(database)
	default:
		return errors.New("Command Not Found")
	}

	return nil
}
