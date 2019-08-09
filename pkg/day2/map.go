/**
map key-value 无序
 */
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func test()  {
	var data=map[string]string{"name":"zzq","age":"20"}
	fmt.Println(data)
	fmt.Println(data["name"])

	//判断key是否存在
	value,ok:=data["age"]
	if ok{
		fmt.Println(value)
	}else {
		fmt.Println("不存在")
	}

	//使用make
	data1:=make(map[string]string,10)
	fmt.Println(len(data1))
	data1["user"]="zzq"
	data1["age"]="19"
	fmt.Println(data1)

	//遍历map
	for k,v := range data1 {
		fmt.Println(k,v)
	}
	//只遍历k
	for k := range data1 {
		fmt.Println(k)
	}
	//删除健值对
	delete(data1,"name")
	for k,v := range data1 {
		fmt.Println(k,v)
	}
}

func main() {

	str := "Studies serve for delight for ornament and for ability Their chief use for delight"
	wordCount(str)
}

//求出单词出现的次数
func wordCount(str string)  {
	strs := strings.Split(str," ")
	fmt.Println(reflect.TypeOf(strs))
	data := make(map[string]int)
	for _,value := range strs{
		v,ok := data[value]
		if ok{
			data[value]=v+1
		}else{
			data[value]=1
		}
	}
	for k,v := range data {
		fmt.Println(k,v)
	}
}
