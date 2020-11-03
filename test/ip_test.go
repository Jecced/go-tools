package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/commutil"
	"testing"
)

func TestIp(t *testing.T) {
	b, a := commutil.GetInternal()
	fmt.Println(b, a)

	external, a := commutil.GetExternal()
	fmt.Println(external, a)
}
