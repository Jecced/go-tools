package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	u := "https://www.baidu.com/bai?a=b&c=d"
	query, err := url.ParseQuery(u)
	fmt.Println(query)
	fmt.Println(err)
}
