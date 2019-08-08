/**
Go 运算符基本 与Java一致  *指针 &地址 与C语言一致
 */
package main

import "fmt"
func main()  {
   var a int =4
   var b *int //声明一个int指针b 单独开辟空间存储
   b=&a  //取出a变量的地址 给指针 b
   fmt.Println(a,*b)



}