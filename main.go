package main

import (
	//	"net/http"
	"fmt"
	"github.com/jinzhu/gorm"
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
	command := _commands[os.Args[1]]
	if command != nil {
		command()
	} else {
		fmt.Println("Command not found")
	}
	//	http.HandleFunc("/", handler)
	//	http.ListenAndServe(":8081", nil)
}
