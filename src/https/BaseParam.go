package https

type BaseParam map[string]string

func (p BaseParam) Add(key, value string) {
	p[key] = value
}

func (p BaseParam) Adds(param map[string]string) {
	for k, v := range param {
		p[k] = v
	}
}

func (p BaseParam) Remove(key string) {
	delete(p, key)
}

func (p *BaseParam) Clear() {
	*p = make(BaseParam)
}

func (p *BaseParam) Mix(other BaseParam) {
	p.Adds(other)
}
