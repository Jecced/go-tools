package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/https"
	"net/url"
	"testing"
)

func TestParam(t *testing.T) {
	p := make(https.Param)
	p.Add("ok", "realy?")
	p.Adds(map[string]string{
		"真的":    "xiongmao",
		"panda": "nonono",
	})
	fmt.Println(p)
	p.Remove("ok")
	fmt.Println(p)

	p.Clear()
	fmt.Println(p)
}

func TestURL(t *testing.T) {
	u := "https://www.baidu.com/bai?a=b&c=d"
	query, err := url.ParseQuery(u)
	fmt.Println(query)
	fmt.Println(err)
}
