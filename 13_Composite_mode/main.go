package main

import (
	"fmt"
	"strings"
)

type ICompany interface {
	Add(c ICompany) //增加
	Remove(c ICompany) //移除
	Display(depth int) //显示
	LineOfDuty() //履行职责，不同部门履行不同职责
}

type ConcreteCompany struct {
	name string
	children map[ICompany]ICompany
}

func NewConcreteCompany(name string) *ConcreteCompany {
	return &ConcreteCompany{
		name:name,
		children: make(map[ICompany]ICompany),
	}
}

func (c *ConcreteCompany)Add(company ICompany)  {
	c.children[company] = company
}

func (c *ConcreteCompany)Remove(company ICompany)  {
	delete(c.children, company)
}

func (c *ConcreteCompany)Display(depth int)  {
	fmt.Println(strings.Repeat("-",depth)," ",c.name)
	for _,v := range c.children{
		v.Display(depth+2)
	}
}

func (c *ConcreteCompany)LineOfDuty()  {
	for _,v := range c.children{
		v.LineOfDuty()
	}
}

type HRDepartment struct {
	name string
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{
		name:name,
	}
}

func (c *HRDepartment)Add(company ICompany)  {
}

func (c *HRDepartment)Remove(company ICompany)  {
}

func (c *HRDepartment)Display(depth int)  {
	fmt.Println(strings.Repeat("-",depth)," ",c.name)
}

func (c *HRDepartment)LineOfDuty()  {
	fmt.Println(c.name,"员工招聘培训管理")
}

type FinanceDepartment struct {
	name string
}

func NewFinanceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{
		name:name,
	}
}

func (c *FinanceDepartment)Add(company ICompany)  {
}

func (c *FinanceDepartment)Remove(company ICompany)  {
}

func (c *FinanceDepartment)Display(depth int)  {
	fmt.Println(strings.Repeat("-",depth)," ",c.name)
}

func (c *FinanceDepartment)LineOfDuty()  {
	fmt.Println(c.name,"公司财务收支管理")
}

func main()  {
	root := NewConcreteCompany("北京总公司")
	root.Add(NewHRDepartment("总公司人力资源部"))
	root.Add(NewHRDepartment("总公司财务部"))

	comp := NewConcreteCompany("上海华东分公司")
	comp.Add(NewHRDepartment("上海华东分公司人力资源部"))
	comp.Add(NewHRDepartment("上海华东分公司财务部"))
	root.Add(comp)

	comp1 := NewConcreteCompany("南京办事处")
	comp1.Add(NewHRDepartment("南京办事处人力资源部"))
	comp1.Add(NewHRDepartment("南京办事处财务部"))
	root.Add(comp1)

	comp2 := NewConcreteCompany("杭州办事处")
	comp2.Add(NewHRDepartment("杭州办事处人力资源部"))
	comp2.Add(NewHRDepartment("杭州办事处财务部"))
	root.Add(comp2)

	fmt.Println("\n结构图：")
	root.Display(1)

	fmt.Println("\n职责：")
	root.LineOfDuty()

}