package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/fileutil"
	"testing"
)

func TestName(t *testing.T) {
	dir := "/Users/ankang/saisheng/slgrpg"
	types := fileutil.FindAllFileTypes(dir)
	fmt.Println(types)
}
