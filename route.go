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
	router.GET("/actions/:ActionId", actionById)

	//TODO
	//router.POST("/actions/:ActionId", postOccurrence)
	//router.GET("/sets", sets)
	//router.GET("/sets/:SetId/actions", actionsFromSet)
	router.GET("/actions/:ActionId/occurrences", occurrences)
}
