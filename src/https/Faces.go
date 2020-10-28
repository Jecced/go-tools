package https

// 公共接口
type comm interface {
	// 增加一个请求头
	AddHeader(key, value string) comm
	// 增加多个请求头
	AddHeaders(BaseParam map[string]string) comm
	// 移除一个请求头
	RemoveHeader(key string) comm
	// 清空所有请求头
	ClearHeader() comm

	// 增加一个cookie
	AddCookie(key, value string) comm
	// 增加多个cookie
	AddCookies(BaseParam map[string]string) comm
	// 移除一个cookie
	RemoveCookie(key string) comm
	// 清空所有cookie
	ClearCookie() comm

	// 设置超时时间
	SetTimeOut(time int) comm
	// 设置连接超时时间
	SetConnTimeOut(time int) comm
	// 设置反馈超时时间
	SetRespTimeOut(time int) comm

	// 代理
	Proxy(proxy string) comm

	// 基础认证
	BasicAuth(user, password string) comm
}

// session 阶段的构造接口
type sessions interface {
	comm
	Get(url string) requests
	Post(url string) requests
}

// request 请求阶段的构造函数
type requests interface {
	comm
	// 增加一个参数
	AddParam(key, value string) requests
	// 增加多个参数
	AddParams(param map[string]string) requests
	// 移除一个参数
	RemoveParam(key string) requests
	// 清空所有参数
	ClearParam() requests

	// 发送
	Send() responses
}

// Send() 请求后的构造接口
// 如果没有进行写入或者读取, 则必须手动关闭Close, 否则会造成泄漏
type responses interface {
	// 关闭流
	Close() error
	// 将流写入到文件
	WriteFile(path string) error
	// 读取文件
	ReadText() (string, error)
}
