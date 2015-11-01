package ambition

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Add routes to http router
// TODO: Add route description parameters and useage
func AddRoutes(router *httprouter.Router) {
	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	router.GET("/", CheckAuth(Index))
	router.GET("/actions", CheckAuth(Actions))
	router.GET("/actions/:ActionId", CheckAuth(ActionById))
	//router.POST("/set/:SetId", PostAction)
	router.GET("/actions/:ActionId/occurrences", Occurrences)
	router.GET("/occurrences/:OccurrenceId", OccurrenceById)

	router.POST("/users", PostUser)
	router.GET("/auth/login", LoginPage)
	router.POST("/auth/login", AuthLogin)

	router.POST("/actions/:ActionId", CheckAuth(PostOccurrence))
	router.POST("/actions/", CheckAuth(PostAction))

	// TODO:
	// router.GET("/sets", sets)
	// router.GET("/sets/:SetId/actions", actionsFromSet)
}
