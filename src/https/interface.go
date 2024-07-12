package https

type ISession interface {
	AddHeader(key string, value string) *p1
	AddHeaders(entries map[string]string) *p1
	RemoveHeader(key string) *p1
	ClearHeader() *p1
	AddCookie(key string, value string) *p1
	AddCookies(entries map[string]string) *p1
	AddCookieString(cookie string) *p1
	RemoveCookie(key string) *p1
	ClearCookie() *p1
	SetTimeOut(time int) *p1
	SetConnTimeOut(time int) *p1
	SetRespTimeOut(time int) *p1
	Proxy(proxy string) *p1
	BasicAuth(user string, password string) *p1
	Retry(count uint) *p1
	SkipSSLVerify(verify bool) *p1
	Get(uri string) *p2
	Post(uri string) *p2
	GetCookies() param
	GetCookie(key string) string
	SetCookieSerializationPath(path string) *p1
}
