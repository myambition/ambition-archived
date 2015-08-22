package ambition

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

var _commands = map[string]func(){
	"seed":   database.seedTables,
	"create": database.createTables,
	"drop":   database.dropTables,
}

var port = os.Getenv("ambition_port")

func Run() {
	// database located in db.go
	defer database.Close()

	// Check fof command line arguments
	if len(os.Args) > 1 {
		command := _commands[os.Args[1]]
		if command != nil {
			command()
		} else {
			fmt.Println("Command not found")
		}
	} else { // If no arguments were found
		// Get a router
		router := httprouter.New()

		// Add the routes in route.go
		AddRoutes(router)

		// Start the http server
		http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	}
}
