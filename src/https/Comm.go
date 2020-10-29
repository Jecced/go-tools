package https

func getAuth(s *session) string {
	auth := s.req.auth
	if "" == auth {
		auth = s.comm.auth
	}
	return auth
}

func getProxy(s *session) string {
	proxy := s.req.proxy
	if "" == proxy {
		proxy = s.comm.proxy
	}
	return proxy
}

func hasRespTimeout(s *session) bool {
	return getRespTimeout(s) != 0
}
func getRespTimeout(s *session) int {
	timeout := s.req.respTimeout
	if 0 == timeout {
		timeout = s.comm.respTimeout
	}
	return timeout
}

func hasConnTimeout(s *session) bool {
	return getConnTimeout(s) != 0
}
func getConnTimeout(s *session) int {
	timeout := s.req.connTimeout
	if 0 == timeout {
		timeout = s.comm.connTimeout
	}
	return timeout
}

// 维护session共享区
func (p *p1) AddHeader(key, value string) *p1 {
	p.comm.header.Add(key, value)
	return p
}
func (p *p1) AddHeaders(entries map[string]string) *p1 {
	p.comm.header.Adds(entries)
	return p
}
func (p *p1) RemoveHeader(key string) *p1 {
	p.comm.header.Remove(key)
	return p
}
func (p *p1) ClearHeader() *p1 {
	p.comm.header.Clear()
	return p
}
func (p *p1) AddCookie(key, value string) *p1 {
	p.comm.cookie.Add(key, value)
	return p
}
func (p *p1) AddCookies(entries map[string]string) *p1 {
	p.comm.cookie.Adds(entries)
	return p
}
func (p *p1) RemoveCookie(key string) *p1 {
	p.comm.cookie.Remove(key)
	return p
}
func (p *p1) ClearCookie() *p1 {
	p.comm.cookie.Clear()
	return p
}
func (p *p1) SetTimeOut(time int) *p1 {
	p.SetRespTimeOut(time)
	p.SetConnTimeOut(time)
	return p
}
func (p *p1) SetConnTimeOut(time int) *p1 {
	p.comm.connTimeout = time
	return p
}
func (p *p1) SetRespTimeOut(time int) *p1 {
	p.comm.respTimeout = time
	return p
}
func (p *p1) Proxy(proxy string) *p1 {
	p.comm.proxy = proxy
	return p
}
func (p *p1) BasicAuth(user, password string) *p1 {
	p.comm.auth = encodeBasicAuth(user, password)
	return p
}

// 维护 request 私有区
func (p *p2) AddHeader(key, value string) *p2 {
	p.comm.header.Add(key, value)
	return p
}
func (p *p2) AddHeaders(entries map[string]string) *p2 {
	p.comm.header.Adds(entries)
	return p
}
func (p *p2) RemoveHeader(key string) *p2 {
	p.comm.header.Remove(key)
	return p
}
func (p *p2) ClearHeader() *p2 {
	p.comm.header.Clear()
	return p
}
func (p *p2) AddCookie(key, value string) *p2 {
	p.comm.cookie.Add(key, value)
	return p
}
func (p *p2) AddCookies(entries map[string]string) *p2 {
	p.comm.cookie.Adds(entries)
	return p
}
func (p *p2) RemoveCookie(key string) *p2 {
	p.comm.cookie.Remove(key)
	return p
}
func (p *p2) ClearCookie() *p2 {
	p.comm.cookie.Clear()
	return p
}
func (p *p2) SetTimeOut(time int) *p2 {
	p.SetRespTimeOut(time)
	p.SetConnTimeOut(time)
	return p
}
func (p *p2) SetConnTimeOut(time int) *p2 {
	p.comm.connTimeout = time
	return p
}
func (p *p2) SetRespTimeOut(time int) *p2 {
	p.comm.respTimeout = time
	return p
}
func (p *p2) Proxy(proxy string) *p2 {
	p.comm.proxy = proxy
	return p
}
func (p *p2) BasicAuth(user, password string) *p2 {
	p.comm.auth = encodeBasicAuth(user, password)
	return p
}
