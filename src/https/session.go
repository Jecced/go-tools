package https

import "net/url"

// 阶段一
// 维护session共享区
// Get(url string) requests
// Post(url string) requests
type p1 session

func (p *p1) Get(uri string) *p2 {
	p.newComm(get, uri)
	return (*p2)(p)
}

func (p *p1) Post(uri string) *p2 {
	p.newComm(post, uri)
	return (*p2)(p)
}

func (p *p1) newComm(method, uri string) {
	p.method = method
	p.uri = uri
	p.param = make(url.Values)
	p.req = &base{
		cookie: make(param),
		header: make(param),
	}
	p.err = nil

	if p.close == false && p.response != nil && !p.response.Close {
		_ = p.response.Body.Close()
	}
	p.close = false
	p.retry = 0
}

// session本身的通用获取
// 优先从私有区获取, 然后从共有区获取
func (s *session) GetConnTimeout() int {
	timeout := s.req.connTimeout
	if 0 == timeout {
		timeout = s.comm.connTimeout
	}
	return timeout
}

func (s *session) GetProxy() string {
	proxy := s.req.proxy
	if "" == proxy {
		proxy = s.comm.proxy
	}
	return proxy
}

func (s *session) HasRespTimeout() bool {
	return s.GetRespTimeout() != 0
}

func (s *session) GetRespTimeout() int {
	timeout := s.req.respTimeout
	if 0 == timeout {
		timeout = s.comm.respTimeout
	}
	return timeout
}

func (s *session) GetRetry() int {
	count := s.req.retry
	if 0 == count {
		count = s.comm.retry
	}
	return int(count)
}

func (s *session) GetAuth() string {
	auth := s.req.auth
	if "" == auth {
		auth = s.comm.auth
	}
	return auth
}

func (p *p1) GetCookies() param {
	return p.comm.cookie
}

func (p *p1) GetCookie(key string) string {
	return p.comm.cookie[key]
}
