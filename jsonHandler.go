package main

import (
	"encoding/json"
)

func PostOccurrenceByActionIdJson(ActionId int, occurrenceJson []byte) error {
	occurrence, err := UnmarshalOccurrence(occurrenceJson)
	database.InsertOccurrenceOfAction(ActionId, occurrence)

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
