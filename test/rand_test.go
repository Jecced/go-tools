package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/randutil"
	"testing"
)

func TestRandom01(t *testing.T) {
	r := randutil.Random(100)
	for i := 0; i < 100000; i++ {
		fmt.Println(r.GetSeed(), r.NextInt(0, 100))
	}
}
