package main

import (
	"fmt"
	"reflect"
)

/**
反射
 */

func main()  {
	var obj s
	obj.Name="this is name"

	getValueByType(obj)
	getValueByType(int64(1))
	getValueByType(bool(true))



	str := "123"
	setValue(&str)
	fmt.Println(str)

	//setValue(&obj)
}

/**
反射获取值类型 并且打印出值
 */
func getValueByType(o interface{})  {
	vk := reflect.ValueOf(o)
    k:=vk.Kind()
	switch k {
	case reflect.Int64:
		fmt.Println(int64(vk.Int()))
	case reflect.Bool:
		fmt.Println(vk.Bool())
	case reflect.Struct:
		fmt.Println(vk)


	}
}

/**
通过反射设置值
 */
func setValue(o interface{})  {
	a:= reflect.ValueOf(o)
	//反射修改值 必须是指针类型
	if a.Kind() == reflect.Ptr{
		a=a.Elem()
		switch a.Kind() {
		case reflect.String:
			a.SetString("反射设置值 必须使用Elem 获取指针")

		}
	}



}

type s struct {
	Name string
}


