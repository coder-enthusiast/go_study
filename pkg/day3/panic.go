/**
异常
 */
package main

import "fmt"

func exception(a int ,b int) int  {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if b==0{
		panic("by /0")
	}

	return a/b
}

func main()  {

	exception(1,0)
	fmt.Println("异常处理")
}