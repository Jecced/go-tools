package https

import (
	"bytes"
	"net/http"
)

// 设置请求头
func setHeader(request *http.Request, s *session) {
	if s.auth != "" {
		s.AddHeader(HeaderAuthorization, s.auth)
	}

	if s.req.method == POST && len(s.req.param) > 0 {
		request.Header.Set(HeaderContentType, "application/x-www-form-urlencoded")
	}

	// 设置cookie
	request.Header.Set(HeaderCookie, cookieFormat(s))

	// 混合请求头
	header := s.header.Clone()
	header.Mix(s.req.header)

	for k, v := range header {
		request.Header.Set(k, v)
	}
}

// 格式化所有 cookie 信息
func cookieFormat(s *session) string {
	// 混合 cookie
	cookie := s.cookie.Clone()
	cookie.Mix(s.req.cookie)

	if len(cookie) == 0 {
		return ""
	}
	bb := bytes.Buffer{}
	for k, v := range cookie {
		bb.WriteString(" ;")
		bb.WriteString(k)
		bb.WriteString("=")
		bb.WriteString(v)
	}

	return bb.String()[2:]
}
