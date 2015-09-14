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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dat, err := ioutil.ReadFile("./html/index.html")
	_ = dat
	check(err)
	t, err := template.ParseFiles("./html/index.html")
	actions, err := database.GetActions()
	t.Execute(w, actions)
}

// TODO:
// Remove encoding/json, create passthrough methods in jsonHandler.go if needed

func Actions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allActions, err := database.GetActions()
	check(err)

	actionJson, err := json.Marshal(allActions)
	check(err)

	fmt.Fprintf(w, "%s", string(actionJson))
}

func ActionById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("ActionId"))
	check(err)

	actionById, err := database.GetActionById(id)
	check(err)

	actionJson, err := json.Marshal(actionById)
	check(err)

	fmt.Fprintf(w, "%s", string(actionJson))
}

func PostAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actionJson, err := ioutil.ReadAll(r.Body)
	check(err)

	id, err := strconv.Atoi(ps.ByName("SetId"))
	check(err)

	err = PostActionBySetIdJson(id, actionJson)
	check(err)
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

func PostOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	occurrenceJson, err := ioutil.ReadAll(r.Body)
	check(err)

	id, err := strconv.Atoi(ps.ByName("ActionId"))
	check(err)

	err = PostOccurrenceByActionIdJson(id, occurrenceJson)
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
