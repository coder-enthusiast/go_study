package main

import "fmt"

/**
函数
 */
//普通函数
func sum(x int ,y int ) int  {
	return x+y
}

//可变长参数 多返回值
func Accumulate(x ...int) (int,int)  {
	sum:=0
	for _,v := range x{
		sum +=v
	}
	return sum,len(x)
}
//返回值 变量命名
func Accumulate1(x ...int) (sum int,lenght int)  {
	sum = 0
	for _,v := range x{
		sum +=v
	}
	lenght = len(x)
	return
}

//声明一个函数为Person 类型
/**
函数声明类型与Person一致的都可以赋值给Person
 */
type Person func(str string) string

func sayHello(str string) string  {
	return "hello "+str
}



//高阶函数 函数作为参数
func funcPara(a int , b int, sum func(x int,y int) int) int  {
	return sum(a,b)
}

//高阶函数 返回值为函数
func operate(op string) (func(a int,b int) int) {
	switch op {
	case "-":
		return func(a int, b int) int {
			return a - b
		}
	case "+":
		return func(a int, b int) int {
			return a + b
		}
	case "*":
		return func(a int, b int) int {
			return a * b
		}
	case "/":
		return func(a int, b int) int {
			return a / b
		}
	case "%":
		return func(a int, b int) int {
			return a % b
		}
	default:
		return func(a int, b int) int {
			return 0
		}
	}
}

//闭包
func bb(x int) func(int) int  {
	fmt.Println("x 的值",x)
	return func(i int) int {
		i+=x
		return i
	}
}


func main() {
	sumResult:= sum(1,2)
	fmt.Println(sumResult)

	acc,len := Accumulate(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(acc,len)


	var person Person
	person=sayHello
	fmt.Println(person("zzq"))

	fmt.Println(funcPara(1,3,sum))

	result := operate("/")
	fmt.Println(result(1,3))


	//匿名函数  把函数给一个变量
	anonymous:=func(){
		fmt.Println("匿名函数")
	}
	anonymous()

	//自执行函数
	func(){
		fmt.Println("自执行函数")
	}()

	b:=bb(10)
	fmt.Println(b(10))
	fmt.Println(b(20))
}

