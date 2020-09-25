package ttlogin

import "github.com/Jecced/rs/src/rs"

func TtLogin(id, secret, code, anonymousCode string) string {
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

	return rs.Get(uri).AddParams(param).Send().ReadText()
}
