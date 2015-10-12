package ambition

import (
	"github.com/julienschmidt/httprouter"
)

// Add routes to http router
// TODO: Add route description parameters and useage
func AddRoutes(router *httprouter.Router) {
	router.GET("/actions", CheckAuth(Actions))
	router.GET("/actions/:ActionId", CheckAuth(ActionById))
	//router.POST("/set/:SetId", PostAction)
	router.GET("/actions/:ActionId/occurrences", Occurrences)
	router.GET("/occurrences/:OccurrenceId", OccurrenceById)

	router.POST("/users", PostUser)

	router.POST("/auth/login", Login)

	// TODO:
	// router.POST("/actions/:ActionId", postOccurrence)
	// router.GET("/sets", sets)
	// router.GET("/sets/:SetId/actions", actionsFromSet)
}
