/**
 * @Author: daipengyuan
 * @Description:
 * @File:  interface_exp.go
 * @Version: 1.0.0
 * @Date: 2021/6/15 09:47
 */

package main

import "fmt"

func printlnInterface(intf interface{}) {
	fmt.Printf("data type is = %T\n", intf)
	fmt.Printf("data content is = %v\n\n", intf)
}

func main() {
	var a = "hello"
	var b = 250
	var c = 2.22
	var d = struct {
		d1 string
	}{d1: "world"}

	printlnInterface(a)
	printlnInterface(b)
	printlnInterface(c)
	printlnInterface(d)
}
