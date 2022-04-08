package main

import "fmt"

//初步实现一个简历

type Resume struct {
	name string
	sex string
	age string
	timeArea string
	company string
}

func (r *Resume)Display()  {
	fmt.Printf("姓名：%s,性别：%s,年龄：%s",r.name,r.sex,r.age)
	fmt.Println()
	fmt.Printf("工作经历：%s,%s",r.timeArea,r.company)
}

//但需要多次实例化时，这种客户端代码就有点麻烦，需要二十次，就实例化二十次
//如果写错一个字，就需要改二十次
func main()  {
	resume := &Resume{
		name: "小明",
		sex:"男",
		age:"14",
		timeArea: "上午",
		company: "科技公司",
	}
	resume.Display()
}
