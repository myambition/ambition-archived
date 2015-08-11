package main

import (
	"github.com/julienschmidt/httprouter"
)

// Add routes to http router
// TODO: Add route description parameters and useage
func AddRoutes(router *httprouter.Router) {
	router.GET("/actions", Actions)
	router.GET("/actions/:ActionId", ActionById)
	router.POST("/set/:SetId", PostAction)
	router.GET("/actions/:ActionId/occrurrences", Occurrences)
	router.GET("/occurrences/:OccurrenceId", OccurrenceById)

	// TODO:
	// router.POST("/actions/:ActionId", postOccurrence)
	// router.GET("/sets", sets)
	// router.GET("/sets/:SetId/actions", actionsFromSet)
}
