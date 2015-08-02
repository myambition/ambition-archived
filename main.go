package main

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

var tempdb, _ = sql.Open("postgres", "dbname=ambition user=ambition password=ambition sslmode=disable")

var database = DB{tempdb}

var _commands = map[string]func(){
	"seed": database.seedTables,
	//"create":  database.createTables,
	//"drop":    database.dropTables,
	//"refresh": database.refreshTables,
}

func main() {

	//defer tempdb.Close()

	if len(os.Args) > 1 {
		command := _commands[os.Args[1]]
		if command != nil {
			command()
		} else {
			fmt.Println("Command not found")
		}
	} else {
		router := httprouter.New()
		//router.GET("/", handler)
		//router.GET("/actions", actions)
		router.GET("/actions/:id", actionById)
		//router.GET("/actions/:id/occurrences", occurrences)
		//router.POST("/actions/:id", postOccurrence)

		//router.GET("/sets", sets)
		//router.GET("/sets/:id/actions", actionsFromSet)

		http.ListenAndServe(":3000", router)

	}
}
