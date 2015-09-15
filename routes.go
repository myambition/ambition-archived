package ambition

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Add routes to http router
// TODO: Add route description parameters and useage
func AddRoutes(router *httprouter.Router) {

	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	router.GET("/", Index)
	router.GET("/actions", Actions)
	router.GET("/actions/:ActionId", ActionById)
	router.POST("/set/:SetId", PostAction)
	router.GET("/actions/:ActionId/occurrences", Occurrences)
	router.GET("/occurrences/:OccurrenceId", OccurrenceById)

	router.POST("/actions/:ActionId", PostOccurrence)
	// TODO:
	// router.GET("/sets", sets)
	// router.GET("/sets/:SetId/actions", actionsFromSet)
}
