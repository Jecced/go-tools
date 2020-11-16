package ttlogin

import (
	"github.com/Jecced/go-tools/src/https"
)

func TtLogin(id, secret, code, anonymousCode string) (string, error) {
	uri := "https://developer.toutiao.com/api/apps/jscode2session"
	param := map[string]string{
		"appid":  id,
		"secret": secret,
	}
	if code != "" {
		param["code"] = code
	}
	if anonymousCode != "" {
		param["anonymous_code"] = anonymousCode
	}

	return https.Get(uri).AddParams(param).Send().ReadText()
}
