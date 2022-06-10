/**
 * @Author: DPY
 * @Description:
 * @File:  test4_test.go
 * @Version: 1.0.0
 * @Date: 2022/2/16 13:52
 */

package simple_exp

import (
	"context"
	"testing"
	"time"
)

func ctx1(ctx context.Context) {
	txx := context.WithValue(ctx, "1", "2")
	txx.Done()
}

func ctx2(ctx context.Context) {
}

func Test4(t *testing.T) {
	ctx := context.Background()
	go ctx1(ctx)
	go ctx2(ctx)
	for {

		select {
		case <-ctx.Done():
			t.Log("done")
			return
		case <-time.After(2 * time.Second):
			t.Log("timeout")
			return
		}
	}
}
