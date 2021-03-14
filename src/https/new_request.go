package https

import (
	"bytes"
	"encoding/json"
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
func (p *p2) getNewRequest() (*http.Request, error) {
	return http.NewRequest(p.method, getUrl(p.uri, p.param), nil)
}

// 生成post请求
func (p *p2) postNewRequest() (*http.Request, error) {
	//return http.NewRequest(p.method, p.uri, strings.NewReader(p.param.Encode()))

	var jsonBytes []byte

	var err error
	if p.usePayload {
		jsonBytes, err = json.Marshal(p.param)
		if err != nil {
			return nil, err
		}
	} else {
		jsonBytes = []byte(p.param.Encode())
	}

	return http.NewRequest(p.method, p.uri, bytes.NewBuffer(jsonBytes))
}
