package https

import (
	"fmt"
	"github.com/Jecced/go-tools/src/fileutil"
	"io/ioutil"
	"net/http"
	"os"
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

	setHeader(request, s)

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

func (s *session) Close() error {
	if s.req == nil {
		return nil
	}
	if s.req.response == nil {
		return nil
	}
	err := s.req.response.Body.Close()
	s.req.close = true
	return err
}

func (s *session) WriteFile(path string) error {
	// 标准化路径
	path = fileutil.PathFormat(path)
	// 尝试创建父文件夹, 防止因为文件夹不存在而出错
	fileutil.MkdirParent(path)

	// 创建文件
	create, err := os.Create(path)
	if err != nil {
		return err
	}
	defer create.Close()

	bytes, err := ioutil.ReadAll(s.req.response.Body)
	if err != nil {
		return err
	}

	_, err = create.Write(bytes)
	if err != nil {
		return err
	}

	return s.Close()
}

func (s *session) ReadText() (string, error) {
	bytes, err := ioutil.ReadAll(s.req.response.Body)
	if err != nil {
		return "", err
	}
	resp := string(bytes)
	err = s.Close()
	if err != nil {
		return "", err
	}
	return resp, nil
}
