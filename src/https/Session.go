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
}
