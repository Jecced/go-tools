package https

import (
	"fmt"
	"net/url"
	"testing"
)

func TestName(t *testing.T) {
	get := Get("http://www.baidu.com?abc=def")
	fmt.Println(get.req)
	fmt.Println(get.req.param)

	parse, err := url.Parse("https://" + "http://127.0.0.12:980")
	fmt.Println(err)
	fmt.Println(parse)
}
