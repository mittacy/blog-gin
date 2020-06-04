package cache

var token string

func CacheToken(str string) {
	token = str
}

func GetCacheToken() (string, bool) {
	if token == "" {
		return token, false
	}
	return token, true
}
