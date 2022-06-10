/**
 * @Author: DPY
 * @Description:
 * @File:  net_test
 * @Version: 1.0.0
 * @Date: 2021/11/12 15:42
 */

package net_exp

import (
	"reflect"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	a := "aaa|bbb"
	t.Log(a[:3])
	t.Log(a[4:])
}

func TestLayout(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	t.Log(len(layout))
	t.Log(layout[0:4])
	t.Log(layout[5:7])
	t.Log(layout[8:10])
	t.Log(layout[11:13])
	t.Log(layout[14:16])
	t.Log(layout[17:19])
}

type a struct {
	A string
	B string
}

func TestTime(t *testing.T) {
	tm := time.Date(2021, time.Month(11), 31, 0, 0, 0, 0, time.Local)
	t.Log(tm.String())
	t.Log(reflect.ValueOf(a{}).NumField())
}
