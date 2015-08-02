package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

var _commands = map[string]func(){
	"seed": database.seedTables,
	//"create":  database.createTables,
	//"drop":    database.dropTables,
	//"refresh": database.refreshTables,
}

func main() {
	defer database.Close()
	if len(os.Args) > 1 {
		command := _commands[os.Args[1]]
		if command != nil {
			command()
		} else {
			fmt.Println("Command not found")
		}
	} else {
		router := httprouter.New()

		AddRoutes(router)

		http.ListenAndServe(":3000", router)
	}
}
