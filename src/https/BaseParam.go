package https

type baseParam map[string]string

func (p baseParam) Add(key, value string) {
	p[key] = value
}

func (p baseParam) Adds(entries map[string]string) {
	for k, v := range entries {
		p[k] = v
	}
}

func (p baseParam) Remove(key string) {
	delete(p, key)
}

func (p *baseParam) Clear() {
	*p = make(baseParam)
}

func (p *baseParam) Mix(other baseParam) {
	p.Adds(other)
}

func (p *baseParam) Clone() baseParam {
	out := make(baseParam)
	for k, v := range *p {
		out[k] = v
	}
	return out
}
