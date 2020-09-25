package wxlogin

import (
	"fmt"
	"github.com/Jecced/rs/src/rs"
)

func WxLogin(id, secret, code string) string {
	uri := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(uri, id, secret, code)
	return rs.Get(url).Send().ReadText()
}
