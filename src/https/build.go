package https

import (
	"net"
	"net/http"
	"time"
)

// 生成建立连接的信息
func (p *p2) buildClient() *http.Client {
	client := &http.Client{
		Transport: p.buildTransport(),
	}
	return client
}

// 生成传输方法
func (p *p2) buildTransport() *http.Transport {
	s := (*session)(p)

	t := &http.Transport{
		Dial: p.dial,
	}
	if s.HasRespTimeout() {
		t.ResponseHeaderTimeout = time.Millisecond * time.Duration(s.GetRespTimeout())
	}

	// 设置代理信息
	setProxy(t, (*session)(p))
	return t
}

func (p *p2) dial(netw, addr string) (net.Conn, error) {
	s := (*session)(p)

	connTimeout := s.GetConnTimeout()

	respTimeout := s.GetRespTimeout()

	var conn net.Conn
	var err error

	if 0 != connTimeout {
		//设置建立连接超时
		conn, err = net.DialTimeout(netw, addr, time.Millisecond*time.Duration(connTimeout))
	} else {
		conn, err = net.Dial(netw, addr)
	}

	if err != nil {
		return nil, err
	}

	if 0 != respTimeout {
		//设置发送接受数据超时
		_ = conn.SetDeadline(time.Now().Add(time.Millisecond * time.Duration(respTimeout)))
	}

	return conn, nil
}
