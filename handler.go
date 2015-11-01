package ambition

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserHandler func(http.ResponseWriter, *http.Request, httprouter.Params, *User)

func AuthLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userId int
	var token string
	var user *User

	if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		var err error
		user, token, err = Login(r.FormValue("username"), r.FormValue("password"))
		if err != nil {
			// Page does not seem to refresh when I do http.StatusUnauthorized, need to look into that
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}

		userId = user.Id

	} else {
		fmt.Println("json")
		userJson, err := ioutil.ReadAll(r.Body)
		check(err)

		token, userId, err = LoginUserJson(userJson)
		user, err = database.GetUserById(userId)
	}

	usernameCookie := http.Cookie{Name: "UserId", Value: strconv.Itoa(userId), Path: "/"}
	tokenCookie := http.Cookie{Name: "Token", Value: token, Path: "/"}
	http.SetCookie(w, &usernameCookie)
	http.SetCookie(w, &tokenCookie)
	http.Redirect(w, r, "/", http.StatusFound)

}

func LoginPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p, err := ioutil.ReadFile("./html/login.html")
	check(err)

	fmt.Fprintf(w, string(p))
}

func PostUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userJson, err := ioutil.ReadAll(r.Body)
	check(err)

	err = PostUserJson(userJson)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params, user *User) {
	t, err := template.ParseFiles("./html/index.html")
	actions, err := user.GetActions()
	check(err)
	t.Execute(w, actions)
}

// TODO:
// Remove encoding/json, create passthrough methods in jsonHandler.go if needed

func Actions(w http.ResponseWriter, r *http.Request, _ httprouter.Params, user *User) {
	allActions, err := user.GetActions()
	check(err)

	actionJson, err := json.Marshal(allActions)
	check(err)

	fmt.Fprintf(w, "%s", string(actionJson))
}

func ActionById(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user *User) {
	id, err := strconv.Atoi(ps.ByName("ActionId"))
	check(err)

	actionById, err := user.GetAction(id)
	check(err)

	actionJson, err := json.Marshal(actionById)
	check(err)

	fmt.Fprintf(w, "%s", string(actionJson))
}

func PostAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user *User) {
	actionJson, err := ioutil.ReadAll(r.Body)
	check(err)

	var action Action
	err = json.Unmarshal(actionJson, &action)

	err = user.CreateAction(action)
}

// TODO:
// Add time as a query string parameter. Allow the user to specify how many they want
// also make a week the default amount
func Occurrences(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("ActionId"))
	check(err)

	occurrences, err := database.GetOccurrencesOfAction(id)
	check(err)

	occurrencesJson, err := json.Marshal(occurrences)
	check(err)

	fmt.Fprintf(w, "%s", string(occurrencesJson))

	//	end := time.Now()
	//	//Ok to subtract time in golang you must add negative time. Time.Sub is not the same thing
	//	start := time.Now().Add(-1 * time.Hour * 24 * 7)
	//
	//	q := r.URL.Query()
	//
	//	if len(q.Get("start")) > 0 {
	//		start, _ = time.Parse(time.RFC3339, q.Get("start"))
	//	}
	//
	//	if len(q.Get("end")) > 0 {
	//		end, _ = time.Parse(time.RFC3339, q.Get("end"))
	//	}
	//
	//	database.d.Where(
	//		"Action_ID = ? AND TIME BETWEEN ? AND ?",
	//		ps.ByName("id"), start.Format(time.RFC3339),
	//		end.Format(time.RFC3339)).
	//		Find(&occurrences)
	//
	//	occurrences_json, err := json.Marshal(occurrences)
	//	check(err)
	//
	//	fmt.Fprintf(w, "%s", string(occurrences_json))
}

func PostOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user *User) {
	occurrenceJson, err := ioutil.ReadAll(r.Body)
	check(err)

	var occurrence Occurrence
	err = json.Unmarshal(occurrenceJson, &occurrence)

	actionId, err := strconv.Atoi(ps.ByName("ActionId"))
	check(err)

	action, err := user.GetAction(actionId)
	check(err)
	action.CreateOccurrence(occurrence)

	check(err)
}

func OccurrenceById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("OccurrenceId"))
	check(err)

	occurrence, err := database.GetOccurrenceById(id)
	check(err)

	occurrenceJson, err := json.Marshal(occurrence)
	check(err)

	fmt.Fprintf(w, "%s", string(occurrenceJson))
}
