package main

import "fmt"

//相比于calculator1.go,除了封装性好，它还使得扩展性便强、耦合性更低，如要要添加新的计算方法，直接新增一个方法
//在新增方法过程中又不会看到其它计算方法或影响到其它计算方法。
//现在问题是怎样让计算器知道我们希望使用哪个算法，这就要用到简单工厂模式。
type Calculate struct {
	a int
	b int
	operate string
}

func (c *Calculate)Add() int {
	return c.a + c.b
}

func (c *Calculate)Sub() int {
	return c.a - c.b
}

func (c *Calculate)Mul() int {
	return c.a * c.b
}

func (c *Calculate)Div() int {
	if c.b == 0{
		return -1
	}
	return c.a / c.b
}

func main()  {
	calculate := &Calculate{}
	fmt.Println("请输入数字A：")
	fmt.Scanf("%d",calculate.a)
	fmt.Println("请选择运算符（+、-、*、/）：")
	fmt.Scanf("%s",calculate.operate)
	fmt.Println("请输入数字B：")
	fmt.Scanf("%d",calculate.b)
	var result int
	switch calculate.operate {
	case "+":
		result = calculate.Add()
	case "-":
		result = calculate.Sub()
	case "*":
		result = calculate.Mul()
	case "/":
		result = calculate.Div()
	default:
		result = -1
	}
	fmt.Println("计算结果为（-1表示有错）：",result)
}


