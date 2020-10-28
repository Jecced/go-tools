package https

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	get := Get("http://www.baidu.com?abc=def")
	fmt.Println(get.req)
	fmt.Println(get.req.param)
}
