package api

const (
	Version                     string = "5.199"
	HostCom                     string = "api.vk.com"
	HostRu                      string = "api.vk.ru"
	MethodEndpoint              string = "https://api.vk.com/method/"
	AuthEndpoint                string = "https://oauth.vk.com/"
	AuthVKIDEndpoint            string = "https://id.vk.com/"
	AuthDirectMethod            string = "token"
	AuthAuthorizeMethod         string = "authorize"
	AuthCodeFlowTokenMethod     string = "access_token"
	AuthVKIDCodeFlowTokenMethod string = "oauth2/auth"
	LanguageDefault             string = "ru"
	RequestLimitUserToken       int    = 3
	RequestLimitGroupToken      int    = 20
	RequestLimitExecuteMethod   int    = 25
)
