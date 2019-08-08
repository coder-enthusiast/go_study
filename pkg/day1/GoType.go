package main

import (
	"fmt"
	"reflect"
)
func main() {

	fmt.Println("初始化默认值int float bool string")
	var a int
	var b float64
	var c bool
	var d string
	fmt.Printf("%v %v %v %q\n", a, b, c, d)

	//e:=1 这种方式声明变量必须是一个新变量
	fmt.Println("类型自动推测int string float bool")
	e:=1
	f:="1"
	g:=1.1
	k:=true
	fmt.Println(reflect.TypeOf(e),reflect.TypeOf(f),reflect.TypeOf(g),reflect.TypeOf(k))



}
