/**
go 常量
 */
package main

import "fmt"

func main()  {
	const PI  = 3.1415926
	const BB string  = "这是一个常量"
	//PI=1 常量不可再次赋值
	fmt.Println(PI,BB)

	fmt.Println("常量还可以用作枚举")
	const (
		SUCCESS=200
		ERROR=500
		NOTFOUNT=404
	)
	fmt.Println(SUCCESS,ERROR,NOTFOUNT)

}

