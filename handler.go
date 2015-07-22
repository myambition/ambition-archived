package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ambition!")
}

func actions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	all_actions := []Action{}

	database.d.Find(&all_actions)
	fmt.Println(all_actions)

	actions_json, err := json.Marshal(all_actions)
	check(err)

	fmt.Fprintf(w, "%s", string(actions_json))
}

func actionById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	action_by_id := Action{}

	database.d.Where("id = ?", ps.ByName("id")).Find(&action_by_id)

	action_json, err := json.Marshal(action_by_id)
	check(err)

	fmt.Fprintf(w, "%s", string(action_json))
}

func occurrences(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	occurrences := []Occurrence{}

	end := time.Now()
	//Ok to subtract time in golang you must add negative time. Time.Sub is not the same thing
	start := time.Now().Add(-1 * time.Hour * 24 * 7)

	q := r.URL.Query()

	if len(q.Get("start")) > 0 {
		start, _ = time.Parse(time.RFC3339, q.Get("start"))
	}

	if len(q.Get("end")) > 0 {
		end, _ = time.Parse(time.RFC3339, q.Get("end"))
	}

	database.d.Where(
		"Action_ID = ? AND TIME BETWEEN ? AND ?",
		ps.ByName("id"), start.Format(time.RFC3339),
		end.Format(time.RFC3339)).
		Find(&occurrences)

	occurrences_json, err := json.Marshal(occurrences)
	check(err)

	fmt.Fprintf(w, "%s", string(occurrences_json))
}

func postOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	occurrenceJSON, err := ioutil.ReadAll(r.Body)
	check(err)

	var occurence Occurrence

	json.Unmarshal(occurrenceJSON, &occurence)
	i, err := strconv.Atoi(ps.ByName("id"))
	check(err)

	occurence.ActionID = i

	database.d.Create(&occurence)
}

func sets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	all_sets := []Set{}

	database.d.Find(&all_sets)

	sets_json, err := json.Marshal(all_sets)
	check(err)

	fmt.Fprintf(w, "%s", sets_json)
}

func actionsFromSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actions := []Action{}

	database.d.Where("Action_ID = ?", ps.ByName("id")).Find(&actions)

	actions_json, err := json.Marshal(actions)
	check(err)

	fmt.Fprintf(w, "%s", actions_json)
}
