package wxlogin

import (
	"fmt"
	"github.com/Jecced/go-tools/src/https"
)

func WxLogin(id, secret, code string) (string, error) {
	uri := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(uri, id, secret, code)
	return https.Get(url).Send().ReadText()
}
