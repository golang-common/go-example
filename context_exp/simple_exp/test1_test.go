// @Author: Perry
// @Date  : 2020/1/3
// @Desc  : 通过context.WithValue来传值

package simple_exp

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("key"), "is cancel")
			return
		default:
			fmt.Println(ctx.Value("key"), "int goroutine")
			time.Sleep(2 * time.Second)
		}
	}
}

func TestTest1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, "key", "add value")
	go watch(valueCtx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
