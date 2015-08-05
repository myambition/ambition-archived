package main

func PostOccurrenceByActionIdJson(ActionId int, occurrenceJson []byte) error {
	occurrence, err := UnmarshalOccurrence(occurrenceJson)
	database.InsertOccurrenceOfAction(ActionId, occurrence)

	return err
}
