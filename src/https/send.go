package https

import (
	"errors"
	"github.com/Jecced/go-tools/src/fileutil"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// 阶段三
// 发送请求后的处理
// // 关闭流
// Close() error
// // 将流写入到文件
// WriteToFile(path string) error
// // 读取文件
// ReadText() (string, error)
type p3 session

func (p *p3) Close() error {
	if p.err != nil {
		return p.err
	}
	if p.req == nil {
		return nil
	}
	if p.response == nil {
		return nil
	}
	err := p.response.Body.Close()
	p.close = true
	return err
}

func (p *p3) WriteToFile(path string) error {
	if p.err != nil {
		return p.err
	}
	// 标准化路径
	path = fileutil.PathFormat(path)
	// 尝试创建父文件夹, 防止因为文件夹不存在而出错
	err := fileutil.MkdirParent(path)
	if err != nil {
		return err
	}

	// 创建文件
	create, err := os.Create(path)
	if err != nil {
		return err
	}
	defer create.Close()

	bytes, err := p.GetBytes()
	if err != nil {
		return err
	}

	_, err = create.Write(bytes)
	if err != nil {
		return err
	}

	return p.Close()
}

func (p *p3) ReadText() (string, error) {
	bytes, err := p.GetBytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (p *p3) GetReader() (io.Reader, error) {
	if p.response == nil {
		return nil, errors.New("没有找到response信息")
	}

	if p.response.Body == nil {
		return nil, errors.New("response中没有找到body流")
	}
	w := p.response.Body
	return w, nil
}

func (p *p3) GetBytes() ([]byte, error) {
	if p.err != nil {
		return nil, p.err
	}
	body, err := p.GetReader()
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	err = p.Close()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (p *p3) GetReaderAndHeader() (io.Reader, http.Header, error) {
	header := p.response.Header
	reader, err := p.GetReader()
	return reader, header, err
}

func (p *p3) GetBytesAndHeader() ([]byte, http.Header, error) {
	header := p.response.Header
	bytes, err := p.GetBytes()
	return bytes, header, err
}

func (p *p3) ReadTextAndHeader() (string, http.Header, error) {
	header := p.response.Header
	text, err := p.ReadText()
	return text, header, err
}

func (p *p3) GetHeader() http.Header {
	return p.response.Header
}
