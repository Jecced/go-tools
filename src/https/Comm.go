package https

type Cookie BaseParam
type Header BaseParam

// session request 公用部分
type comm struct {
	// cookie 请求的时候会合并 session 的 cookie
	// 请求结束后服务器返回的 cookie 会合并到 session 部分的 cookie
	cookie Cookie

	// header 请求的时候回合并 session 的 header
	header Header

	// 如果为空, 则使用 session 的 proxy
	proxy string
	// 如果为空, 则使用 session 的 auth
	auth string

	// 相应建立连接的超时时间
	connTimeout int
	// 相应Resp的超时时间
	respTimeout int
}
