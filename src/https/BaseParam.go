package https

type param map[string]string

func (p param) Add(key, value string) {
	p[key] = value
}

func (p param) Adds(param map[string]string) {
	for k, v := range param {
		p[k] = v
	}
}

func (p param) Remove(key string) {
	delete(p, key)
}

func (p *param) Clear() {
	*p = make(param)
}
