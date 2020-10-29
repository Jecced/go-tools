package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/gts"
	"testing"
)

func TestTranslate(t *testing.T) {
	cn := gts.TranslateEn2Cn("fuck your computer")
	fmt.Println(cn)
}
