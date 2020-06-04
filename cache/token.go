package cache

var token string

func SetToken(str string) {
	token = str
}

func GetToken() (string, bool) {
	if token == "" {
		return token, false
	}
	return token, true
}
