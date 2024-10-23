package actor

type JWTVkIDTokenData struct {
	Host     string // 'iis': 'VK'
	UserID   int    // 'sub': 000000000
	ClientID int    // 'app': 51812311
	Exired   int64  // 'exp': 1618310760
	TS       int64  // 'iat': 1709018605
	UID      int    // 'jti': 21
}
