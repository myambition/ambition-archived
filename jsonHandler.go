package main

import (
	"encoding/json"
)

func PostOccurrenceByActionIdJson(ActionId int, occurrenceJson []byte) error {
	var occurrence Occurrence
	err := json.Unmarshal(occurrenceJson, occurrence)

	occurrence.ActionId = ActionId
	database.InsertOccurrence(&occurrence)

	return err
}

func PostActionBySetIdJson(SetId int, actionJson []byte) error {
	var action Action
	err := json.Unmarshal(actionJson, action)

	action.SetId = SetId
	database.InsertAction(&action)

	return err
}

func PostArrayOfSetsJson(setJson []byte) error {
	var sets []Set
	json.Unmarshal(setJson, &sets)
	var err error
	for _, set := range sets {
		err = database.InsertSet(&set)
	}

	return err
}

func PostArrayOfActionsJson(actionJson []byte) error {
	var actions []Action
	json.Unmarshal(actionJson, &actions)
	var err error
	for _, action := range actions {
		err = database.InsertAction(&action)
	}

	return err
}

func PostArrayOfOccurrencesJson(occurrenceJson []byte) error {
	var occurrences []Occurrence
	json.Unmarshal(occurrenceJson, &occurrences)
	var err error
	for _, occurrence := range occurrences {
		err = database.InsertOccurrence(&occurrence)
	}

	return err
}
