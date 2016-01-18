package ambition

import (
	"errors"
)

func Login(username, password string) (u *User, token string, err error) {

	user, err := database.GetUserByUserName(username)

	authed := CompareSaltAndPasswordToHash(user.PasswordSalt, user.HashedPassword, []byte(password))

	if authed {
		token, _ := GenerateRandomString(32)
		database.InsertSession(user.Id, HashToken(token))
		return user, token, err
	}

	return nil, "", errors.New("Invalid username or password")
}

func (u User) GetActions() ([]Action, error) {
	actions, err := database.GetActionsByUserId(u.Id)
	return actions, err
}

func (u User) GetAction(actionId int) (*Action, error) {
	action, err := database.GetActionById(actionId)
	check(err)
	if action.UserId == u.Id {
		return action, err
	}
	return nil, errors.New("Permission Denied")
}

func (a Action) CreateOccurrence(occurrence Occurrence) error {
	occurrence.ActionId = a.Id

	err := database.InsertOccurrence(&occurrence)
	return err
}
