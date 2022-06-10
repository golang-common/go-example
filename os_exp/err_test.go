/**
 * @Author: DPY
 * @Description:
 * @File:  err_test
 * @Version: 1.0.0
 * @Date: 2021/11/9 14:35
 */

package os_exp

import (
	"errors"
	"net"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewSyscallError(t *testing.T) {
	err := os.NewSyscallError("test", errors.New("hello world"))
	t.Logf("%+v", err)
	t.Log(reflect.TypeOf(err))
}

func TestIsExist(t *testing.T) {
	_, err := os.OpenFile("./hw.txt", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		t.Log(os.IsExist(err))
	}
}

func TestIsNotExist(t *testing.T) {
	_, err := os.Open("./hw2.txt")
	if err != nil {
		t.Log(os.IsNotExist(err))
	}
}

func TestIsPermission(t *testing.T) {
	_, err := os.OpenFile("./hw5.txt", os.O_CREATE, 0444)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.OpenFile("./hw5.txt", os.O_RDWR, 0444)
	if err != nil {
		t.Log(os.IsPermission(err))
	}
}

func TestIsTimeout(t *testing.T) {
	_, err := net.DialTimeout("tcp", "www.baidu.com:19999", 2*time.Second)
	if err != nil {
		t.Log(err)
		t.Log(os.IsTimeout(err))
	}
}
