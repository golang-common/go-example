// @Author: Perry
// @Date  : 2020/1/3
// @Desc  : 
/*
1-main函数
	-用cancel创建一个context
	-随机超时后调用取消函数

2-doWorkContext函数
	-派生一个超时context
	-这个context将被取消当
	 	-main调用取消函数
	 	-或超时到
	 	-或doWorkContext调用它的取消函数
	-启动goroutine传入派生上下文执行一些慢处理
	-等待goroutine完成或上下文被maingoroutine取消，以优先发生者为准

4-sleepRandomContext函数
	-开启一个goroutine去做些缓慢的处理
	-等待该goroutine完成或，
	-等待context被maingoroutine取消，超时或它自己的取消函数被调用

5-sleepRandom函数
	-随机时间休眠
	-此示例使用休眠来模拟随机处理时间，在实际示例中，您可以使用通道来通知此函数，
	以开始清理并在通道上等待它，以确认清理已完成。
*/

package simple_exp2

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func sleepRandom(fromFunction string, ch chan int) {
	defer func() { log.Println(fromFunction, "sleepRandom complete") }()

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleepTime := randomNumber + 100
	log.Println(fromFunction, "starting sleep for", sleepTime, "ms")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	log.Println(fromFunction, "Waking up,slept for ", sleepTime, "ms")
	if ch != nil {
		ch <- sleepTime
	}
}

func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		log.Println("sleepRandomContext complete")
		ch <- true
	}()
	sleepTimeChan := make(chan int)
	go sleepRandom("sleepRandomContext", sleepTimeChan)
	select {
	case <-ctx.Done():
		log.Println("sleepRandomContext:Time to return")
	case sleepTime := <-sleepTimeChan:
		log.Println("Slept for ", sleepTime, "ms")
	}
}

func doWorkContext(ctx context.Context) {
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	defer func() {
		log.Println("doWorkContext complete")
		cancelFunc()
	}()

	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)
	select {
	case <-ctx.Done():
		log.Println("doWorkContext: Time to return")
	case <-ch:
		log.Println("sleepRandomContext returned")
	}
}

func Test1(t *testing.T) {
	ctx := context.Background()
	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	defer func() {
		log.Println("Main Defer: canceling context")
		cancelFunc()
	}()
	go func() {
		sleepRandom("Main", nil)
		cancelFunc()
		fmt.Println("Main sleep complete. canceling context")
	}()

	doWorkContext(ctxWithCancel)
}
