package https

import (
	"fmt"
	"log"
	"net/http"
)

// 阶段二
// 维护单个请求私有区
// // 发送
// Send() responses
type p2 session

// 增加一个参数
func (p *p2) AddParam(key, value string) *p2 {
	get := p.param.Get(key)
	if "" == get {
		p.param.Add(key, value)
	} else {
		p.param.Set(key, value)
	}
	return p
}

func (p *p2) ChangeMethod(method string) *p2 {
	p.method = method
	return p
}

// 增加多个请求参数
func (p *p2) AddParams(param map[string]string) *p2 {
	for k, v := range param {
		p.AddParam(k, v)
	}
	return p
}

// 移除一个请求参数
func (p *p2) RemoveParam(key string) *p2 {
	p.param.Del(key)
	return p
}

// 清空所有请求参数
func (p *p2) ClearParam() *p2 {
	for k := range p.param {
		p.RemoveParam(k)
	}
	return p
}

// 使用payload post请求
func (p *p2) UsePayload(body interface{}) *p2 {
	p.usePayload = body
	return p
}

func (p *p2) Send() *p3 {
	s := (*session)(p)
	retry := s.GetRetry()
	if 0 == retry {
		p.err = p.send()
		return (*p3)(p)
	}
	var err error
	for i := 0; i < retry; i++ {
		err = p.send()
		if err == nil {
			break
		}
		log.Println(fmt.Sprintf("\nretry: %d, uri: %s\ncause: %v", i+1, p.uri, err))
	}
	if err != nil {
		p.err = err
	}
	return (*p3)(p)
}

func (p *p2) send() error {
	var request *http.Request
	var err error

	switch p.method {
	case get:
		request, err = p.getNewRequest()
	case post:
		request, err = p.postNewRequest()
	}

	if err != nil {
		return err
	}

	client := p.buildClient()

	s := (*session)(p)

	s.SetHeader(request)

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	// 处理response返回的cookie
	cookies := response.Cookies()
	for _, cookie := range cookies {
		p.comm.cookie.Add(cookie.Name, cookie.Value)
	}

	p.response = response
	return nil
}
