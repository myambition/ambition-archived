package ambition

func (u User) GetActions() ([]Action, error) {
	actions, err := database.GetActionByUserId(u.Id)
	return actions, err
}
