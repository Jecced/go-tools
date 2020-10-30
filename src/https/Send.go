package https

import (
	"github.com/Jecced/go-tools/src/fileutil"
	"io/ioutil"
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
	fileutil.MkdirParent(path)

	// 创建文件
	create, err := os.Create(path)
	if err != nil {
		return err
	}
	defer create.Close()

	bytes, err := ioutil.ReadAll(p.response.Body)
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
	if p.err != nil {
		return "", p.err
	}
	bytes, err := ioutil.ReadAll(p.response.Body)
	if err != nil {
		return "", err
	}
	resp := string(bytes)
	err = p.Close()
	if err != nil {
		return "", err
	}
	return resp, nil
}
