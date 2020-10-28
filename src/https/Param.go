package https

// 增加一个参数
func (s *session) AddParam(key, value string) requests {
	get := s.req.param.Get(key)
	if "" == get {
		s.req.param.Add(key, value)
	} else {
		s.req.param.Set(key, value)
	}
	return s
}

// 增加多个请求参数
func (s *session) AddParams(param map[string]string) requests {
	for k, v := range param {
		s.AddParam(k, v)
	}
	return s
}

// 移除一个请求参数
func (s *session) RemoveParam(key string) requests {
	s.req.param.Del(key)
	return s
}

// 清空所有请求参数
func (s *session) ClearParam() requests {
	for k := range s.req.param {
		s.RemoveParam(k)
	}
	return s
}
