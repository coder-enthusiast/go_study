/**
 go中没有类的概念 使用可以使用struct来封装成类
 */
package main

import (
	"fmt"
	"reflect"
)

//声明一个MyInt类型 具有int的特性
type MyInt int

//int的别名
type alinInt = int

type Person struct {
	 name string
	 age int
}

func main()  {
	var myint MyInt
	var aint alinInt
	fmt.Println(reflect.TypeOf(myint),reflect.TypeOf(aint))

	var person Person
	person.age=18
	person.name="zzq"
	fmt.Println(person)
}