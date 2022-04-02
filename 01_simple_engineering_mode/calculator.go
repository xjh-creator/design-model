package main

import "fmt"

func main()  {
	var A,B,result int
	var operate string
	fmt.Println("请输入数字A：")
	fmt.Scanf("%d",&A)
	fmt.Println("请选择运算符（+、-、*、/）：")
	fmt.Scanf("%s",&operate)
	fmt.Println("请输入数字B：")
	fmt.Scanf("%d",&B)
	switch operate {
	case "+":
		result = A + B
	case "-":
		result = A - B
	case "*":
		result = A * B
	case "/":
		result = A / B
	default:
		result = -1
	}
	fmt.Println("计算结果为（-1表示有错）：",result)
}
