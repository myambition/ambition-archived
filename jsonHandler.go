package ambition

import (
	"encoding/json"
	"fmt"
)

func LoginUserJson(userJson []byte) (string, int, error) {
	var userJsonMap map[string]interface{}

	err := json.Unmarshal(userJson, &userJsonMap)
	fmt.Println(userJsonMap)

	user, _ := database.GetUserByUserName(userJsonMap["username"].(string))

	passwordSalt := user.PasswordSalt
	password := []byte(userJsonMap["password"].(string))
	hashedPassword := user.HashedPassword
	auth := CompareSaltAndPasswordToHash(passwordSalt, hashedPassword, password)
	if auth {
		token, _ := GenerateRandomString(32)
		database.InsertSession(user.Id, HashToken(token))
		return token, user.Id, nil
	}

	return "", 0, err
}

// Creates a User in the database from json.
func PostUserJson(userJson []byte) error {
	var userJsonMap map[string]interface{}

	err := json.Unmarshal(userJson, &userJsonMap)

	var user User
	password := userJsonMap["password"].(string)
	user.UserName = userJsonMap["username"].(string)
	user.Email = userJsonMap["email"].(string)

	// Create User salt and hashed password for storage and later authentication
	user.PasswordSalt, user.HashedPassword, err = CreateSaltAndHashedPassword([]byte(password))

	database.InsertUser(&user)

	return err
}

func PostOccurrenceByActionIdJson(ActionId int, occurrenceJson []byte) error {
	var occurrence Occurrence
	err := json.Unmarshal(occurrenceJson, &occurrence)

	occurrence.ActionId = ActionId
	database.InsertOccurrence(&occurrence)

	return err
}

func PostActionBySetIdJson(SetId int, actionJson []byte) error {
	var action Action
	err := json.Unmarshal(actionJson, &action)

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
