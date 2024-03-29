package main
/**
一道面试题 说出输出结果
 */
import "fmt"

type student struct {
	name string
	age  int
}

func getNewStudent(name string,age int) *student  {
	return &student{
		name,
		age,
	}
}

func (s *student)setName(name string)  {
	s.name=name
}

func (s *student) getName() string  {
	return s.name
}

func main()  {
	stu:=getNewStudent("zzq",20)
	//因为stu是一个指针类型 正常情况下应根据22行代码取值 22行和23行结果一致 是因为go的语法糖
	fmt.Println((*stu).age)
	fmt.Println(stu.age)

	stu.setName("qwe")
	fmt.Println(stu.getName())

}

func test() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for index, stu := range stus {
		//fmt.Printf("%p",&stu)
		//fmt.Println()
		fmt.Printf("%p",&stus[index])
		fmt.Println()

		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	/**
	输出结果如下
	小王子 => 大王八
	娜扎 => 大王八
	大王八 => 大王八
	为什么都是大王八呢？
	看20行代码中 stu变量 因为这个变量在遍历的时候 只会改变 内容 而不会改变地址
	看22行代码 是取第20行中的stu变量的地址
	最后遍历的内容是{name: "大王八", age: 9000}  那&stu就指向了这条数据 而map里的value 存的都是&stu  所以都是 大王八
	代码21行可以看出stu是同一块内存
	 */
}
