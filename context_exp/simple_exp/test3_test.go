// @Author: Perry
// @Date  : 2020/1/3
// @Desc  : 截止时间取消 context.WithDeadline

package simple_exp

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTest3(t *testing.T) {
	d := time.Now().Add(8 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("oversleep")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}

}
