package main

import "fmt"

type Operation interface {
	GetResult() float64
}

type Number struct {
	NumberA float64
	NumberB float64
}

type OperationAdd struct {
	*Number
}

func (a *OperationAdd) GetResult() float64 {
	return a.NumberA + a.NumberB
}

type OperationSub struct {
	*Number
}

func (s *OperationSub) GetResult() float64 {
	return s.NumberA - s.NumberB
}

type OperationMul struct {
	*Number
}

func (m *OperationMul) GetResult() float64 {
	return m.NumberA * m.NumberB
}

type OperationDiv struct {
	*Number
}

func (a *OperationDiv) GetResult() float64 {
	if a.NumberB == 0 {
		return -1
	}
	return a.NumberA / a.NumberB
}

func NewOperation(operate string, a, b float64) Operation {
	number := &Number{NumberA: a, NumberB: b}
	switch operate {
	case "+":
		return &OperationAdd{number}
	case "-":
		return &OperationSub{number}
	case "*":
		return &OperationMul{number}
	case "/":
		return &OperationDiv{number}
	default:
		return nil
	}
}

type OperationFactory struct {
}

//输入运算符，工厂就实例化出合适的对象
//通过多态形式，返回父类的方式实现了计算器的结果
func (f *OperationFactory) createOperate(operate string, a, b float64) Operation {
	number := &Number{NumberA: a, NumberB: b}
	switch operate {
	case "+":
		return &OperationAdd{number}
	case "-":
		return &OperationSub{number}
	case "*":
		return &OperationMul{number}
	case "/":
		return &OperationDiv{number}
	default:
		return nil
	}
}

func main() {
	var operate string
	var a, b float64
	fmt.Println("请输入数字A：")
	fmt.Scanf("%f", &a)
	fmt.Println("请选择运算符（+、-、*、/）：")
	fmt.Scanf("%s", &operate)
	fmt.Println("请输入数字B：")
	fmt.Scanf("%f", &b)

	factory := new(OperationFactory)
	calculate := factory.createOperate(operate, a, b)
	fmt.Println("计算结果为（-1表示有错）：", calculate.GetResult())

}
