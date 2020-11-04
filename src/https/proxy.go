package https

import (
	"net/http"
	"net/url"
)

// 给 translate 设置代理
func (s *session) setProxy(transport *http.Transport) {
	if s.GetProxy() != "" {
		transport.Proxy = s.proxyFun
	}
}

// 代理方法
func (s *session) proxyFun(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://" + s.GetProxy())
}
