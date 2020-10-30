package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/https"
	"testing"
)

func TestHttp01(t *testing.T) {

	text, err := https.Get("https://www.google.com/").
		Proxy("localhost:1081").
		Retry(3).
		Send().
		ReadText()
	fmt.Println(text, err)
}
