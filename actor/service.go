package actor

type Service struct {
	AccessToken string `json:"access_token"`
}

func NewService(accessToken string) *Service {
	return &Service{
		AccessToken: accessToken,
	}
}

func (s *Service) GetID() int {
	return -1
}

func (s *Service) GetAccessToken() string {
	return s.AccessToken
}
