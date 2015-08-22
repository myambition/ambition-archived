package ambition

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Database type to extend with custom functions
type DB struct {
	*sql.DB
}

// Database initialization
// TODO: Add seting up postgresql
var tempdb, _ = sql.Open("postgres", "dbname=ambition user=ambition password=ambition sslmode=disable")

// Create a database type to extend
var database = DB{tempdb}

// ----------------------------- Sets  ----------------------------- //
func (db DB) GetSets() ([]Set, error) {
	const query = `SELECT * FROM sets`
	var reval []Set

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var set Set
		err := rows.Scan(&set.Id, &set.SetName)
		check(err)
		reval = append(reval, set)
	}
	return reval, err
}

func (db DB) GetSetById(id int) (*Set, error) {
	const query = `SELECT set_name FROM sets WHERE id = $1`
	var reval Set
	err := db.QueryRow(query, id).Scan(&reval.SetName)
	reval.Id = id
	return &reval, err
}

func (db DB) InsertSet(set *Set) error {
	const query = `INSERT INTO sets (set_name) VALUES ($1)`

	_, err := db.Exec(query, set.SetName)

	return err
}

func (db DB) DeleteSetById(setId int) error {
	const query = `DELETE FROM sets WHERE id = $1`

	_, err := db.Exec(query, setId)

	return err
}

// ----------------------------- Actions  ----------------------------- //

func (db DB) GetActions() ([]Action, error) {
	const query = `SELECT * FROM actions`
	var reval []Action

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var action Action
		err := rows.Scan(&action.Id, &action.ActionName)
		check(err)
		reval = append(reval, action)
	}
	return reval, err
}

func (db DB) GetActionById(id int) (*Action, error) {
	const query = `SELECT action_name, set_id FROM actions WHERE id = $1`
	var reval Action
	err := db.QueryRow(query, id).Scan(&reval.ActionName, &reval.SetId)
	reval.Id = id

	return &reval, err
}

func (db DB) InsertAction(action *Action) error {
	const query = `INSERT INTO actions (action_name, set_id) VALUES ($1, $2))`

	_, err := db.Exec(query, action.ActionName, action.SetId)

	return err
}

func (db DB) DeleteActionById(actionId int) error {
	const query = `DELETE FROM actions WHERE id = $1`

	_, err := db.Exec(query, actionId)

	return err
}

// ----------------------------- Occurrences  ----------------------------- //

func (db DB) GetOccurrenceById(id int) (*Occurrence, error) {
	const query = `SELECT (action_name, time) FROM occurrences WHERE id = $1`
	var reval Occurrence
	err := db.QueryRow(query, id).Scan(&reval.ActionId, &reval.Time)
	reval.Id = id
	return &reval, err
}

func (db DB) GetOccurrencesOfAction(id int) ([]Occurrence, error) {
	const query = `SELECT * FROM occurrences WHERE action_id = $1`
	var reval []Occurrence

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var occurrence Occurrence
		err := rows.Scan(&occurrence.Id, &occurrence.ActionId, &occurrence.Time)
		check(err)
		reval = append(reval, occurrence)
	}
	return reval, err
}

func (db DB) InsertOccurrence(occurrence *Occurrence) error {
	const query = `INSERT INTO occurrences (action_id, time) VALUES ($1, $2)`

	_, err := db.Exec(query, occurrence.ActionId, occurrence.Time)

	return err
}

func (db DB) DeleteOccurrenceById(occurrenceId int) error {
	const query = `DELETE FROM occurrences WHERE id = $1`

	_, err := db.Exec(query, occurrenceId)

	return err
}

// ------------ Table Creation and Dropping -------------------

func (db DB) CreateSetTable() error {
	const query = `CREATE TABLE sets(id SERIAL PRIMARY KEY, set_name varchar(255))`

	_, err := db.Exec(query)

	return err
}

func (db DB) DropSetTable() error {
	const query = `DROP TABLE sets`

	_, err := db.Exec(query)

	return err
}

func (db DB) CreateActionTable() error {
	const query = `CREATE TABLE actions(id SERIAL PRIMARY KEY, action_name varchar(255))`

	_, err := db.Exec(query)

	return err
}

func (db DB) DropActionTable() error {
	const query = `DROP TABLE actions`

	_, err := db.Exec(query)

	return err
}

func (db DB) CreateOccurrenceTable() error {
	const query = `CREATE TABLE occurrences(id SERIAL PRIMARY KEY, action_id varchar(255), time timestamp)`

	_, err := db.Exec(query)

	return err
}

func (db DB) DropOccurrenceTable() error {
	const query = `DROP TABLE occurrences`

	_, err := db.Exec(query)

	return err
}

// FUTURE:
// Will allow combining CreateTable and DropTable functions
func getTable(obj interface{}) string {
	switch obj.(type) {
	default:
		return "unknown"
	case Action, *Action:
		return "actions"
	case Set, *Set:
		return "sets"
	}
}
