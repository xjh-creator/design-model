package main

import "fmt"
//相比于calculator.go,把业务逻辑封装起来，把ta与逻辑界面分开
func calculate(A,B int,operate string) int {
	switch operate {
	case "+":
		return A + B
	case "-":
		return A - B
	case "*":
		return A * B
	case "/":
		if B == 0{
			return -1
		}
		return A / B
	default:
		return -1
	}
}
func main()  {
	var A,B int
	var operate string
	fmt.Println("请输入数字A：")
	fmt.Scanf("%d",&A)
	fmt.Println("请选择运算符（+、-、*、/）：")
	fmt.Scanf("%s",&operate)
	fmt.Println("请输入数字B：")
	fmt.Scanf("%d",&B)
	result := calculate(A,B,operate)
	fmt.Println("计算结果为（-1表示有错）：",result)
}

