package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	str:="abc朱志强"
	//字节数
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	//字符数

	fmt.Println()
	for _,s:=range str  {
		if(len(string(s))>=3){
			fmt.Printf(string(s))
		}
	}

	fmt.Println(str)

}
