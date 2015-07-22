package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

var mygorm, _ = gorm.Open("postgres", "user=ambition dbname=ambition password=ambition")

var database = Database{mygorm}

var _commands = map[string]func(){
	"seed":    database.seedTables,
	"create":  database.createTables,
	"drop":    database.dropTables,
	"refresh": database.refreshTables,
}

func main() {
	database.d.LogMode(true)
	database.d.DB()

	if len(os.Args) > 1 {
		command := _commands[os.Args[1]]
		if command != nil {
			command()
		} else {
			fmt.Println("Command not found")
		}
	} else {
		router := httprouter.New()
		router.GET("/", handler)
		router.GET("/actions", actions)
		router.GET("/actions/:id", actionById)
		router.GET("/actions/:id/occurrences", occurrences)
		router.POST("/actions/:id", postOccurrence)

		router.GET("/sets", sets)
		router.GET("/sets/:id/actions", actionsFromSet)

		http.ListenAndServe(":3000", router)
	}
}
