package main

import "fmt"

//测试
func test()  {
	//定义一个长度为3的数组
	/**
	长度不一致的数组 或者类型不一致 不可赋值交换
	*/
	//会自动初始化 [0,0,0]
	var a[3] int
	fmt.Println(a)

	//数组初始化
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
	//数组初始化自动推测长度
	var b =[...]int{1,2,3,4,5}
	var c  = []string{"a","b","c"}//切片 可变数组 引用类型
	fmt.Println(len(b),len(c))

	//数组遍历
	for i:=0;i<len(c) ;i++  {
		fmt.Print(c[i])
	}
	fmt.Println()

	for index,value:=range b {
		fmt.Println(index,value)
	}


	//二维数组
	aa := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(aa) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(aa[2][1]) //支持索引取值:重庆
	//当有返回值必须接收 上下文不用 可以使用 _ 接收
	//二维数组遍历
	for _,v1:= range aa {
		for _,v2 := range v1 {
			fmt.Print(v2)
		}
		fmt.Println()
	}

}


func main()  {
	//test()


}



