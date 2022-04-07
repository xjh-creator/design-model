package main

import "fmt"


type LeiFeng interface {
	Sweep()
	Wash()
	BuyRice()
}

type Person struct {
	name string
}

func (s *Person)Sweep()  {
	fmt.Println(s.name,"扫地")
}

func (s *Person)Wash()  {
	fmt.Println(s.name,"洗衣服")
}

func (s *Person)BuyRice()  {
	fmt.Println(s.name,"买米")
}

type Undergraduate struct {
	*Person
}

func NewUndergraduate() *Undergraduate{
	return &Undergraduate{&Person{"毕业生"}}
}

type Volunteer struct {
	*Person
}

func NewVolunteer() *Volunteer{
	return &Volunteer{&Person{"志愿者"}}
}

type SimpleFactory struct {

}

func NewSimpleFactory(personType string) LeiFeng{
	switch personType {
	case "在校生":
		return NewUndergraduate()
	case "志愿者":
		return NewVolunteer()
	default:
		return nil
	}
}

//这里有几个问题
//1.就是如果都是创建在校生的实例，NewSimpleFactory("在校生")这个方法就会重复很多遍，这就有些不够优雅
//2.如果要增加出在校生、志愿者等学雷锋的角色，就必须修改工厂方法的case分支条件，修改了原有的类
//这不仅是对扩展开放了，也对修改开放了，违背了开放-封闭原则（软件实体可以扩展，但不可以修改）
func main()  {
	studentA := NewSimpleFactory("在校生")
	studentA.Wash()
	studentA.Sweep()
	studentA.BuyRice()

	volunteer := NewSimpleFactory("志愿者")
	volunteer.Wash()
	volunteer.Sweep()
	volunteer.BuyRice()
}