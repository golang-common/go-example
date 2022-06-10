/**
 * @Author: DPY
 * @Description:
 * @File:  signal_test
 * @Version: 1.0.0
 * @Date: 2021/11/10 00:17
 */

package os_exp

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"
)

/*
exec.LookPath()
exec.CommandContext()
exec.ErrNotFound
exec.Cmd{}.
exec.ExitError{}
exec.Error{}
*/

func TestCommand(t *testing.T) {
	cmd := exec.Command("ls", "-l", `/Users/lyonsdpy/Coding/go`)
	rst, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(rst))
}

func TestLookPath(t *testing.T) {
	path, err := exec.LookPath("ls")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
}

func TestExecExp1(t *testing.T) {
	cmd := exec.Command("ls", "-l")
	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

func TestExecExp2(t *testing.T) {
	cmd := exec.Command("ls", "-l")
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b.Bytes()))
}

func TestExecExp3(t *testing.T) {
	var (
		bufRead bytes.Buffer
		cmd     = new(exec.Cmd)
	)
	cmd.Stdout = &bufRead
	cmd.Path = `/bin/ls`
	cmd.Args = []string{"-l", "/tmp"}
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bufRead.Bytes()))
}

func TestExecExp4(t *testing.T) {
	var (
		readBuf  bytes.Buffer
		writeBuf bytes.Buffer
	)
	cmd := &exec.Cmd{
		Path:   "/bin/bash",
		Stdin:  &writeBuf,
		Stdout: &readBuf,
		Stderr: &readBuf,
	}
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	n, err := writeBuf.WriteString("echo hello\n")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("成功写入了%d字节的数据", n)
	n, err = writeBuf.WriteString("echo world\n")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("成功写入了%d字节的数据", n)
	err = cmd.Wait()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(readBuf.String())
}

func TestExecExp5(t *testing.T) {
	var wg sync.WaitGroup
	cmd := &exec.Cmd{
		Path: "/bin/bash",
	}
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		t.Fatal(err)
	}
	inPipe, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	wg.Add(1)
	// 开启一个协程，分三次写入数据，最后一次人为制造一个stderr
	go func() {
		defer wg.Done()
		var n int
		var e error
		// 第一次写入
		n, e = inPipe.Write([]byte("echo hello\n"))
		if e != nil {
			fmt.Printf("写入出错，err=%s", err)
		}
		fmt.Printf("成功写入%d字节数据\n", n)
		// 第二次写入
		time.Sleep(1 * time.Second)
		n, e = inPipe.Write([]byte("echo world\n"))
		if err != nil {
			fmt.Printf("写入出错，err=%s", err)
		}
		fmt.Printf("成功写入%d字节数据\n", n)
		// 第三次写入
		time.Sleep(1 * time.Second)
		n, e = inPipe.Write([]byte("ddd\n"))
		if err != nil {
			fmt.Printf("写入出错，err=%s", err)
		}
		fmt.Printf("成功写入%d字节数据\n", n)
		fmt.Println("写入完成")
	}()
	// 开启协程持续读取信息，直到读取到非EOF错误再退出
	wg.Add(1)
	go func() {
		var n int
		var e error
		for i := 1; ; {
			readBytes := make([]byte, 1024)
			n, e = outPipe.Read(readBytes)
			if e != nil && e != io.EOF {
				fmt.Printf("第%d次读取标准输出失败，退出读取，err=%s\n", i, e)
				break
			}
			if n > 0 {
				fmt.Printf("第%d次成功读取标准输出，大小%d字节，内容=%s\n", i, n, string(readBytes[:n]))
				i++
			}
			if i >= 3 {
				wg.Done()
			}
		}
	}()
	// 开启协程持续读取信息，一旦读取到实际内容则打印内容并退出
	wg.Add(1)
	go func() {
		defer wg.Done()
		var n int
		var e error
		for i := 1; ; i++ {
			readBytes := make([]byte, 1024)
			n, e = errPipe.Read(readBytes)
			if e != nil && e != io.EOF {
				fmt.Printf("第%d次读取标准错误失败，退出读取，err=%s\n", i, e)
				break
			}
			if n > 0 {
				fmt.Printf("第%d次成功读取标准错误，退出读取，大小%d字节，内容=%s\n", i, n, string(readBytes[:n]))
				break
			}
		}
	}()
	wg.Wait()
	// 向系统发送kill信号结束进程
	fmt.Printf("向进程发送 %s 系统信号，进程结束\n", os.Kill)
	err = cmd.Process.Signal(os.Kill)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("退出")
}
