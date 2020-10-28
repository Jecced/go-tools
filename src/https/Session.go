package https

// session环境
// 默认公用Cookie
// 默认公用请求头
// 默认Proxy
// 默认BasicAuth
// 默认超时时间
type session struct {
	comm
	req *request
}

func (s *session) commReq(url string, method RequestType) *session {
	s.req = &request{
		uri:    url,
		method: method,
	}
	return s
}

func (s *session) Get(url string) *session {
	return s.commReq(url, GET)
}

func (s *session) Post(url string) *session {
	return s.commReq(url, POST)
}

func (s *session) GetProxy() string {
	proxy := s.proxy
	if s.req != nil && s.req.proxy != "" {
		proxy = s.req.proxy
	}
	return proxy
}
