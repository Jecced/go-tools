package https

// session环境
// 默认公用Cookie
// 默认公用请求头
// 默认Proxy
// 默认BasicAuth
// 默认超时时间
type session struct {
	base
	req *request
}

func (s *session) commReq(url string, method requestType) *session {
	req := &request{
		uri:    url,
		method: method,
	}
	req.header = make(baseParam)
	req.cookie = make(baseParam)
	s.req = req
	return s
}

func (s *session) Get(url string) requests {
	return s.commReq(url, get)
}

func (s *session) Post(url string) requests {
	return s.commReq(url, post)
}

func (s *session) GetProxy() string {
	proxy := s.proxy
	if s.req != nil && s.req.proxy != "" {
		proxy = s.req.proxy
	}
	return proxy
}
