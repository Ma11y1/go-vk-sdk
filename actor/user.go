package actor

type User struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ID          int    `json:"user_id"`
	Email       string `json:"email,omitempty"`
	State       string `json:"state,omitempty"`
}

func (*User) GetType() Type {
	return UserType
}

func (user *User) GetID() int {
	return user.ID
}

func (user *User) GetAccessToken() string {
	return user.AccessToken
}

type UserVKID struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ID           int    `json:"user_id"`
	IDToken      string `json:"id_token"`
	Scope        int    `json:"scope"`
}

func (*UserVKID) GetType() Type {
	return UserVKIDType
}

func (user *UserVKID) GetID() int {
	return user.ID
}

func (user *UserVKID) GetAccessToken() string {
	return user.AccessToken
}
