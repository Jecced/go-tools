package https

import "net/url"

type https struct {
	uri string

	param url.Values
}

// session环境
// 共享Cookie
// 共享请求头
// 共享Proxy
// 共享BasicAuth
// 共享超时时间
type session struct {
	cookie Param
	header Param
}
