package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hi there, I loves %s!", r.URL.Path[1:])
}

func actions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	all_actions := []Action{}

	database.d.Find(&all_actions)
	fmt.Println(all_actions)

	actions_json, err := json.Marshal(all_actions)
	check(err)

	fmt.Fprintf(w, "%s", string(actions_json))
}

func action_id(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	action_by_id := Action{}

	database.d.Find(&action_by_id).Where("id = ?", ps.ByName("id"))

	action_json, err := json.Marshal(action_by_id)
	check(err)

	fmt.Fprintf(w, "%s", string(action_json))
}

func post_occurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	occurrenceJSON, err := ioutil.ReadAll(r.Body)
	check(err)

	var occurence Occurrence

	json.Unmarshal(occurrenceJSON, &occurence)
	i, err := strconv.Atoi(ps.ByName("id"))
	check(err)

	occurence.ActionID = i

	database.d.Create(&occurence)
}
