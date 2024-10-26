package actor

type Service struct {
	AccessToken string `json:"access_token"`
}

func (*Service) GetType() Type {
	return ServiceType
}

func (*Service) GetID() int {
	return -1
}

func (s *Service) GetAccessToken() string {
	return s.AccessToken
}
