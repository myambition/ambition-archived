package main

import (
	"net/http"
	"os"
)

var _commands = map[string]func(){
	"seed":   seedDatabase,
	"clear":  clearDatabase,
	"create": createDatabase,
}

func main() {
	_commands[os.Args[1]]()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
