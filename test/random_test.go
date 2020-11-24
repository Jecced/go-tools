package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/randutil"
	"testing"
)

func TestRandom(t *testing.T) {
	rand := randutil.Random(100)
	fmt.Println(rand.Next(100, 200))
	fmt.Println(rand.NextInt(100, 200))
	fmt.Println(rand.NextBool())
	fmt.Println(rand.NextInt32(100, 200))
	fmt.Println(rand.NextInt64(100, 200))

	fmt.Println("种子: ", rand.GetSeed())

	rand.SetSeed(100)
	fmt.Println(rand.Next(100, 200))
	fmt.Println(rand.NextInt(100, 200))
	fmt.Println(rand.NextBool())
	fmt.Println(rand.NextInt32(100, 200))
	fmt.Println(rand.NextInt64(100, 200))
}
