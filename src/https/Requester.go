package https

type Requester interface {
	// 增加一个参数
	AddParam(key, value string)
	// 增加多个参数
	AddParams(param map[string]string)
	// 移除一个参数
	RemoveParam(key string)
	// 清空所有参数
	ClearParam()

	// 增加一个请求头
	AddHeader(key, value string)
	// 增加多个请求头
	AddHeaders(param map[string]string)
	// 移除一个请求头
	RemoveHeader(key string)
	// 清空所有请求头
	ClearHeader()

	// 增加一个cookie
	AddCookie(key, value string)
	// 增加多个cookie
	AddCookies(param map[string]string)
	// 移除一个cookie
	RemoveCookie(key string)
	// 清空所有cookie
	ClearCookie()

	// 设置超时时间
	SetTimeOut(time int)
	// 设置连接超时时间
	SetConnTimeOut(time int)
	// 设置反馈超时时间
	SetRespTimeOut(time int)

	// 代理
	Proxy(proxy string)

	// 基础认证
	BasicAuth(user, password string)
}
