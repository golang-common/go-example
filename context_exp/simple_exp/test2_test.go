// @Author: Perry
// @Date  : 2020/1/3
// @Desc  : 超时取消context.WithTimeout

package simple_exp

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func TestTest2(t *testing.T) {
	//定义一个5秒超时的context(5秒后ctx.done管道可取值)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("Hey,I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()
}
