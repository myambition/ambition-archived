package ambition

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
)

var port int

func Run() {
	// Get a router
	router := httprouter.New()

	// Add the routes in route.go
	AddRoutes(router)

	// Start the http server
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func Init() {
	config := ReadConfiguration("./config.json")

	var dbString string

	if config.DBLocal {
		dbString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			config.DBUser, config.DBPassword, config.DBName, config.DBSSL)
	} else {
		dbString = fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=%s",
			config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.DBSSL)
	}

	var tempdb, _ = sql.Open("postgres", dbString)
	database = DB{tempdb}
	port = config.AmbitionPort
}
