package https

import (
	"bytes"
	"net/http"
)

// 设置请求头
func setHeader(request *http.Request, s *session) {
	auth := getAuth(s)
	if auth != "" {
		s.req.header.Add(headerAuthorization, auth)
	}

	if s.method == post && len(s.param) > 0 {
		request.Header.Set(headerContentType, "application/x-www-form-urlencoded")
	}

	// 设置cookie
	request.Header.Set(headerCookie, cookieFormat(s))

	// 混合请求头
	s.req.header.Mix(s.comm.header)

	for k, v := range s.req.header {
		request.Header.Set(k, v)
	}
}

// 格式化所有 cookie 信息
func cookieFormat(s *session) string {
	// 混合 cookie
	s.req.cookie.Mix(s.comm.cookie)

	if len(s.req.cookie) == 0 {
		return ""
	}
	bb := bytes.Buffer{}
	for k, v := range s.req.cookie {
		bb.WriteString(" ;")
		bb.WriteString(k)
		bb.WriteString("=")
		bb.WriteString(v)
	}

	return bb.String()[2:]
}
