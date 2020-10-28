package https

import (
	"net"
	"net/http"
	"time"
)

// 生成建立连接的信息
func buildClient(s *session) *http.Client {
	client := &http.Client{
		Transport: buildTransport(s),
	}
	return client
}

// 生成传输方法
func buildTransport(s *session) *http.Transport {
	t := &http.Transport{
		Dial:                  s.dial,
		ResponseHeaderTimeout: time.Millisecond * time.Duration(s.respTimeout),
	}

	// 设置代理信息
	setProxy(t, s)
	return t
}

func (s *session) dial(netw, addr string) (net.Conn, error) {
	connTimeout := s.req.connTimeout
	if 0 == connTimeout {
		connTimeout = s.connTimeout
	}

	respTimeout := s.req.respTimeout
	if 0 == respTimeout {
		respTimeout = s.respTimeout
	}

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
