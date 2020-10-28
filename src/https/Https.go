package https

type requestType string

const (
	get  requestType = "get"
	post requestType = "post"
)

// 枚举, 固定请求头
const (
	headerContentType   = "Content-Type"
	headerAuthorization = "Authorization"
	headerCookie        = "Cookie"
)

func Get(url string) requests {
	return Session().Get(url)
}

func Post(url string) requests {
	return Session().Post(url)
}

func Session() sessions {
	s := &session{}
	// 默认请求超时时间
	s.SetTimeOut(30_000)
	s.header = make(baseParam)
	s.cookie = make(baseParam)
	return s
}
