package main

import "fmt"

func main()  {
	str:="abc朱志强"

	for _,s:=range str  {
		if(len(string(s))>=3){
			fmt.Printf(string(s))
		}
	}

}
