package pck

import "fmt"

func init() {
	fmt.Println("pac包")
}

type Staff struct {
	Name string
	JobName string
	No string
}


func (s Staff) GetNameRemark() string {
	return "姓名"
}

func (s Staff) GetJobNameRemark() string  {
	return "职位名称"
}

func (s Staff) GetNoRemark() string  {
	return "工号"
}

