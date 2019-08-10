package main

import (
	"fmt"
	"reflect"
)

/**
通过反射获取结构体中的字段信息
 */

type fi struct {
	a string `json:"A"`
	b int
	c bool `json:"C"`
}

func main()  {
	var obj=&fi{
		a: "aaa",
		b: 1,
		c: true,
	}
	printField(obj)
}

func printField(o interface{}){
	ot:= reflect.TypeOf(o)
    ov:=reflect.ValueOf(o).Elem()
		ot = ot.Elem()
		for i := 0;i<ot.NumField();i++  {
			fmt.Println("field name",ot.Field(i).Name)
			fmt.Println("field type",ot.Field(i).Type)
			fmt.Println("field value",ov.Field(i))
			fmt.Println("field tag",ot.Field(i).Tag.Get("json"))
			fmt.Println()
		}


}