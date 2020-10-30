package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/gts"
	"testing"
)

func TestTranslate(t *testing.T) {
	test("fuck your computer")
	test("fuck")
	test("what are you doing")
}
func test(text string) {
	cn, err := gts.TranslateEn2Cn(text)
	fmt.Println(cn)
	fmt.Println(err)
}
