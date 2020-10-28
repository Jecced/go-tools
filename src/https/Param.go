package https

// 增加一个参数
func (h *https) AddParam(key, value string) Requests {
	get := h.param.Get(key)
	if "" == get {
		h.param.Add(key, value)
	} else {
		h.param.Set(key, value)
	}
	return h
}

// 增加多个请求参数
func (h *https) AddParams(param map[string]string) Requests {
	for k, v := range param {
		h.AddParam(k, v)
	}
	return h
}

// 移除一个请求参数
func (h *https) RemoveParam(key string) Requests {
	h.param.Del(key)
	return h
}

// 清空所有请求参数
func (h *https) ClearParam() Requests {
	for k := range h.param {
		h.RemoveParam(k)
	}
	return h
}
