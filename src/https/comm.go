package https

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
func (p *p1) Retry(count uint) *p1 {
	p.comm.retry = count
	return p
}

func (p *p1) SkipSSLVerify(verify bool) *p1{
	p.comm.skipSSLVerify = verify
	return p
}

// 维护 request 私有区
func (p *p2) AddHeader(key, value string) *p2 {
	p.req.header.Add(key, value)
	return p
}
func (p *p2) AddHeaders(entries map[string]string) *p2 {
	p.req.header.Adds(entries)
	return p
}
func (p *p2) RemoveHeader(key string) *p2 {
	p.req.header.Remove(key)
	return p
}
func (p *p2) ClearHeader() *p2 {
	p.req.header.Clear()
	return p
}
func (p *p2) AddCookie(key, value string) *p2 {
	p.req.cookie.Add(key, value)
	return p
}
func (p *p2) AddCookies(entries map[string]string) *p2 {
	p.req.cookie.Adds(entries)
	return p
}
func (p *p2) RemoveCookie(key string) *p2 {
	p.req.cookie.Remove(key)
	return p
}
func (p *p2) ClearCookie() *p2 {
	p.req.cookie.Clear()
	return p
}
func (p *p2) SetTimeOut(time int) *p2 {
	p.SetRespTimeOut(time)
	p.SetConnTimeOut(time)
	return p
}
func (p *p2) SetConnTimeOut(time int) *p2 {
	p.req.connTimeout = time
	return p
}
func (p *p2) SetRespTimeOut(time int) *p2 {
	p.req.respTimeout = time
	return p
}
func (p *p2) Proxy(proxy string) *p2 {
	p.req.proxy = proxy
	return p
}
func (p *p2) BasicAuth(user, password string) *p2 {
	p.req.auth = encodeBasicAuth(user, password)
	return p
}
func (p *p2) Retry(count uint) *p2 {
	p.req.retry = count
	return p
}
func (p *p2) SkipSSLVerify(verify bool) *p2{
	p.req.skipSSLVerify = verify
	return p
}