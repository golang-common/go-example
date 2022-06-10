/**
 * @Author: DPY
 * @Description:
 * @File:  sys_test
 * @Version: 1.0.0
 * @Date: 2021/11/9 09:59
 */

package os_exp

import (
	"fmt"
	"os"
	"testing"
)

func TestHostName(t *testing.T) {
	name, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}

func TestExpand(t *testing.T) {
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "dpy",
			"hello": "hi",
		}
		if m[key] != "" {
			return m[key]
		}
		return key
	}
	//  hello,world，由于hello world之前没有$符号，则无法利用map规则进行转换
	s := "hello,world"

	//  hi,dpy finish，finish没有在map规则中，所以还是返回原来的值
	s1 := "$hello,$world $finish"
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))
}

func TestExpandEnv(t *testing.T) {
	var str = "my home is $HOME"
	es := os.ExpandEnv(str)
	t.Log(es)
}

func TestClearenv(t *testing.T) {
	err := os.Setenv("hello", "11")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv("world", "22")
	if err != nil {
		t.Fatal(err)
	}
	e1 := os.Getenv("hello")
	e2 := os.Getenv("world")
	t.Log(e1)
	t.Log(e2)
	os.Clearenv()
	e1 = os.Getenv("hello")
	e2 = os.Getenv("world")
	t.Log(e1)
	t.Log(e2)
}

func TestEnviron(t *testing.T) {
	envList := os.Environ()
	for _, env := range envList {
		t.Log(env)
	}
}

func TestExecutable(t *testing.T) {
	ecb, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ecb)
}

func TestFindProcess(t *testing.T) {
	proc, err := os.FindProcess(38523)
	if err != nil {
		t.Fatal(err)
	}
	err = proc.Kill()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetgid(t *testing.T) {
	t.Log(os.Getgid())
}

func TestGetegid(t *testing.T) {
	t.Log(os.Getegid())
}

func TestGetgroups(t *testing.T) {
	grps, err := os.Getgroups()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(grps)
}

func TestGetuid(t *testing.T) {
	t.Log(os.Getuid())
}

func TestGeteuid(t *testing.T) {
	t.Log(os.Geteuid())
}

func TestGetpid(t *testing.T) {
	t.Log(os.Getpid())
}

func TestGetppid(t *testing.T) {
	t.Log(os.Getppid())
}

func TestGetpagesize(t *testing.T) {
	t.Log(os.Getpagesize())
}

func TestGetenv(t *testing.T) {
	t.Log(os.Getenv("HOME"))
}

func TestLookupEnv(t *testing.T) {
	env, ok := os.LookupEnv("HOME")
	if ok {
		t.Log(env)
	}
}

func TestSetenv(t *testing.T) {
	err := os.Setenv("dpy", "hello world")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(os.Getenv("dpy"))
}

func TestStartProcess(t *testing.T) {
	attr := new(os.ProcAttr)
	_, err := os.StartProcess("/Users/lyonsdpy/Coding/go/bin/dlv", []string{"dlv"}, attr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnsetenv(t *testing.T) {
	err := os.Setenv("dpy", "hello world")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(os.Getenv("dpy"))
	err =os.Unsetenv("dpy")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(os.Getenv("dpy"))
}
