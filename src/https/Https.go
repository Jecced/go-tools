package https

type RequestType string

const (
	GET  RequestType = "GET"
	POST RequestType = "POST"
)

// 枚举, 固定请求头
const (
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
	HeaderCookie        = "Cookie"
)

func Get(url string) *session {
	return Session().Get(url)
}

func Post(url string) *session {
	return Session().Post(url)
}

func Session() *session {
	s := &session{}
	// 默认请求超时时间
	s.SetTimeOut(30_000)
	return s
}
