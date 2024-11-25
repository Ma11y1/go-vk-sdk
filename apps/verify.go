package apps

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/logger"
	"net/http"
	"net/url"
	"strings"
)

// Doc: https://dev.vk.com/ru/mini-apps/development/launch-params

// ParamsVerification represents verification struct.
//
//	Doc: https://dev.vk.com/ru/mini-apps/development/launch-params-sign
type ParamsVerification struct {
	Secret string
}

// Sign return signature in base64.
func (p *ParamsVerification) Sign(data []byte) string {
	mac := hmac.New(sha256.New, []byte(p.Secret))
	_, _ = mac.Write(data)
	expectedMAC := mac.Sum(nil)

	base64Sign := base64.StdEncoding.EncodeToString(expectedMAC)
	base64Sign = strings.ReplaceAll(base64Sign, "+", "-")
	base64Sign = strings.ReplaceAll(base64Sign, "/", "_")
	base64Sign = strings.TrimRight(base64Sign, "=")

	return base64Sign
}

// Verify
//
// link := "https://domen.domen/?user_id=494075&vk_app_id=11111&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA"
//
// pv := vkapps.NewParamsVerification(clientSecret)
//
// v, _ := pv.Verify(link)
func (p *ParamsVerification) Verify(u *url.URL) (bool, error) {
	rawValues, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return false, internalErrors.ErrorLog("Apps.ParamsVerification.Verify()", "Error parsing query string: "+err.Error())
	}

	if len(rawValues["sign"]) == 0 || len(rawValues["sign_keys"]) == 0 {
		return false, nil
	}

	singParams := make(url.Values)
	signKeys := strings.Split(rawValues["sign_keys"][0], ",")

	for _, key := range signKeys {
		values, ok := rawValues[key]
		if !ok {
			continue
		}

		for _, value := range values {
			singParams.Add(key, value)
		}
	}

	base64Sign := p.Sign([]byte(singParams.Encode())) // sign params sorted by key.

	return base64Sign == rawValues["sign"][0], nil
}

// VerifyMiddleware
//
// pv := vkapps.NewParamsVerification(clientSecret)
//
// http.HandleFunc("/api/user/", pv.VerifyMiddleware(UserHandler))
// http.HandleFunc("/api/user/details", pv.VerifyMiddleware(UserDetailsHandler))
func (p *ParamsVerification) VerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok, err := p.Verify(r.URL)
		if err != nil {
			logger.Log("Apps.ParamsVerification.VerifyMiddleware()", "Error 400 Bad request url: "+r.URL.String())
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if ok {
			next.ServeHTTP(w, r)
		} else {
			logger.Log("Apps.ParamsVerification.VerifyMiddleware()", "Error 403 Forbidden url: "+r.URL.String())
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// Verify
//
//	link := "https://domen.domen/?user_id=494075&vk_app_id=11111&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA"
//
//	v, _ := apps.Verify(link, "ewo123SecretMd2390")
func Verify(link, secret string) (bool, error) {
	pv := &ParamsVerification{Secret: secret}

	u, err := url.Parse(link)
	if err != nil {
		return false, internalErrors.ErrorLog("Apps.Verify()", "Error parse link("+link+"): "+err.Error())
	}

	return pv.Verify(u)
}
