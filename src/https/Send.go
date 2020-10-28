package https

import (
	"fmt"
	"net/http"
)

func (s *session) Send() *session {
	var request *http.Request
	var err error

	if s.req == nil {
		fmt.Println("没有请求参数req, 请执行Get或者Post")
		return s
	}

	switch s.req.method {
	case GET:
		request, err = s.getNewRequest()
	case POST:
		request, err = s.postNewRequest()
	}

	if err != nil {
		fmt.Println("生成请求对象错误", err.Error())
	}

	client := buildClient(s)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("发送请求失败", err.Error())
		return s
	}
	if response == nil {
		return s
	}

	// 处理response返回的cookie
	cookies := response.Cookies()
	for _, cookie := range cookies {
		s.AddCookie(cookie.Name, cookie.Value)
	}

	s.req.response = response

	return s
}
