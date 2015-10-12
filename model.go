package ambition

import "errors"

func (u User) GetActions() ([]Action, error) {
	actions, err := database.GetActionsByUserId(u.Id)
	return actions, err
}

func (u User) GetAction(actionId int) (*Action, error) {
	action, err := database.GetActionById(actionId)
	if action.UserId == u.Id {
		return action, err
	}
	return nil, errors.New("Permission Denied")
}
