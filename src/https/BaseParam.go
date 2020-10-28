package https

type Param map[string]string

func (p Param) Add(key, value string) {
	p[key] = value
}

func (p Param) Adds(param map[string]string) {
	for k, v := range param {
		p[k] = v
	}
}

func (p Param) Remove(key string) {
	delete(p, key)
}

func (p *Param) Clear() {
	*p = make(Param)
}
