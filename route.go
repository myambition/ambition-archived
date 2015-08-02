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
	router.GET("/actions", actions)
	router.GET("/actions/:id", actionById)

	//TODO
	//router.POST("/actions/:id", postOccurrence)
	//router.GET("/sets", sets)
	//router.GET("/sets/:id/actions", actionsFromSet)
	//router.GET("/actions/:id/occurrences", occurrences)
}
