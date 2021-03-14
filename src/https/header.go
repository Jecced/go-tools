package https

import (
	"bytes"
	"net/http"
)

// 设置请求头
func (s *session) SetHeader(request *http.Request) {
	auth := s.GetAuth()
	if auth != "" {
		s.req.header.Add(headerAuthorization, auth)
	}

	if s.method == post {
		if s.usePayload {
			request.Header.Set(headerContentType, "application/json")
		} else if len(s.param) > 0 {
			request.Header.Set(headerContentType, "application/x-www-form-urlencoded")
		}
	}

	// 设置cookie
	request.Header.Set(headerCookie, s.CookieFormat())

	// 混合请求头
	s.req.header.Mix(s.comm.header)

	for k, v := range s.req.header {
		request.Header.Set(k, v)
	}
}

// 格式化所有 cookie 信息
func (s *session) CookieFormat() string {
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
