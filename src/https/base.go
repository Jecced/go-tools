package https

type base struct {
	// cookie 请求的时候会合并 session 的 cookie
	// 请求结束后服务器返回的 cookie 会合并到 session 部分的 cookie
	cookie param

	// cookie 序列化 反序列化地址 序列化到一个本地文件
	cookieSerializationPath string

	// header 请求的时候回合并 session 的 header
	header param

	// 如果为空, 则使用 session 的 proxy
	proxy string
	// 如果为空, 则使用 session 的 auth
	auth string

	// 相应建立连接的超时时间
	connTimeout int
	// 相应Resp的超时时间
	respTimeout int

	// 是否跳过https校验
	skipSSLVerify bool

	// 重试次数
	retry uint
}
