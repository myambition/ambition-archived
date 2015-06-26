package main

import (
	//	"net/http"
	"fmt"
	"os"
)

var _commands = map[string]func(){
	"seed":   seedTables,
	"clear":  clearTables,
	"create": createTables,
	"drop":   dropTables,
}

func main() {
	command := _commands[os.Args[1]]
	if command != nil {
		command()
	} else {
		fmt.Println("Command not found")
	}
	//	http.HandleFunc("/", handler)
	//	http.ListenAndServe(":8081", nil)
}
