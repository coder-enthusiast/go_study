/**
条件语句
 */
package main

import "fmt"

func main()  {
	//if else
	a:=false
	if a{
		fmt.Println(a)
	}else{
		fmt.Println(!a)
	}


	//for
	for i:=0;i<10 ;i++  {
		fmt.Print(i)
		if i%2==0{
			//可用于结束循环
			goto breaks
		}
	}
	//goto标签
	breaks:
		fmt.Println("结束循环")




	//switch 变量
	b:="bb"
	switch b {

	case "bb":
		fmt.Println("bb")
		//继续执行下一个case（不管是否符合）
		fallthrough
	case "aa":
		fmt.Println("aa")
	default:
		fmt.Println("default")
	}
	//switch 表达式 switch后不跟变量
	o := 8
	switch  {
	case o %2 == 0 && o % 4 ==0:
		fmt.Println("2")
	case o % 4 ==0:
		fmt.Println("4")
	}

	//break 具有与Java一致的功能
	/**
	break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上
	 */
BREAK:
	for i:=0;i<10 ;i++  {
		if(i>=5){
			break BREAK
		}
	}
	//continue 具有与Java一致的功能
	/**
	在 continue语句后添加标签时，表示开始标签对应的循环
	 */
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}


}
