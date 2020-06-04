package cache

var adminToken string
const TokenName = "adminToken"

func SetToken(str string) {
	adminToken = str
}

func GetToken() (string, bool) {
	if adminToken == "" {
		return adminToken, false
	}
	return adminToken, true
}
