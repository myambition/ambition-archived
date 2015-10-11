package ambition

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// CheckAuth
func CheckAuth(handle UserHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := r.Cookie("1")
		check(err)

		dbHashedToken, _ := database.GetSessionKeyByUserId(1)
		if CompareHashAndToken(dbHashedToken, token.Value) {
			user, err := database.GetUserById(1)
			check(err)
			handle(w, r, ps, user)
		} else {
			w.WriteHeader(401)
		}
	}

}
