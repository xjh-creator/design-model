package main

import "fmt"

//总计金额
var total float32

//商城简单收银软件
func main()  {
	var totalPrice float32
	for{
		var A float32
		var B int
		fmt.Println("请输入单价：")
		fmt.Scanf("%f",&A)
		fmt.Println("请输入数量：")
		fmt.Scanf("%d",&B)
		totalPrice = A * float32(B)
		total += totalPrice
		fmt.Println("总计金额为：",total)

		var operate string
		fmt.Println("请输入接下来操作(1重置、0退出程序、其它 继续计算)：")
		fmt.Scanf("%s",&operate)
		if operate == "1"{
			total = 0
		}else if operate == "0"{
			break
		}
	}
}