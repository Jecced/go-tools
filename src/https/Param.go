package https

// 通用map类型参数接口
type param map[string]string

func (p param) Add(key, value string) {
	p[key] = value
}

func (p param) Adds(entries map[string]string) {
	for k, v := range entries {
		p[k] = v
	}
}

func (p param) Remove(key string) {
	delete(p, key)
}

func (p *param) Clear() {
	*p = make(param)
}

func (p *param) Mix(other param) {
	p.Adds(other)
}

func (p *param) Clone() param {
	out := make(param)
	for k, v := range *p {
		out[k] = v
	}
	return out
}
