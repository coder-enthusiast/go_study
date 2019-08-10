package main

import (
	"fmt"
	"reflect"
)

/**
go的接口不仅具有接口的功能 还具有Java Object的功能 空接口可以接收任何值
 */

type Sayer interface {
	Say() string
	HiSay()
}

type zzq struct {
	Name string
}


func (z zzq) HiSay() {
	fmt.Printf("Hi Say")
}

func (z zzq) Say() string {
	return fmt.Sprintf("%s say\n",z.Name)
}

func main()  {
	var say  Sayer
	z := zzq{Name:"zzq"}
	say=z
	say.Say()
	say.HiSay()
	getType(1)
	getType(2.2)
	getType(float64(2.2))
	getType(z)
	getType(true)
}

func getType(obj interface{})  {
	fmt.Println("------",reflect.TypeOf(obj))
	switch obj.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case zzq:
		fmt.Println("zzq")
	default:
		fmt.Println("未知类型")

	}
}
