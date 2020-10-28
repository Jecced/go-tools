package https

import "net/url"

type Param url.Values

// 独立请求模块
// 属于 session 部分
// 每个请求都是一个 request
type request struct {
	comm
	// 每个request都有自己的 uri 地址, param 参数
	uri   string
	param Param

	// 用来确定是 GET 请求还是 POST 请求, 默认为 GET 请求
	method RequestType
}
