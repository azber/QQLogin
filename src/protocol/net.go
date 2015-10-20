package protocol

import "net/http"

func GetCookie(key string, cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == key {
			return cookie.Value
		}
	}
	return ""
}
