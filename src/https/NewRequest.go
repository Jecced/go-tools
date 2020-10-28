package https

import (
	"net/http"
	"net/url"
	"strings"
)

func getUrl(uri string, param url.Values) string {
	// 参数列表为空 直接返回
	if 0 == len(param) {
		return uri
	}

	body := param.Encode()

	markWithEnd := strings.HasSuffix(uri, "?")

	// 问号结尾, 直接返回uri + param
	if markWithEnd {
		return uri + body
	}

	hasMark := strings.Index(uri, "?") != -1
	// 有问号, 但是在中间
	if hasMark {
		return uri + "&" + body
	}

	// 没有问号
	return uri + "?" + body
}

// get请求将生成get请求url
func (s *session) getNewRequest() (*http.Request, error) {
	return http.NewRequest(string(s.req.method), getUrl(s.req.uri, s.req.param), nil)
}

// 生成post请求
func (s *session) postNewRequest() (*http.Request, error) {
	return http.NewRequest(string(s.req.method), s.req.uri, strings.NewReader(s.req.param.Encode()))
}
