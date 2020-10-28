package https

// 公共接口
type CommFace interface {
	// 增加一个请求头
	AddHeader(key, value string) CommFace
	// 增加多个请求头
	AddHeaders(BaseParam map[string]string) CommFace
	// 移除一个请求头
	RemoveHeader(key string) CommFace
	// 清空所有请求头
	ClearHeader() CommFace

	// 增加一个cookie
	AddCookie(key, value string) CommFace
	// 增加多个cookie
	AddCookies(BaseParam map[string]string) CommFace
	// 移除一个cookie
	RemoveCookie(key string) CommFace
	// 清空所有cookie
	ClearCookie() CommFace

	// 设置超时时间
	SetTimeOut(time int) CommFace
	// 设置连接超时时间
	SetConnTimeOut(time int) CommFace
	// 设置反馈超时时间
	SetRespTimeOut(time int) CommFace

	// 代理
	Proxy(proxy string) CommFace

	// 基础认证
	BasicAuth(user, password string) CommFace
}

// 请求前的构造接口
type Requests interface {
	// 增加一个参数
	AddParam(key, value string) Requests
	// 增加多个参数
	AddParams(param map[string]string) Requests
	// 移除一个参数
	RemoveParam(key string) Requests
	// 清空所有参数
	ClearParam() Requests

	// 发送
	//Send() Responses
}

// 请求后的构造接口
type Responses interface {
	// 关闭流
	Close()
	// 将流写入到文件
	WriteFile(path string)
	// 读取文件
	ReadText() string
}
