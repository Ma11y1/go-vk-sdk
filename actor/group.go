package actor

type Group struct {
	ID          int    `json:"id"`
	AccessToken string `json:"access_token"`
}

func (*Group) GetType() Type {
	return GroupType
}

func (group *Group) GetID() int {
	return group.ID
}

func (group *Group) GetAccessToken() string {
	return group.AccessToken
}

type Groups struct {
	Groups    []Group `json:"groups"`
	ExpiresIn int     `json:"expires_in"`
}

func (*Groups) GetType() Type {
	return GroupsType
}

func (*Groups) GetID() int {
	return -1
}

func (*Groups) GetAccessToken() string {
	return ""
}
