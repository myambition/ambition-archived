package ambition

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var port = os.Getenv("ambition_port")

func Run() {
	// Get a router
	router := httprouter.New()

	// Add the routes in route.go
	AddRoutes(router)

	// Start the http server
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
