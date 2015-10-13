package ambition

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// CheckAuth
func CheckAuth(handle UserHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userIdCookie, err := r.Cookie("UserId")
		if err != nil {
			LoginPage(w, r, ps)
			return
		}
		userId, err := strconv.Atoi(userIdCookie.Value)

		token, err := r.Cookie("Token")

		dbHashedToken, _ := database.GetSessionKeyByUserId(userId)
		if CompareHashAndToken(dbHashedToken, token.Value) {
			user, err := database.GetUserById(userId)
			check(err)
			handle(w, r, ps, user)
		} else {
			w.WriteHeader(401)
		}
	}
}
