package main

import "fmt"

type PersonalInfo struct {
	name string
	sex  string
	age  string
}

type WorkExperience struct {
	timeArea string
	company  string
}

type Resume struct {
	PersonalInfo
	WorkExperience
}

func (this *Resume) SetPersonalInfo(name string, sex string, age string) {
	this.name = name
	this.sex = sex
	this.age = age
}

func (this *Resume) SetWorkExperience(timeArea string, company string) {
	this.timeArea = timeArea
	this.company = company
}

func (this *Resume) Display() {
	fmt.Println(this.name, this.sex, this.age)
	fmt.Println("工作经历：", this.timeArea, this.company)
}

func (this *Resume) Clone() *Resume {
	//resume := &Resume{
	//	this.PersonalInfo,
	//	this.WorkExperience,
	//}
	resume := *this
	//拷贝了一个值，这种写法比new快
	return &resume
}

func main() {
	r1 := &Resume{}
	r1.SetPersonalInfo("大鸟", "男", "29")
	r1.SetWorkExperience("1998-2000", "xx公司")
	fmt.Println("地址1",&r1)

	r2 := r1.Clone()
	r2.SetWorkExperience("2001-2006", "yy公司")
	fmt.Println("地址1",&r2)

	r3 := r1.Clone()
	r3.SetPersonalInfo("大鸟", "男", "24")
	r3.SetWorkExperience("1998-2003", "zz公司")
	fmt.Println("地址1",&r3)

	r1.Display()
	r2.Display()
	r3.Display()
}

