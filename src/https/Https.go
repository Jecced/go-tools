package https

type RequestType string

const (
	GET  RequestType = "GET"
	POST RequestType = "POST"
)

func Get(url string) *session {
	return Session().Get(url)
}

func Post(url string) *session {
	return Session().Post(url)
}

func Session() *session {
	return &session{}
}
