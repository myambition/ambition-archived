package main

import (
	"github.com/julienschmidt/httprouter"
)

/*
var getRoutes = map[string]func(){
	"/actions":     actions,
	"/actions/:id": actionById,
}
*/
func AddRoutes(router *httprouter.Router) {
	router.GET("/actions", Actions)
	router.GET("/actions/:ActionId", ActionById)
	router.POST("/set/:SetId", PostAction)

	//TODO
	//router.POST("/actions/:ActionId", postOccurrence)
	//router.GET("/sets", sets)
	//router.GET("/sets/:SetId/actions", actionsFromSet)
	router.GET("/actions/:ActionId/occrurrences", Occurrences)
	router.GET("/occurrences/:OccurrenceId", OccurrenceById)
}
