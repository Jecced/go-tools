package test

import (
	"context"
	"fmt"
	"testing"
)

func TestContext01(t *testing.T) {
	ctx := context.Background()
	fmt.Println(ctx)

	cancel, cancelFunc := context.WithCancel(ctx)
	fmt.Println(cancel)
	fmt.Println(cancelFunc)

}
