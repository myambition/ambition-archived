package ambition

import (
	"encoding/json"
	"errors"
)

// LoginuserJson takes json with a username and a password key
// unmarshals this json and then creates a session if the
// pass authentication information is valid
func LoginUserJson(userJson []byte) (string, int, error) {
	var userJsonMap map[string]interface{}

	err := json.Unmarshal(userJson, &userJsonMap)

	user, token, err := Login(userJsonMap["username"].(string), userJsonMap["password"].(string))

	return token, user.Id, err

}

// Creates a User in the database from json.
func PostUserJson(userJson []byte) error {
	var userJsonMap map[string]interface{}
	err := json.Unmarshal(userJson, &userJsonMap)

	var user User
	password := userJsonMap["password"].(string)
	user.UserName = userJsonMap["username"].(string)
	user.Email = userJsonMap["email"].(string)

	_, err = database.GetUserByUserName(user.UserName)
	if err != nil {
		// Create User salt and hashed password for storage and later authentication
		user.PasswordSalt, user.HashedPassword, err = CreateSaltAndHashedPassword([]byte(password))

		database.InsertUser(&user)
	} else {
		err = errors.New("Username exists")
	}

	return err
}

func PostOccurrenceByActionIdJson(ActionId int, occurrenceJson []byte) error {
	var occurrence Occurrence
	err := json.Unmarshal(occurrenceJson, &occurrence)

	occurrence.ActionId = ActionId
	database.InsertOccurrence(&occurrence)

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
