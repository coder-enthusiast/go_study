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

//Go中 开头字母大写是公开访问的 小写是私有的
type Person struct {
	 Name string
	 Age int
}

func (s *Person) setName(name string)  {
	s.Name=name
}

func (s *Person) setAge(age int)  {
	s.Age=age
}

func main() {
	var myint MyInt
	var aint alinInt
	fmt.Println(reflect.TypeOf(myint), reflect.TypeOf(aint))

	var person Person
	person.setName("zzq")
	person.setAge(18)
	fmt.Println(person)
}