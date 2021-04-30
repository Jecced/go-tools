package https

import (
	"net/http"
	"net/url"
)

type session struct {
	comm *base
	req  *base

	uri      string
	param    url.Values
	method   string
	response *http.Response
	close    bool

	err error

	// 已经重试的次数
	retry uint

	usePayload interface{} // 是否使用json payload 格式传输数据
}

func Session() *p1 {
	p := &p1{}
	p.comm = &base{
		cookie: make(param),
		header: make(param),
	}
	return p
}

func Get(url string) *p2 {
	return Session().Get(url)
}

func Post(uri string) *p2 {
	return Session().Post(uri)
}

const (
	get  = "GET"
	post = "POST"
)

// 枚举, 固定请求头
const (
	headerContentType   = "Content-Type"
	headerAuthorization = "Authorization"
	headerCookie        = "Cookie"
)
