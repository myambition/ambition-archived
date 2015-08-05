package main

import (
	"encoding/json"
)

func UnmarshalAction(actionJson []byte) (*Action, error) {
	var action Action
	err := json.Unmarshal(actionJson, action)

	return &action, err
}

func UnmarshalOccurrence(occurrenceJson []byte) (*Occurrence, error) {
	var occurrence Occurrence
	err := json.Unmarshal(occurrenceJson, occurrence)

	return &occurrence, err
}
