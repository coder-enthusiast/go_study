package main

import "fmt"

func main()  {
	/**
	下面这端代码会报错 为什么报错？指针变量需要手动分配地址 不会自动初始化
	 */
	//var a *int;*a = 100;fmt.Println(a)//panic: runtime error: invalid memory address or nil pointer dereference
    //new  申请指针内存地址
	var a  = new(int)
	*a=10
	fmt.Println(a);fmt.Println(&*a)
	modify1(*a)
	fmt.Println(*a)
	modify2(&*a)
	fmt.Println(*a)

	b := 10
	c:=&b
	//&c是c的内存位置 存的是b的内存位置 *(&c)取c值
	fmt.Println(&b==*(&c))
	//*c是&b的值  &*c取b的内存值
	fmt.Println(&b==&*c)


}

func modify1(a int)  {
	a=100
}

//类似引用传递
func modify2(a *int)  {
	*a=100
}