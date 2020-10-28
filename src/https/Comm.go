package https

import "encoding/base64"

// session request 公用部分
type comm struct {
	// cookie 请求的时候会合并 session 的 cookie
	// 请求结束后服务器返回的 cookie 会合并到 session 部分的 cookie
	cookie BaseParam

	// header 请求的时候回合并 session 的 header
	header BaseParam

	// 如果为空, 则使用 session 的 proxy
	proxy string
	// 如果为空, 则使用 session 的 auth
	auth string

	// 相应建立连接的超时时间
	connTimeout int
	// 相应Resp的超时时间
	respTimeout int
}

func (c *comm) AddCookie(key, value string) CommFace {
	c.cookie.Add(key, value)
	return c
}

func (c *comm) AddCookies(entries map[string]string) CommFace {
	c.cookie.Adds(entries)
	return c
}

func (c *comm) RemoveCookie(key string) CommFace {
	c.cookie.Remove(key)
	return c
}

func (c *comm) ClearCookie() CommFace {
	c.cookie.Clear()
	return c
}

func (c *comm) AddHeader(key, value string) CommFace {
	c.header.Add(key, value)
	return c
}

func (c *comm) AddHeaders(entries map[string]string) CommFace {
	c.header.Adds(entries)
	return c
}

func (c *comm) RemoveHeader(key string) CommFace {
	c.header.Remove(key)
	return c
}

func (c *comm) ClearHeader() CommFace {
	c.header.Clear()
	return c
}

func (c *comm) Proxy(proxy string) CommFace {
	c.proxy = proxy
	return c
}

func (c *comm) BasicAuth(user, password string) CommFace {
	c.auth = "Basic " +
		base64.URLEncoding.EncodeToString([]byte(user+":"+password))
	return c
}

func (c *comm) SetTimeOut(time int) CommFace {
	c.SetConnTimeOut(time)
	c.SetRespTimeOut(time)
	return c
}

func (c *comm) SetConnTimeOut(time int) CommFace {
	c.connTimeout = time
	return c
}

func (c *comm) SetRespTimeOut(time int) CommFace {
	c.respTimeout = time
	return c
}