package ambition

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// This function runs the command
// $ docker run -e POSTGRES_PASSWORD=ambition -e POSTGRES_USER=ambition -p 5432:5432 -d postgres
// which runs the latest postgres image linking port 5432 and setting the params
// This function should eventually be populated by some kind of config
// I am thinking a json file I open and univerally parse
func dockerDB() {
	out, err := exec.Command("docker", "run", "-e", "POSTGRES_PASSWORD=ambition", "-e", "POSTGRES_USER=ambition", "-p", "5432:5432", "-d", "postgres").Output()
	fmt.Printf("%s", err)
	fmt.Printf("%s", out)

}

func createTables(db DB) {
	err := db.CreateUserTable()
	fmt.Print(err)
	db.CreateSessionTable()
	db.CreateSetTable()
	db.CreateActionTable()
	db.CreateOccurrenceTable()
}

func dropTables(db DB) {
	db.DropUserTable()
	db.DropSessionTable()
	db.DropSetTable()
	db.DropActionTable()
	db.DropOccurrenceTable()
}

func seedTables() {
	//setJson, err := ioutil.ReadFile("../testdata/seed-data/sets-seed.json")
	//check(err)
	actionJson, err := ioutil.ReadFile("./testdata/seed-data/actions-seed.json")
	check(err)
	occurrenceJson, err := ioutil.ReadFile("./testdata/seed-data/occurrences-seed.json")
	check(err)
	userJson, err := ioutil.ReadFile("./testdata/post-data/user.json")
	check(err)

	//PostArrayOfSetsJson(setJson)
	PostArrayOfActionsJson(actionJson)
	PostArrayOfOccurrencesJson(occurrenceJson)
	PostUserJson(userJson)
}

func CallCommand(command string) error {
	switch command {
	case "dockerDB":
		dockerDB()
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
