package https

import (
	"fmt"
	"testing"
)

func TestParam(t *testing.T) {
	p := make(param)
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
