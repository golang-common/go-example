/**
 * @Author: DPY
 * @Description:
 * @File:  dir_test
 * @Version: 1.0.0
 * @Date: 2021/11/6 16:27
 */

package os_exp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestGetwd(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pwd)
}

func TestReadDir(t *testing.T) {
	entries, err := os.ReadDir("/Users/lyonsdpy/Coding/go/src")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		t.Logf("%+v", entry)
	}
}

func TestChdir(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("修改前的当前目录=", pwd)
	err = os.Chdir("/Users/lyonsdpy")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("修改当前目录成功")
	pwd, err = os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("修改后的当前目录=", pwd)
}

func TestOpen(t *testing.T) {
	fd, err := os.Open("./hw.txt")
	if err != nil {
		t.Fatal(err)
	}
	fileBytes, err := ioutil.ReadAll(fd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(fileBytes))
}

func TestCreate(t *testing.T) {
	fd, err := os.Create("./hw.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fd.Name())
}

func TestRenameFile(t *testing.T) {
	// 创建文件dpy1并写入数据
	fd1, err := os.Create("./dpy1.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = fd1.Write([]byte("11111111"))
	if err != nil {
		t.Fatal(err)
	}
	// 创建文件dpy2并写入数据
	fd2, err := os.Create("./dpy2.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = fd2.Write([]byte("22222222"))
	if err != nil {
		t.Fatal(err)
	}
	// 将文件dpy1重命名为dpy2，此时dpy2将会被替换
	err = os.Rename("./dpy1.txt", "./dpy2.txt")
	if err != nil {
		t.Fatal(err)
	}
	// 查看dpy2文件
	fd2, err = os.Open("./dpy2.txt")
	if err != nil {
		t.Fatal(err)
	}
	fdBytes2, err := ioutil.ReadAll(fd2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(fdBytes2))
	// 查看dpy1文件，发现文件不存在(已被重命名)
	fd2, err = os.Open("./dpy1.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenameDir(t *testing.T) {
	err := os.Rename("/Users/lyonsdpy/tmp/dpy1", "/Users/lyonsdpy/tmp/dpy2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestChmod(t *testing.T) {
	// 创建文件并写入内容
	err := os.WriteFile("./dpy1.txt", []byte("helloworld"), 0666)
	if err != nil {
		t.Fatal(err)
	}
	// 修改文件权限为只读
	err = os.Chmod("./dpy1.txt", 0444)
	if err != nil {
		t.Fatal(err)
	}
	// 尝试再次写入这个文件，会报错permission denied
	err = os.WriteFile("./dpy1.txt", []byte("helloworld"), 0666)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemove(t *testing.T) {
	err := os.RemoveAll("./dpy1.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveAll(t *testing.T) {
	err := os.RemoveAll("./")
	if err != nil {
		t.Fatal(err)
	}
}

func TestChown(t *testing.T) {
	err := os.Chown("./hw.txt", os.Getuid(), os.Getgid())
	if err != nil {
		t.Fatal(err)
	}
}

func TestLink(t *testing.T) {
	err := os.Link("./hw.txt", "./hw2.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSymlink(t *testing.T) {
	err := os.Symlink("./hw.txt", "./hw3.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestChtimes(t *testing.T) {
	const layout = `2006-01-02 15:04:05`
	tm, err := time.Parse(layout, "2017-07-17 00:00:00")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chtimes("./hw.txt", tm, tm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadLink(t *testing.T) {
	l, err := os.Readlink("./hw3.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}

func TestReadFile(t *testing.T) {
	fBytes, err := os.ReadFile("./hw.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(fBytes))
}

func TestNewFile(t *testing.T) {
	fd, err := syscall.Open("./hw.txt", 0, 0777)
	if err != nil {
		t.Fatal(err)
	}
	file := os.NewFile(uintptr(fd), "hw.txt")
	fBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(fBytes))
}

func TestNewFile2(t *testing.T) {
	file := os.NewFile(1, "")
	_, err := file.WriteString("hello\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = file.WriteString("world\n")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewFile3(t *testing.T) {
	sokFds, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		t.Fatal(err)
	}
	f1 := os.NewFile(uintptr(sokFds[0]), "")
	defer f1.Close()
	f2 := os.NewFile(uintptr(sokFds[1]), "")
	defer f2.Close()
	c1, err := net.FileConn(f1)
	if err != nil {
		t.Fatal(err)
	}
	c2, err := net.FileConn(f2)
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_, err = c1.Write([]byte("hello world\n"))
		if err != nil {
			t.Fatal(err)
		}
		defer wg.Done()
		defer c1.Close()
	}()

	wg.Add(1)
	go func() {
		var b = make([]byte, 11)
		_, err = c2.Read(b)
		if err != nil {
			t.Fatal(err)
		}
		defer wg.Done()
		fmt.Println(string(b))
	}()

	wg.Wait()
}

func TestTruncate(t *testing.T) {
	err := os.Truncate("./hw.txt", 15)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStat(t *testing.T) {
	finfo, err := os.Stat("./hw.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(finfo.Name())
	t.Log(finfo.Size())
	t.Log(finfo.Mode())
	t.Log(finfo.IsDir())
	t.Log(indentJson(finfo.Sys()))
	t.Log(finfo.ModTime())
}

func indentJson(obj interface{}) string {
	ret, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err.Error()
	}
	//os.Pipe()
	return string(ret)
}

func TestCreateTemp(t *testing.T) {
	f, err := os.CreateTemp("./", "hw1.txt*")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write([]byte("hello hw1"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDirFS(t *testing.T) {
	dir := os.DirFS("/Users/lyonsdpy/Coding/go/src/golang_exp")
	file, err := dir.Open("os_exp")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(stat.IsDir())
	t.Log(stat.Name())
}

func TestTempDir(t *testing.T) {
	t.Log(os.TempDir())
}

func TestMkdirTemp(t *testing.T) {
	dirName, err := os.MkdirTemp("./", "dpy")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dirName)
}

func TestUserConfigDir(t *testing.T) {
	dirName, err := os.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dirName)
}

func TestUserCacheDir(t *testing.T) {
	dirName, err := os.UserCacheDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dirName)
}

func TestSameFile(t *testing.T) {
	f1, _ := os.Open("./hw1.txt")
	f1stt, _ := f1.Stat()
	f2, _ := os.Open("./hw.txt")
	f2stt, _ := f2.Stat()
	t.Log(os.SameFile(f1stt, f2stt))
}

func TestWriteFile(t *testing.T) {
	err := os.WriteFile("./hw.txt", []byte("hello write"), os.ModeAppend|0777)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPipe(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	_, err = w.Write([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	var rst = make([]byte, 11)
	_, err = r.Read(rst)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(rst))
}

func TestUserHomeDir(t *testing.T){
	dirName, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dirName)
}