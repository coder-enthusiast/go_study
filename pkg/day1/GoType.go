package main

import (
	"fmt"
	"reflect"
)
func main() {
	fmt.Println("这是一个控制台输出函数 在fmt包下")

	fmt.Println("变量声明")
	var aa int =1
	var bb float64=1.1
	var cc bool=true
	var dd string="golang"
	fmt.Println(aa,bb,cc,dd)

	fmt.Println("多变量声明")
	var aaa, bbb, ccc int=1,2,3
	fmt.Println(aaa,bbb,ccc)



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
