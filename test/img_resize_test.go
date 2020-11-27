package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/imgutil"
	"testing"
)

func TestResize(t *testing.T) {
	in := "/Users/ankang/Desktop/600.jpeg"
	out := "/Users/ankang/Desktop/601.jpeg"
	image, err := imgutil.LoadImage(in)
	if err != nil {
		fmt.Println(err)
		return
	}
	resize := imgutil.Resize(image, 100, 100)
	err = imgutil.SaveImage(out, resize)
	fmt.Println(err)
}
