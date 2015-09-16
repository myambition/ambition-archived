package ambition

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CheckAuth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := r.Cookie("1")
		check(err)

		dbHashedToken, _ := database.GetSessionKeyByUserId(1)
		if CompareHashAndToken(dbHashedToken, token.Value) {
			handle(w, r, ps)
		} else {
			w.WriteHeader(401)
		}
	}

}
