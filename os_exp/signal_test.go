/**
 * @Author: DPY
 * @Description:
 * @File:  signal_test
 * @Version: 1.0.0
 * @Date: 2021/11/10 23:21
 */

package os_exp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestSigNotify(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := <-c
	fmt.Println("获取到信号:", s)
}

func TestSigNotifyAll(t *testing.T) {
	c := make(chan os.Signal, 1)
	fmt.Println(os.Getpid())
	signal.Notify(c)
	s := <-c
	fmt.Println("Got signal:", s)
}

func TestSigNotifyContext(t *testing.T) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Signal(os.Interrupt); err != nil {
		t.Fatal(err)
	}

	select {
	case <-time.After(time.Second):
		fmt.Println("missed signal")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		// ctx.Done后，尽快调用stop函数停止接收信号
		stop()
	}
}

// 调用Reset前发送一次信号，打印helloworld
// 调用Reset后发送一次信号，执行系统默认行为
func TestSigReset(t *testing.T) {
	// 获取进程对象
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	// 拦截信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// 向进程第一次发送Interrupt信号
	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Fatal(err)
	}
	s := <-c
	fmt.Println("拦截到信号", s)

	// 调用Reset恢复Interrupt的默认行为
	signal.Reset(os.Interrupt)

	// 向进程第二次发送Interrupt信号
	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Fatal(err)
	}
	// 执行了Interrupt默认动作，进程以错误码1退出
	// 一下代码都不会被执行
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("结束,超时")
	case s = <-c:
		fmt.Println("结束，接收到信号", s)
	}
	fmt.Println("hello")
}

func TestSigStop(t *testing.T) {
	// 获取进程对象
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	// 拦截信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// 向进程第一次发送Interrupt信号
	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Fatal(err)
	}
	s := <-c
	fmt.Println("拦截到信号", s)

	// 调用Stop停止c接收任何信号
	signal.Stop(c)

	// 向进程第二次发送Interrupt信号,但不会被接收
	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Fatal(err)
	}
	s = <-c
	fmt.Println("尝试拦截信号", s)
}

func TestSigIgnore(t *testing.T) {
	// 获取进程对象
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	// 拦截信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// 向进程第一次发送Interrupt信号
	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Fatal(err)
	}

	s := <-c
	fmt.Println("拦截到信号", s)
	// 调用Ignore忽略信号
	signal.Ignore(os.Interrupt)

	// 向进程第二次发送Interrupt信号,但不会被接收
	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSigIgnored(t *testing.T) {
	signal.Ignore(os.Interrupt)
	t.Log(signal.Ignored(os.Interrupt))
}
