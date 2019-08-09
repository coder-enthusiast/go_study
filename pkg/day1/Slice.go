/**
切片
 */
package main

import "fmt"




func main()  {
	//声明并初始化切片
	a:=[]int{}
	//fmt.Println(a)
	//切片追加值
	a=append(a,1,2,3,4,5)
	fmt.Println(a,len(a),cap(a))

	b := []string{}
	b = append(b,"a","中","国")
	fmt.Println(b,len(b),cap(b))


	var aa = make([]string, 3,10)
	fmt.Println(len(aa),cap(aa))//长度为3 容量为10
	for i := 0; i < 10; i++ {
		aa = append(aa, fmt.Sprintf("%v", i))
	}


	fmt.Println(aa,len(aa),cap(aa))//[   0 1 2 3 4 5 6 7 8 9] 13 20  append追加


	for i := 0; i < 10; i++ {
		aa[i] = fmt.Sprint(i)
	}
	fmt.Println(aa,len(aa),cap(aa))
}