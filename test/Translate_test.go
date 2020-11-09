package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/gts"
	"testing"
)

func TestTranslate(t *testing.T) {
	go test("fuck your computer")
	go test("fuck")
	go test("what are you doing")
	select {}
}
func test(text string) {
	cn, err := gts.TranslateEn2Cn(text)
	fmt.Println(cn, err)
}
