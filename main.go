package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

var mygorm, _ = gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")

var database = Database{mygorm}

var _commands = map[string]func(){
	"seed":   database.seedTables,
	"clear":  database.clearTables,
	"create": database.createTables,
	"drop":   database.dropTables,
}

func main() {
	database.d.DB()
	//need to add check that os.Args is greater than 1
	if len(os.Args) > 2 {
		command := _commands[os.Args[1]]
		if command != nil {
			command()
		} else {
			fmt.Println("Command not found")
		}
	}
	router := httprouter.New()
	router.GET("/", handler)

	http.ListenAndServe(":3000", router)
}
