package https

type BaseParam map[string]string

func (p BaseParam) Add(key, value string) {
	p[key] = value
}

func (p BaseParam) Adds(entries map[string]string) {
	for k, v := range entries {
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

func (p *BaseParam) Clone() BaseParam {
	out := make(BaseParam)
	for k, v := range *p {
		out[k] = v
	}
	return out
}
