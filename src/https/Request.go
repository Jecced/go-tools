package https

import (
	"net/http"
	"net/url"
)

// 独立请求模块
// 属于 session 部分
// 每个请求都是一个 request
type request struct {
	base
	// 每个request都有自己的 uri 地址, param 参数
	uri   string
	param url.Values

	// 用来确定是 get 请求还是 post 请求, 默认为 get 请求
	method requestType

	// 用于优化
	// 用来判断当前请求是否已经close
	// 当处于session阶段, 每次进行新的get和set阶段
	// 将判断上一次请求是否已经关闭, 如果没有关闭, 则自动关闭
	close bool

	response *http.Response
}
