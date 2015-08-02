package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

var tempdb, _ = sql.Open("postgres", "dbname=ambition user=ambition password=ambition sslmode=disable")

var database = DB{tempdb}

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

func (db DB) GetActionById(id int) (*Action, error) {
	const query = `SELECT action_name from actions where id = $1`
	var reval Action
	err := db.QueryRow(query, id).Scan(&reval.ActionName)
	check(err)
	reval.Id = id
	return &reval, err
}

func (db DB) InsertAction(action *Action) error {
	const query = `INSERT INTO actions (action_name) VALUES ($1)`

	_, err := db.Exec(query, action.ActionName)
	check(err)

	return err
}

func (db DB) GetActions() ([]Action, error) {
	const query = `SELECT * from actions`
	var reval []Action

	rows, err := db.Query(query)
	check(err)
	defer rows.Close()
	for rows.Next() {
		var action Action
		err := rows.Scan(&action.Id, &action.ActionName)
		check(err)
		reval = append(reval, action)
	}
	return reval, err

}

func (db DB) DropActionTable() error {
	const query = `DROP TABLE actions`

	_, err := db.Exec(query)

	return err
}

func (db DB) CreateActionTable() error {
	const query = `CREATE TABLE actions(id SERIAL PRIMARY KEY, action_name varchar(255))`

	_, err := db.Exec(query)

	return err
}
