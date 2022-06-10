/**
 * @Author: DPY
 * @Description:
 * @File:  reflect_test
 * @Version: 1.0.0
 * @Date: 2022/1/1 11:35
 */

package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Tif interface {
	Tint() string
}

type TFunc struct {
	Name string `json:"name"`
}

func (t TFunc) Tint() string {
	return t.Name
}

func Tvar(i Tif) {
	fmt.Println(reflect.TypeOf(i))
}

func Test1(t *testing.T) {
	a := TFunc{Name: "dpy"}
	Tvar(a)
}

func testAppend(data interface{}) {
	v := reflect.ValueOf(data).Elem()
	v.Set(reflect.Append(v, reflect.ValueOf("a")))
	//fmt.Printf("%p\n", data)
	//fmt.Printf("%p\n", v.Interface())
}

func TestAppend(t *testing.T) {
	var a = []string{"1", "2"}
	testAppend(&a)
	fmt.Println(a)
}

func TestBAdd(t *testing.T){
	a := rune(97)
	fmt.Println(string(a))
	a++
	fmt.Println(string(a))
}