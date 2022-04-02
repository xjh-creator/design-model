package main

import "fmt"

var a,b int

type Calculate3 interface {
	GetResult() int
}

type Operation struct {
	a int
	b int
}

type OperationAdd struct {
	*Operation
}

func (a *OperationAdd)GetResult() int {
	return a.a + a.b
}

type OperationSub struct {
	*Operation
}

func (s *OperationSub)GetResult() int {
	return s.a - s.b
}

type OperationMul struct {
	*Operation
}

func (m *OperationMul)GetResult() int {
	return m.a * m.b
}

type OperationDiv struct {
	*Operation
}

func (a *OperationDiv)GetResult() int {
	if a.b == 0{
		return -1
	}
	return a.a / a.b
}

type OperationFactory struct {

}

//输入运算符，工厂就实例化出合适的对象
//通过多态形式，返回父类的方式实现了计算器的结果
func (f *OperationFactory)createOperate(operate string) Calculate3 {
	operation := &Operation{a:a,b:b}
	switch operate {
	case "+":
		return &OperationAdd{operation}
	case "-":
		return &OperationSub{operation}
	case "*":
		return &OperationMul{operation}
	case "/":
		return &OperationDiv{operation}
	default:
		return nil
	}
}

func main()  {
	var operate string
	fmt.Println("请输入数字A：")
	fmt.Scanf("%d",&a)
	fmt.Println("请选择运算符（+、-、*、/）：")
	fmt.Scanf("%s",&operate)
	fmt.Println("请输入数字B：")
	fmt.Scanf("%d",&b)

	factory := new(OperationFactory)
	calculate := factory.createOperate(operate)
	fmt.Println("计算结果为（-1表示有错）：",calculate.GetResult())

}
